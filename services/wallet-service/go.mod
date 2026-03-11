module github.com/gameengine/wallet-service

go 1.24.0

require (
	github.com/google/uuid v1.5.0
	github.com/lib/pq v1.10.9
	github.com/nats-io/nats.go v1.48.0
	github.com/redis/go-redis/v9 v9.4.0
	google.golang.org/grpc v1.60.1
	gopkg.in/yaml.v3 v3.0.1
)


replace github.com/gameengine/gen/go => ../../proto/gen/go
