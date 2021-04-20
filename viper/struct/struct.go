package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// ConfYaml ...
type ConfYaml struct {
	// Endpoints   SectionStorageEndpoint // `mapstructure:"endpoints"` // YAML
	Endpoints []SectionStorageEndpoint // `mapstructure:"endpoints"`  // HCL
	// DockerConfs []DockerConf             `mapstructure:"docker"`
	Uses []Use `mapstructure:"use"`
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

// Use ...
type Use struct {
	DockerConfs   []DockerConf   `mapstructure:"docker"`
	SkaffoldConfs []SkaffoldConf `mapstructure:"skaffold"`
}

// SkaffoldConf ...
type SkaffoldConf struct {
	Image  string
	Tag    string
	Kaniko bool
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

	// set default config
	//viper.SetDefault("ContentDir", "content")
	viper.SetDefault("ContentDir", "0")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomy", map[string]string{"tag": "tags", "category": "categories"})
	viper.SetDefault("Taxonomies", []map[string]string{{"tag": "tags", "category": "categories"}})

	fmt.Println(viper.GetBool("ContentDir"))
	fmt.Println(viper.GetString("LayoutDir"))
	fmt.Println(viper.GetStringMapString("Taxonomy"))
	fmt.Println(viper.GetStringMapString("Taxonomy")["category"])
	fmt.Println(viper.Get("Taxonomies").([]map[string]string)[0])
	fmt.Println(viper.Get("Taxonomies").([]map[string]string)[0]["category"])

	fmt.Println("#")
	fmt.Println(viper.AllSettings())
	fmt.Println("Get", viper.Get("endpoints"))
	// fmt.Println(viper.Get("endpoints")[0]) // interface does not support indexing
	fmt.Println("Get+interface1", viper.Get("endpoints").([]map[string]interface{})[0])
	fmt.Println("Get+interface2", viper.Get("endpoints").([]map[string]interface{})[0])
	fmt.Println("Get+interface3", viper.Get("endpoints").([]map[string]interface{})[0]["location"])
	fmt.Println("Get+interface3", viper.Get("endpoints").([]map[string]interface{})[0]["middle"])
	fmt.Println("Get+interface3", viper.Get("endpoints").([]map[string]interface{})[0]["middle"])
	fmt.Println("Get+interface4", viper.Get("endpoints").([]map[string]interface{})[0]["middle"].([]map[string]interface{})[0])
	fmt.Println("Get+interface5", viper.Get("endpoints").([]map[string]interface{})[0]["middle"].([]map[string]interface{})[0]["key02"])
	fmt.Println("#")
	fmt.Println("GetStringMap", viper.GetStringMap("endpoints"))
	fmt.Println(viper.GetStringMapString("endpoints"))
	fmt.Println(viper.GetStringSlice("endpoints"))
	fmt.Println("#")

	conf := &ConfYaml{}
	err = viper.Unmarshal(conf)
	if err != nil {
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
	/*
		for k, v := range conf.DockerConfs {
			fmt.Println(k, v)
			fmt.Println(v.Image)
		}
	*/

	for i, use := range conf.Uses {
		fmt.Println(i, use)
		for j, dockerConf := range use.DockerConfs {
			fmt.Println(j, dockerConf.Image)
		}
		for j, skaffoldConf := range use.SkaffoldConfs {
			fmt.Println(j, skaffoldConf.Kaniko)
		}

	}

}
