//
//author:abel
//date:2023/9/1
package main

import (
	"github.com/spf13/viper"
	"gt/cmd"
	"log"
)

func init() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".gitlabtools")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("没有找到配置文件")
		} else {
			log.Println("配置文件读取错误")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}
func main() {
	cmd.Execute()
}
