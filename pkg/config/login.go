package config

import "github.com/spf13/viper"

type Credentials struct {
	Username string
	Password string
}

func GetCredentials() (c Credentials) {
	viper.UnmarshalKey("credentials", &c)
	return c
}

func SetCredentials(c Credentials) error {
	viper.Set("credentials", c)

	return viper.WriteConfig()
}
