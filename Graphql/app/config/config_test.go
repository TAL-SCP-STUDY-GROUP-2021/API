package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	workDir, _ := os.Getwd()
	//fmt.Printf("%s\n", workDir)
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/../../conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s", viper.GetString("mysql"))
}
