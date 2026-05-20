# Makefile for building and running all services

MAVEN_SERVICES = affiliate-service bonus-service commission-service payment-service
GO_SERVICES = agent-service user-service winners-showcase-service loyalty-service \
              banner-service betting game-registry card-games auth-service \
              referral-service dice-games wallet-service tournament chat \
              live-dealer-service sports-betting-service slot-games game-engine \
              multiplayer notification leaderboard-service merchant-service \
              rng-service jackpot-service
GATEWAY_SERVICES = gateway/admin gateway/merchant gateway/player gateway/agent
PYTHON_SERVICES = aml-service fraud-service risk-service kyc-service
ELIXIR_SERVICES = websocket-gateway

BIN_DIR = bin

.PHONY: all build build-java build-go build-python build-elixir run run-java run-go run-python run-elixir clean

all: build

build: build-java build-go build-python build-elixir

build-java:
	@echo "Building Java services..."
	@for service in $(MAVEN_SERVICES); do \
		echo "Building $$service..."; \
		(cd $$service && mvn clean package -DskipTests); \
	done

build-go:
	@echo "Building Go services..."
	@mkdir -p $(BIN_DIR)
	@for service in $(GO_SERVICES); do \
		echo "Building $$service..."; \
		go build -o $(BIN_DIR)/$$service ./$$service/cmd/main.go; \
	done
	@for service in $(GATEWAY_SERVICES); do \
		name=$$(basename $$service); \
		echo "Building gateway-$$name..."; \
		go build -o $(BIN_DIR)/gateway-$$name ./$$service/main.go; \
	done

build-python:
	@echo "Building Python services (installing dependencies)..."
	@for service in $(PYTHON_SERVICES); do \
		echo "Preparing $$service..."; \
		(cd $$service && if [ ! -d venv ]; then python3 -m venv venv; fi && \
		./venv/bin/pip install -r requirements.txt); \
	done

build-elixir:
	@echo "Building Elixir services..."
	@for service in $(ELIXIR_SERVICES); do \
		echo "Building $$service..."; \
		(cd $$service && mix local.hex --force && mix local.rebar --force && mix deps.get && mix compile); \
	done

run:
	@echo "Running all services (in background)..."
	@echo "Note: This will start a lot of processes. Use 'make stop' to kill them."
	@$(MAKE) run-java run-go run-python run-elixir

run-java:
	@mkdir -p logs
	@for service in $(MAVEN_SERVICES); do \
		echo "Starting $$service..."; \
		(cd $$service && nohup java -jar target/*.jar > ../logs/$$service.log 2>&1 &); \
	done

run-go:
	@mkdir -p logs
	@for service in $(GO_SERVICES); do \
		echo "Starting $$service..."; \
		nohup ./$(BIN_DIR)/$$service > logs/$$service.log 2>&1 & \
	done
	@for service in $(GATEWAY_SERVICES); do \
		name=$$(basename $$service); \
		echo "Starting gateway-$$name..."; \
		nohup ./$(BIN_DIR)/gateway-$$name > logs/gateway-$$name.log 2>&1 & \
	done

run-python:
	@mkdir -p logs
	@for service in $(PYTHON_SERVICES); do \
		echo "Starting $$service..."; \
		(cd $$service && nohup ./venv/bin/uvicorn app.main:app --host 0.0.0.0 --port $$(grep EXPOSE Dockerfile | awk '{print $$2}') > ../logs/$$service.log 2>&1 &); \
	done

run-elixir:
	@mkdir -p logs
	@for service in $(ELIXIR_SERVICES); do \
		echo "Starting $$service..."; \
		(cd $$service && nohup mix phx.server > ../logs/$$service.log 2>&1 &); \
	done

stop:
	@echo "Stopping all services..."
	-pkill -f "java -jar"
	-pkill -f "./bin/"
	-pkill -f "uvicorn"
	-pkill -f "mix phx.server"

status:
	@echo "Checking running services..."
	@echo "Java services:"
	@pgrep -f "java -jar" | xargs -r ps -fp || echo "None"
	@echo "\nGo services:"
	@pgrep -f "./bin/" | xargs -r ps -fp || echo "None"
	@echo "\nPython services:"
	@pgrep -f "uvicorn" | xargs -r ps -fp || echo "None"
	@echo "\nElixir services:"
	@pgrep -f "mix phx.server" | xargs -r ps -fp || echo "None"

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR) logs
	@for service in $(MAVEN_SERVICES); do (cd $$service && mvn clean); done
	@for service in $(ELIXIR_SERVICES); do (cd $$service && mix clean); done
