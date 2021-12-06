package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
	"path/filepath"
)

var (
	GlobalConfig     *Config
	GlobalServerOpts *ServerOpts
)

type ServerOpts struct {
	Config *Config
	DB     *sqlx.DB
}

type Config struct {
	Database *Database `ini:"database"`
}
type Database struct {
	Type    string `ini:"type"`
	Address string `ini:"address"`
}

func InitServerOpts() *ServerOpts {
	InitConfig()
	db, err := buildDB(GlobalConfig.Database.Type, GlobalConfig.Database.Address)
	if err != nil {
		panic(err)
	}
	return &ServerOpts{
		Config: GlobalConfig,
		DB:     db,
	}
}

func InitConfig() { //init  初始化
	config, err := parseConfig()
	if err != nil {
		panic(err)
	}
	GlobalConfig = config
}

func parseConfig() (*Config, error) { //parse  解析
	var config *Config
	filePath := "./env/local/config.ini"
	filePaths, err := filepath.Abs(filePath)
	fmt.Println(filePaths, err)
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, err
	}
	err = cfg.MapTo(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
