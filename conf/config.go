package conf

import (
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

var (
	confFile string
	Conf     Config
)

func init() {
	flag.StringVar(&confFile, "configfile", "./conf/config.toml", " set im_proxy config file path")
}

type Config struct {
	Server      Serverinfo        `toml:"serverinfo"`
	Worker      WorkerInfo        `toml:"worker"`
	Application ApplicationInfo   `toml:"application"`
}

type Serverinfo struct {
	Bind    string
	Version int
}

type WorkerInfo struct {
	JobSize    int `toml:"job_size"`
	WorkerSize int `toml:"worker_size"`
}

type ApplicationInfo struct {
	GinMode           string        `toml:"gin_mode"`
}

func InitConfig() error {
	if _, err := toml.DecodeFile(confFile, &Conf); err != nil {
		return err
	}
	b, err := json.Marshal(&Conf)
	if err != nil {
		return err
	}
	glog.Infof("配置文件信息：%s", string(b))
	return nil
}
