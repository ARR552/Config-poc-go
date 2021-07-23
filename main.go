package main

import (
	"io/ioutil"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"github.com/imdario/mergo"
	"github.com/caarlos0/env/v6"
)
func main() {
	Init()
	fmt.Println("Default configuration: ",ConfigurationDef)

	// Set env variables. Done outside of the app
	os.Setenv("POC_VARIABLE", "DEFAULT_VARIABLE")
	os.Setenv("Option.Number", "5")

	data, err := LoadConf("./config.toml")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	fmt.Println("ConfigFile Configuration (ENV+TOML): ",data)
	mergo.Merge(data, ConfigurationDef)
	fmt.Println("Final Configuration (ENV+TOML+Defaul): ",data)
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
// Load loads a generic config.
func Load(path string, cfg interface{}) error {
	bs, err := ioutil.ReadFile(path) //nolint:gosec
	if err != nil {
		return err
	}
	cfgToml := string(bs)
	if _, err := toml.Decode(cfgToml, cfg); err != nil {
		return err
	}
	return nil
}
// LoadConf loads the configuration from path.
func LoadConf(path string) (*Configuration, error) {
	var cfg Configuration
	if err := Load(path, &cfg); err != nil {
		return nil, fmt.Errorf("error loading configuration file: %w", err)
	}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg, nil
}
var ConfigurationDef Configuration
func Init() {
	data := Data {
		Try: 1,
		URL: "Default_url",
	}
	var datas []Data
	datas = append(datas, data)
	datas = append(datas, data)
	option := Option {
		Number: 1,
	}
	conf := Conf {
		Datas: datas,
		Option: option,
	}
	config := Config {
		Op: "default_op",
		Extr: "default_extr",
	}
	ConfigurationDef = Configuration {
		Config: config,
		Conf: conf,
		POC_VARIABLE: "default_Poc_variable",
	}
}