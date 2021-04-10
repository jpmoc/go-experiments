package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Conf struct {
	NATS []struct {
		HTTPPort int
		Port     int
		Username string
		Password string
	}
}

func main() {
	var c Conf
	// config file
	viper.SetConfigName("draft")
	viper.AddConfigPath(".")
	viper.SetConfigType("hcl")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.Get("NATS")) // gives [map[port:10041 username:cl1 password:__Psw__4433__ http_port:10044]]

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.NATS[0].Username)
}
