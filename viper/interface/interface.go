package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	testvar := viper.Get("element")
	fmt.Println(testvar)
	// array on anything
	elementsMap := testvar.([]interface{})
	// loop because array!
	for k, vmap := range elementsMap {
		fmt.Print("Key: ", k)
		fmt.Println(" Value: ", vmap)
		eachElementsMap := vmap.(map[interface{}]interface{})

		for k, vEachValMap := range eachElementsMap {
			fmt.Printf("%v: %v \n", k, vEachValMap)
			vEachValDataMap := vEachValMap.(map[interface{}]interface{})

			for k, v := range vEachValDataMap {
				fmt.Printf("%v: %v \n", k, v)
			}
		}
	}
}
