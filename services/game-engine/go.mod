module github.com/game-engine/game-engine

go 1.25

require (
	game_engine/gen/go v0.0.0
	google.golang.org/grpc v1.60.1
	gopkg.in/yaml.v3 v3.0.1
)

replace game_engine/gen/go => ../../proto/gen/go
