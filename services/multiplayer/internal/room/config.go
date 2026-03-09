package room

// Config holds multiplayer service configuration
type Config struct {
	Server struct {
		GRPCPort int `yaml:"grpc_port"`
		HTTPPort int `yaml:"http_port"`
	} `yaml:"server"`

	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`

	NATS struct {
		URLs      []string `yaml:"urls"`
		ClusterID string   `yaml:"cluster_id"`
		ClientID  string   `yaml:"client_id"`
	} `yaml:"nats"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"database"`

	Multiplayer struct {
		DefaultMinPlayers int  `yaml:"default_min_players"`
		DefaultMaxPlayers int  `yaml:"default_max_players"`
		DefaultMaxTables  int  `yaml:"default_max_tables"`
		TableTimeout      int  `yaml:"table_timeout"`
		IdleTimeout       int  `yaml:"idle_timeout"`
		MaxSpectators     int  `yaml:"max_spectators"`
		AutoDisconnect    bool `yaml:"auto_disconnect"`
	} `yaml:"multiplayer"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}
