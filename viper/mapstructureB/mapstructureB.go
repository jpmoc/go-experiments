package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type DockerConfig struct {
	Image string // image
	Tag   string // tag
}
type QueryConfig struct {
	// Top most field contains a dictionary where keys map to an array of CountryConfigS
	CountryQueries map[string][]DockerConfig
}

func NewQueryConfig() QueryConfig {
	queryConfig := QueryConfig{}
	queryConfig.CountryQueries = map[string][]DockerConfig{}
	return queryConfig
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	queryConfig := NewQueryConfig()
	if err != nil {
		log.Panic("error:", err)
	} else {
		// mapstructure.Decode(viper.AllSettings(), &queryConfig)
		fmt.Println(viper.Get("toto"))
		mapstructure.Decode(viper.Get("toto"), &queryConfig)
	}
	for _, config := range queryConfig.CountryQueries["sg"] {

		fmt.Println("qtype:", config.Image, "qplacetype:", config.Tag)
	}
}
