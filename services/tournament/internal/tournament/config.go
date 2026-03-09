package tournament

// Config holds tournament service configuration
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

	Tournament struct {
		DefaultMinPlayers int  `yaml:"default_min_players"`
		DefaultMaxPlayers int  `yaml:"default_max_players"`
		DefaultStartDelay int  `yaml:"default_start_delay"`
		AutoStartEnabled  bool `yaml:"auto_start_enabled"`
		RebuyEnabled      bool `yaml:"rebuy_enabled"`
		AddonEnabled      bool `yaml:"addon_enabled"`
	} `yaml:"tournament"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}
