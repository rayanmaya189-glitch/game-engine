package room

// Config holds chat service configuration
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

	Chat struct {
		MaxMessageLength     int `yaml:"max_message_length"`
		MaxMessagesPerRoom   int `yaml:"max_messages_per_room"`
		MessageRetentionDays int `yaml:"message_retention_days"`
		RateLimitPerMinute   int `yaml:"rate_limit_per_minute"`
	} `yaml:"chat"`

	ProfanityFilter struct {
		Enabled         bool   `yaml:"enabled"`
		ReplacementChar string `yaml:"replacement_char"`
		FilterLevel     string `yaml:"filter_level"`
	} `yaml:"profanity_filter"`

	Moderation struct {
		AutoMuteThreshold   int  `yaml:"auto_mute_threshold"`
		MuteDurationMinutes int  `yaml:"mute_duration_minutes"`
		BanDurationHours    int  `yaml:"ban_duration_hours"`
		RequiresModerator   bool `yaml:"requires_moderator"`
	} `yaml:"moderation"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}
