package config

import "github.com/spf13/viper"

type Config struct {
	PORT                     string `mapstructure:"PORT"`
	NATS_CLUSTER             string `mapstructure:"NATS_CLUSTER"`
	USER_SERVICE_ADDRESS     string `mapstructure:"USER_SERVICE_ADDRESS"`
	PARTY_SERVICE_ADDRESS    string `mapstructure:"PARTY_SERVICE_ADDRESS"`
	STORY_SERVICE_ADDRESS    string `mapstructure:"STORY_SERVICE_ADDRESS"`
	RELATION_SERVICE_ADDRESS string `mapstructure:"RELATION_SERVICE_ADDRESS"`
	COMMENT_SERVICE_ADDRESS  string `mapstructure:"COMMENT_SERVICE_ADDRESS"`
	TOKEN_SECRET             string `mapstructure:"TOKEN_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
