#!/bin/bash
# =============================================================================
# Game Engine Casino - K3s Automated Deployment
# =============================================================================
# Usage: ./k3s-deploy.sh
# Auto-generates all secrets, deploys everything to K3s
# =============================================================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log() { echo -e "${GREEN}[✓]${NC} $1"; }
warn() { echo -e "${YELLOW}[!]${NC} $1"; }
error() { echo -e "${RED}[✗]${NC} $1"; }
step() { echo -e "\n${BLUE}=== $1 ===${NC}"; }

# =============================================================================
# STEP 1: Install K3s
# =============================================================================
step "Step 1: Installing K3s"

if command -v k3s &>/dev/null; then
    log "K3s already installed"
else
    curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server \
        --disable traefik \
        --disable servicelb \
        --write-kubeconfig-mode 644 \
        --kubelet-arg='max-pods=250'" sh -
    
    # Wait for K3s
    sleep 10
    log "K3s installed"
fi

# Setup kubeconfig
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config 2>/dev/null || true
sudo chown $USER:$USER ~/.kube/config 2>/dev/null || true
export KUBECONFIG=~/.kube/config
echo 'export KUBECONFIG=~/.kube/config' >> ~/.bashrc 2>/dev/null || true

# Verify
kubectl get nodes
log "K3s cluster ready"

# =============================================================================
# STEP 2: Install MetalLB
# =============================================================================
step "Step 2: Installing MetalLB"

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.12/config/manifests/metallb-native.yaml 2>/dev/null || true
kubectl wait --namespace metallb-system --for=condition=ready pod --selector=app=metallb --timeout=120s 2>/dev/null || true

SERVER_IP=$(hostname -I | awk '{print $1}')
[ -z "$SERVER_IP" ] && SERVER_IP="127.0.0.1"

cat <<EOF | kubectl apply -f - 2>/dev/null || true
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: game-engine-pool
  namespace: metallb-system
spec:
  addresses:
  - ${SERVER_IP}/32
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: game-engine-l2
  namespace: metallb-system
spec:
  ipAddressPools:
  - game-engine-pool
EOF

log "MetalLB configured with IP ${SERVER_IP}"

# =============================================================================
# STEP 3: Generate Secrets
# =============================================================================
step "Step 3: Auto-generating secrets"

DB_PASSWORD=$(openssl rand -base64 16 | tr -d '\n/+')
REDIS_PASSWORD=$(openssl rand -base64 16 | tr -d '\n/+')
JWT_SECRET=$(openssl rand -base64 64 | tr -d '\n')
WS_SECRET=$(openssl rand -base64 64 | tr -d '\n')
ADMIN_PASSWORD=$(openssl rand -base64 12 | tr -d '\n/+')

# Generate RSA keys
mkdir -p certs
openssl genrsa -out certs/jwt-private.pem 4096 2>/dev/null
openssl rsa -in certs/jwt-private.pem -pubout -out certs/jwt-public.pem 2>/dev/null

log "All secrets auto-generated"

# =============================================================================
# STEP 4: Create Namespaces
# =============================================================================
step "Step 4: Creating namespaces"

kubectl create namespace game-engine 2>/dev/null || true
kubectl create namespace databases 2>/dev/null || true
kubectl create namespace monitoring 2>/dev/null || true

log "Namespaces created"

# =============================================================================
# STEP 5: Create Secrets and ConfigMaps
# =============================================================================
step "Step 5: Creating Kubernetes secrets"

# Create TLS secret from RSA keys
kubectl create secret tls jwt-tls \
    --cert=certs/jwt-public.pem \
    --key=certs/jwt-private.pem \
    -n game-engine 2>/dev/null || true

# Create application secrets
kubectl create secret generic game-engine-secrets \
    --from-literal=DB_PASSWORD="${DB_PASSWORD}" \
    --from-literal=REDIS_PASSWORD="${REDIS_PASSWORD}" \
    --from-literal=JWT_SECRET="${JWT_SECRET}" \
    --from-literal=WS_SECRET_KEY_BASE="${WS_SECRET}" \
    --from-literal=ADMIN_PASSWORD="${ADMIN_PASSWORD}" \
    -n game-engine 2>/dev/null || kubectl create secret generic game-engine-secrets \
    --from-literal=DB_PASSWORD="${DB_PASSWORD}" \
    --from-literal=REDIS_PASSWORD="${REDIS_PASSWORD}" \
    --from-literal=JWT_SECRET="${JWT_SECRET}" \
    --from-literal=WS_SECRET_KEY_BASE="${WS_SECRET}" \
    --from-literal=ADMIN_PASSWORD="${ADMIN_PASSWORD}" \
    -n game-engine --dry-run=client -o yaml | kubectl apply -f -

# Database secrets
kubectl create secret generic database-credentials \
    --from-literal=username=postgres \
    --from-literal=password="${DB_PASSWORD}" \
    -n databases 2>/dev/null || kubectl create secret generic database-credentials \
    --from-literal=username=postgres \
    --from-literal=password="${DB_PASSWORD}" \
    -n databases --dry-run=client -o yaml | kubectl apply -f -

# Create shared configmap
kubectl create configmap game-engine-config \
    --from-literal=DB_HOST=postgres.databases.svc.cluster.local \
    --from-literal=DB_PORT=5432 \
    --from-literal=DB_USER=postgres \
    --from-literal=REDIS_HOST=redis.databases.svc.cluster.local \
    --from-literal=REDIS_PORT=6379 \
    --from-literal=NATS_HOST=nats.databases.svc.cluster.local \
    --from-literal=NATS_PORT=4222 \
    --from-literal=ENVIRONMENT=production \
    --from-literal=LOG_LEVEL=info \
    -n game-engine 2>/dev/null || kubectl create configmap game-engine-config \
    --from-literal=DB_HOST=postgres.databases.svc.cluster.local \
    --from-literal=DB_PORT=5432 \
    --from-literal=DB_USER=postgres \
    --from-literal=REDIS_HOST=redis.databases.svc.cluster.local \
    --from-literal=REDIS_PORT=6379 \
    --from-literal=NATS_HOST=nats.databases.svc.cluster.local \
    --from-literal=NATS_PORT=4222 \
    --from-literal=ENVIRONMENT=production \
    --from-literal=LOG_LEVEL=info \
    -n game-engine --dry-run=client -o yaml | kubectl apply -f -

log "Secrets and ConfigMaps created"

# =============================================================================
# STEP 6: Deploy Databases
# =============================================================================
step "Step 6: Deploying databases"

# Create init script configmap
kubectl create configmap postgres-init \
    --from-file=01-create-databases.sql=infrastructure/docker/postgres/init/01-create-databases.sql \
    -n databases 2>/dev/null || kubectl create configmap postgres-init \
    --from-file=01-create-databases.sql=infrastructure/docker/postgres/init/01-create-databases.sql \
    -n databases --dry-run=client -o yaml | kubectl apply -f -

# PostgreSQL PVC
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgres-data
  namespace: databases
spec:
  accessModes: [ReadWriteOnce]
  resources:
    requests:
      storage: 200Gi
  storageClassName: local-path
EOF

# PostgreSQL Deployment
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
  namespace: databases
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
        - name: postgres
          image: postgres:16-alpine
          ports:
            - containerPort: 5432
          env:
            - name: POSTGRES_USER
              value: postgres
            - name: POSTGRES_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: database-credentials
                  key: password
          resources:
            requests:
              memory: "4Gi"
              cpu: "2"
            limits:
              memory: "16Gi"
              cpu: "8"
          args:
            - "postgres"
            - "-c" 
            - "shared_buffers=4GB"
            - "-c"
            - "effective_cache_size=12GB"
            - "-c"
            - "work_mem=64MB"
            - "-c"
            - "max_connections=500"
          volumeMounts:
            - name: postgres-data
              mountPath: /var/lib/postgresql/data
            - name: init-scripts
              mountPath: /docker-entrypoint-initdb.d
      volumes:
        - name: postgres-data
          persistentVolumeClaim:
            claimName: postgres-data
        - name: init-scripts
          configMap:
            name: postgres-init
---
apiVersion: v1
kind: Service
metadata:
  name: postgres
  namespace: databases
spec:
  selector:
    app: postgres
  ports:
    - port: 5432
      targetPort: 5432
  type: ClusterIP
EOF

# Redis
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
  namespace: databases
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
        - name: redis
          image: redis:7-alpine
          ports:
            - containerPort: 6379
          resources:
            requests:
              memory: "512Mi"
              cpu: "250m"
            limits:
              memory: "2Gi"
              cpu: "1"
          command: ["redis-server", "--maxmemory", "1536mb", "--maxmemory-policy", "allkeys-lru"]
---
apiVersion: v1
kind: Service
metadata:
  name: redis
  namespace: databases
spec:
  selector:
    app: redis
  ports:
    - port: 6379
      targetPort: 6379
  type: ClusterIP
EOF

# NATS
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nats
  namespace: databases
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nats
  template:
    metadata:
      labels:
        app: nats
    spec:
      containers:
        - name: nats
          image: nats:2.10-alpine
          ports:
            - containerPort: 4222
            - containerPort: 8222
          resources:
            requests:
              memory: "256Mi"
              cpu: "250m"
            limits:
              memory: "1Gi"
              cpu: "1"
---
apiVersion: v1
kind: Service
metadata:
  name: nats
  namespace: databases
spec:
  selector:
    app: nats
  ports:
    - name: client
      port: 4222
      targetPort: 4222
    - name: monitoring
      port: 8222
      targetPort: 8222
  type: ClusterIP
EOF

log "Databases deployed"

# Wait for databases
echo -n "  Waiting for PostgreSQL..."
for i in $(seq 1 60); do
    if kubectl exec -n databases deploy/postgres -- pg_isready -U postgres &>/dev/null; then
        echo " ✓"
        break
    fi
    echo -n "."
    sleep 2
done

echo -n "  Waiting for Redis..."
for i in $(seq 1 30); do
    if kubectl exec -n databases deploy/redis -- redis-cli ping &>/dev/null; then
        echo " ✓"
        break
    fi
    echo -n "."
    sleep 1
done

# =============================================================================
# STEP 7: Build and Deploy Application Services
# =============================================================================
step "Step 7: Building and deploying services"

# Deploy Go services
GO_SERVICES="auth-service:4433 user-service:4435 wallet-service:4437 game-registry:4439 card-games:9040 dice-games:9041 slot-games:9042 rng-service:9043 betting:9044 tournament:9020 jackpot-service:9091 leaderboard-service:9024 winners-showcase-service:9025 merchant-service:9095 agent-service:9083 loyalty-service:9096 live-dealer-service:9092 sports-betting-service:9093 game-engine:9094 multiplayer:9021 chat:9022 notification:9023 banner-service:9097 referral-service:9098"

for svc_port in $GO_SERVICES; do
    svc=$(echo $svc_port | cut -d: -f1)
    port=$(echo $svc_port | cut -d: -f2)
    
    # Build image
    if [ -f "services/$svc/Dockerfile" ]; then
        docker build -t game-engine/${svc}:latest services/${svc}/ -q 2>/dev/null || true
    fi
    
    # Deploy
    cat <<EOF | kubectl apply -f - 2>/dev/null
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${svc}
  namespace: game-engine
spec:
  replicas: 2
  selector:
    matchLabels:
      app: ${svc}
  template:
    metadata:
      labels:
        app: ${svc}
    spec:
      containers:
        - name: ${svc}
          image: game-engine/${svc}:latest
          imagePullPolicy: Never
          ports:
            - containerPort: ${port}
              name: grpc
          envFrom:
            - configMapRef:
                name: game-engine-config
          env:
            - name: SERVER_GRPC_PORT
              value: "${port}"
            - name: DB_NAME
              value: "${svc//-/_}"
            - name: DB_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: game-engine-secrets
                  key: DB_PASSWORD
          resources:
            requests:
              memory: "64Mi"
              cpu: "100m"
            limits:
              memory: "256Mi"
              cpu: "500m"
---
apiVersion: v1
kind: Service
metadata:
  name: ${svc}
  namespace: game-engine
spec:
  selector:
    app: ${svc}
  ports:
    - port: ${port}
      targetPort: ${port}
      name: grpc
  type: ClusterIP
EOF

    echo "  Deployed ${svc}:${port}"
done

log "Go services deployed"

# =============================================================================
# STEP 8: Deploy Gateways
# =============================================================================
step "Step 8: Deploying gateways"

# Player Gateway
docker build -t game-engine/gateway-player:latest services/gateway/ -f services/gateway/Dockerfile.player -q 2>/dev/null || true
cat <<EOF | kubectl apply -f - 2>/dev/null
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gateway-player
  namespace: game-engine
spec:
  replicas: 3
  selector:
    matchLabels:
      app: gateway-player
  template:
    metadata:
      labels:
        app: gateway-player
    spec:
      containers:
        - name: gateway-player
          image: game-engine/gateway-player:latest
          imagePullPolicy: Never
          ports:
            - containerPort: 8080
          env:
            - name: JWT_SECRET
              valueFrom:
                secretKeyRef:
                  name: game-engine-secrets
                  key: JWT_SECRET
            - name: REDIS_HOST
              value: "redis.databases.svc.cluster.local"
          resources:
            requests:
              memory: "128Mi"
              cpu: "200m"
            limits:
              memory: "512Mi"
              cpu: "1"
---
apiVersion: v1
kind: Service
metadata:
  name: gateway-player
  namespace: game-engine
spec:
  type: LoadBalancer
  selector:
    app: gateway-player
  ports:
    - port: 80
      targetPort: 8080
      name: http
EOF

# Admin Gateway
docker build -t game-engine/gateway-admin:latest services/gateway/admin/ -q 2>/dev/null || true
cat <<EOF | kubectl apply -f - 2>/dev/null
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gateway-admin
  namespace: game-engine
spec:
  replicas: 2
  selector:
    matchLabels:
      app: gateway-admin
  template:
    metadata:
      labels:
        app: gateway-admin
    spec:
      containers:
        - name: gateway-admin
          image: game-engine/gateway-admin:latest
          imagePullPolicy: Never
          ports:
            - containerPort: 8081
          env:
            - name: JWT_SECRET
              valueFrom:
                secretKeyRef:
                  name: game-engine-secrets
                  key: JWT_SECRET
            - name: REDIS_HOST
              value: "redis.databases.svc.cluster.local"
          resources:
            requests:
              memory: "128Mi"
              cpu: "200m"
            limits:
              memory: "512Mi"
              cpu: "1"
---
apiVersion: v1
kind: Service
metadata:
  name: gateway-admin
  namespace: game-engine
spec:
  type: LoadBalancer
  selector:
    app: gateway-admin
  ports:
    - port: 80
      targetPort: 8081
      name: http
EOF

# WebSocket Gateway
docker build -t game-engine/websocket-gateway:latest services/websocket-gateway/ -q 2>/dev/null || true
cat <<EOF | kubectl apply -f - 2>/dev/null
apiVersion: apps/v1
kind: Deployment
metadata:
  name: websocket-gateway
  namespace: game-engine
spec:
  replicas: 3
  selector:
    matchLabels:
      app: websocket-gateway
  template:
    metadata:
      labels:
        app: websocket-gateway
    spec:
      containers:
        - name: websocket-gateway
          image: game-engine/websocket-gateway:latest
          imagePullPolicy: Never
          ports:
            - containerPort: 4000
          env:
            - name: REDIS_HOST
              value: "redis.databases.svc.cluster.local"
            - name: SECRET_KEY_BASE
              valueFrom:
                secretKeyRef:
                  name: game-engine-secrets
                  key: WS_SECRET_KEY_BASE
          resources:
            requests:
              memory: "128Mi"
              cpu: "200m"
            limits:
              memory: "512Mi"
              cpu: "1"
---
apiVersion: v1
kind: Service
metadata:
  name: websocket-gateway
  namespace: game-engine
spec:
  type: LoadBalancer
  selector:
    app: websocket-gateway
  ports:
    - port: 80
      targetPort: 4000
      name: ws
EOF

# Admin Panel
docker build -t game-engine/admin-panel:latest admin/ -q 2>/dev/null || true
cat <<EOF | kubectl apply -f - 2>/dev/null
apiVersion: apps/v1
kind: Deployment
metadata:
  name: admin-panel
  namespace: game-engine
spec:
  replicas: 2
  selector:
    matchLabels:
      app: admin-panel
  template:
    metadata:
      labels:
        app: admin-panel
    spec:
      containers:
        - name: admin-panel
          image: game-engine/admin-panel:latest
          imagePullPolicy: Never
          ports:
            - containerPort: 80
          resources:
            requests:
              memory: "32Mi"
              cpu: "50m"
            limits:
              memory: "128Mi"
              cpu: "200m"
---
apiVersion: v1
kind: Service
metadata:
  name: admin-panel
  namespace: game-engine
spec:
  type: LoadBalancer
  selector:
    app: admin-panel
  ports:
    - port: 80
      targetPort: 80
      name: http
EOF

log "Gateways deployed"

# =============================================================================
# STEP 9: Wait and Verify
# =============================================================================
step "Step 9: Verifying deployment"

echo -n "  Waiting for pods to start..."
sleep 20
echo " ✓"

# Count pods
TOTAL_PODS=$(kubectl get pods -n game-engine --no-headers 2>/dev/null | wc -l)
RUNNING_PODS=$(kubectl get pods -n game-engine --field-selector=status.phase=Running --no-headers 2>/dev/null | wc -l)
DB_PODS=$(kubectl get pods -n databases --field-selector=status.phase=Running --no-headers 2>/dev/null | wc -l)

log "Game Engine pods: ${RUNNING_PODS}/${TOTAL_PODS} running"
log "Database pods: ${DB_PODS}/3 running"

# Get service IPs
echo ""
echo "Service URLs:"
kubectl get svc -n game-engine -o wide 2>/dev/null | grep LoadBalancer | awk '{print "  " $1 ": http://" $4}'

# =============================================================================
# STEP 10: Save Credentials
# =============================================================================
step "Step 10: Saving credentials"

cat > k3s-credentials.txt << EOF
===========================================
Game Engine Casino - K3s Deployment
Generated: $(date)
===========================================

Cluster:
  Server IP: ${SERVER_IP}
  K3s: $(k3s --version 2>/dev/null | head -1 || echo "installed")

Access URLs:
  Player API:    http://${SERVER_IP}:80
  Admin API:     http://${SERVER_IP}:8081 (or via LoadBalancer)
  WebSocket:     ws://${SERVER_IP}:4000
  Admin Panel:   http://${SERVER_IP}:3000 (or via LoadBalancer)

Admin Login:
  Username: admin
  Password: ${ADMIN_PASSWORD}

Database:
  Host: postgres.databases.svc.cluster.local
  Port: 5432
  User: postgres
  Password: ${DB_PASSWORD}

Secrets:
  JWT: ${JWT_SECRET}
  Redis: ${REDIS_PASSWORD}

Useful Commands:
  kubectl get pods -n game-engine
  kubectl get svc -n game-engine
  kubectl logs -n game-engine -l app=auth-service
  kubectl scale deployment -n game-engine gateway-player --replicas=5
  kubectl delete namespace game-engine databases monitoring

===========================================
KEEP THIS FILE SECURE!
===========================================
EOF

chmod 600 k3s-credentials.txt
log "Credentials saved to k3s-credentials.txt"

# =============================================================================
# COMPLETE
# =============================================================================
echo ""
echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN}  K3s Deployment Complete!${NC}"
echo -e "${GREEN}============================================${NC}"
echo ""
echo "  Running pods: ${RUNNING_PODS}/${TOTAL_PODS}"
echo "  Credentials:  ./k3s-credentials.txt"
echo "  Admin Login:  admin / ${ADMIN_PASSWORD}"
echo ""
echo "  kubectl get pods -n game-engine"
echo "  kubectl get svc -n game-engine"
echo ""
