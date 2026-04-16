module github.com/game_engine/betting

go 1.25.0

replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.9.0

require (
	github.com/game_engine/gen/go v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.9.1
	google.golang.org/grpc v1.66.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)

replace github.com/game_engine/gen/go => ../../proto/gen/go
