package config

type (
	AppConfig struct {
		GRPC         GRPCConfig         `mapstructure:"grpc"`
		FeatureRedis FeatureRedisConfig `mapstructure:"feature_redis"`
	}

	GRPCConfig struct {
		Port int `mapstructure:"address"`
	}

	FeatureRedisConfig struct {
		Addr string `mapstructure:"address"`
		DB   int    `mapstructure:"db"`
	}
)
