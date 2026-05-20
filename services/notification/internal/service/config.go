package service

type Config struct {
	Server struct {
		GRPCPort int `yaml:"grpc_port"`
	} `yaml:"server"`
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`
	Notification struct {
		Push struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"push"`
		Email struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"email"`
		SMS struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"sms"`
		InApp struct {
			Enabled       bool `yaml:"enabled"`
			MaxPerUser    int  `yaml:"max_per_user"`
			RetentionDays int  `yaml:"retention_days"`
		} `yaml:"in_app"`
		Batch struct {
			Enabled      bool `yaml:"enabled"`
			MaxBatchSize int  `yaml:"max_batch_size"`
		} `yaml:"batch"`
	} `yaml:"notification"`
}
