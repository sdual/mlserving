package config

type (
	AppConfig struct {
		GRPC         GRPC         `mapstructure:"grpc"`
		Logger       Logger       `mapstructure:"logger"`
		FeatureRedis FeatureRedis `mapstructure:"feature_redis"`
	}

	GRPC struct {
		Port int `mapstructure:"port"`
	}

	Logger struct {
		LogLevel string `mapstructure:"log_level"`
	}

	FeatureRedis struct {
		Addr string `mapstructure:"address"`
		DB   int    `mapstructure:"db"`
	}
)
