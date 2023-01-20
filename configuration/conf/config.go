package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var shared *_Configuration

type _Configuration struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeOut  time.Duration `json:"read_timeout"`
		WriteTimeOut time.Duration `json:"write_timeout"`
	} `json:"server"`

	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
}

func init() {
	if shared != nil {
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	read, err := ioutil.ReadFile(filepath.Join(dir, "conf", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	shared = new(_Configuration)
	err = json.Unmarshal(read, &shared)
	if err != nil {
		panic(err)
		return
	}

}

func Configuration() _Configuration {
	return *shared
}
