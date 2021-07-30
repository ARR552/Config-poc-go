package main

import (
	"fmt"
	"os"
	"strconv"
	configLibrary "github.com/hermeznetwork/go-hermez-config"

)
func main() {
	ConfigurationDef := `
Miss = 30
POC_VARIABLE = "Hola_default"
[Conf.Option]
Number = 101
[[Conf.Datas]]
URL = "default_url"
Try = 50
[Config]
Op = "default"
Extr = "yes_default"
POC_VARIABLE2 = "File_default"
`

	// Set env variables. Done outside of the app
	os.Setenv("POC_VARIABLE", "DEFAULT_VARIABLE")
	os.Setenv("Option.Number", "5")
	os.Setenv("Data.URL", "http://localhost:3000/api")
	os.Setenv("Data.Try", "1000")

	path := "./config.toml"
    
	var cfg Configuration
	err := configLibrary.LoadConfig(path, ConfigurationDef, &cfg)
	if err != nil {
        //Handle error
        fmt.Println(err)
    }
	var data Data
	if os.Getenv("Data.URL") != "" {
		data = Data {
			URL: os.Getenv("Data.URL"),
		}
		if os.Getenv("Data.Try") != "" {
			i, err := strconv.Atoi(os.Getenv("Data.Try"))
			if err != nil {
				fmt.Println(err)
			}
			data.Try = i
		}
		var datas []Data
		datas = append(datas, data)
		cfg.Conf.Datas = datas
	}
    fmt.Println("Configuration: ", cfg)
}
type Data struct {
	URL string `env:"Data.URL"`
	Try int    `env:"Data.Try"`
}
type Option struct {
	Number int  `env:"Option.Number"`
}
type Conf struct {
	Option Option
	Datas []Data
}
type Config struct {
	Op string   `env:"Config.Op"`
	Extr string  `env:"Config.Extr"`
}
type Configuration struct {
	Config Config
	Conf Conf
	POC_VARIABLE string  `env:"POC_VARIABLE"`
}