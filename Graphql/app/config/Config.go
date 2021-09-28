package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	Db   string
	HOST string
)

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	///Users/guomiao/Desktop/go/API/Graphql/app/mysql
	viper.AddConfigPath("/Users/guomiao/Desktop/go/API/Graphql/conf")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	Db = viper.GetString("mysql")
	HOST = viper.GetString("host")
}
