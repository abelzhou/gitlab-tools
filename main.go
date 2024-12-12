//
//author:abel
//date:2023/9/1
package main

import (
	"github.com/spf13/viper"
	"gt/cmd"
	"gt/pkg/gitlab"
	"log"
)

func init() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".gitlabtools")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Configuration file not found")
		} else {
			log.Println("Error reading configuration file")
		}
		log.Fatal(err) // Fatal error: failed to read configuration file
	}
}
func main() {
	gitlab.InitGitlabClient()
	cmd.Execute()
}
