package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Init Config Variables
func Init() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}
	fmt.Println(viper.GetString("reddit.secret"))
}
