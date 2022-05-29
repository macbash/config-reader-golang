package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func configread(keyname string) []string {
	var resArray []string
	viper.AddConfigPath("./")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()
	str := fmt.Sprintf("%v", viper.Get(keyname))
	for _, word := range strings.Split(str, ",") {
		resArray = append(resArray, word)
	}
	return resArray
}
func main() {
	excl := configread("test.exclude")
	fmt.Println(excl)
	name := configread("test.name")
	fmt.Println(name)
	fmt.Println(configread("test.env"))
	fmt.Println(configread("prod.env"))
	fmt.Println(configread("prod.exclude"))
}
