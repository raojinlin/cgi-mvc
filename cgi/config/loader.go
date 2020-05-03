package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Loader(filepath string) *Config {
	fs, err := os.OpenFile(filepath, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("open file error")
		panic(err)
	}

	content, err := ioutil.ReadAll(fs)
	if err != nil {
		fmt.Println("config file read error")
		panic(err)
	}

	var config Config = Config{}
	err = json.Unmarshal(content, &config)

	if err != nil {
		fmt.Println("config unmarshal error")
		panic(err)
	}

	return &config
}
