package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	DBInstance DBInstanceConfigurations
}

// EXported
type DBInstanceConfigurations struct {
	Name               string
	Id                 string
	StorageGB          string
	DBInstanceClass    string
	EngineVersion      string
	StorageType        string
	Region             string
	MultiAZ            bool
	PubliclyAccessible bool
	whiteList          string `mapstructure:"path_map"`
}

func main() {
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("DBInstanceName is\t", configuration.DBInstance.Name)
	fmt.Println("DBInstanceID is\t", configuration.DBInstance.Id)
	fmt.Println("Storage is\t", configuration.DBInstance.StorageGB)
	fmt.Println("DBInstanceClass is\t", configuration.DBInstance.DBInstanceClass)
	fmt.Println("EngineVersion is\t", configuration.DBInstance.EngineVersion)
	fmt.Println("MultiAZ is\t", configuration.DBInstance.MultiAZ)
	fmt.Println("PubliclyAccessible is\t", configuration.DBInstance.PubliclyAccessible)
	fmt.Println("StorageType is\t", configuration.DBInstance.StorageType)
	fmt.Println("Region is\t", configuration.DBInstance.Region)
	fmt.Println("Whitelist is\t", configuration.DBInstance.whiteList)
}
