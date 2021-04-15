package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// ConfYaml ...
type ConfYaml struct {
	// Endpoints   SectionStorageEndpoint // `mapstructure:"endpoints"` // YAML
	Endpoints   []SectionStorageEndpoint // `mapstructure:"endpoints"`  // HCL
	DockerConfs []DockerConf             `mapstructure:"docker"`
}

// SectionStorageEndpoint ...
type SectionStorageEndpoint struct {
	Url       string // `mapstructure:"url"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	UseSSL    bool   `mapstructure:"use_ssl"` //defaults to false if not found!
	Location  string // `mapstructure:"location"`
}

// DockerConf ...
type DockerConf struct {
	Image string
	Tag   string
}

func main() {
	const format string = "hcl"
	// viper.SetConfigType("yaml")
	viper.SetConfigType(format)
	viper.SetConfigName("manifest")
	// viper.AddConfigPath("/etc/myapp/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	// error checking ...
	conf := &ConfYaml{}
	err = viper.Unmarshal(conf)
	if err == nil {
		fmt.Println(err)
	}

	/*
		if format == "yaml" {
			fmt.Println("conf", conf)
			fmt.Println("conf.Endpoints", conf.Endpoints)
			fmt.Println("conf.Endpoints.Url", conf.Endpoints.Url)
			fmt.Println("conf.EndpointsAccessKey", conf.Endpoints.AccessKey)
			url := viper.Get("endpoints.url")
			fmt.Println("Get url", url)
		}
	*/

	if format == "hcl" {
		for k, v := range conf.Endpoints {
			fmt.Println(k, v)
			fmt.Println(v.Url)
		}
	}
	for k, v := range conf.DockerConfs {
		fmt.Println(k, v)
		fmt.Println(v.Image)
	}

}
