package config

import "github.com/spf13/viper"

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func Parse(conf any) error {
	return viper.Unmarshal(conf)
}
