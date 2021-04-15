package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type CountryConfig struct {
	Qtype      string
	Qplacetype string
}
type QueryConfig struct {
	CountryQueries map[string][]CountryConfig
}

func NewQueryConfig() QueryConfig {
	queryConfig := QueryConfig{}
	queryConfig.CountryQueries = map[string][]CountryConfig{}
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
		mapstructure.Decode(viper.AllSettings(), &queryConfig)
	}
	for _, config := range queryConfig.CountryQueries["sg"] {

		fmt.Println("qtype:", config.Qtype, "qplacetype:", config.Qplacetype)
	}
}
