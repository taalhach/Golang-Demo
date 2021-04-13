package internal

import (
	"fmt"
	ini "github.com/nanitor/goini"
	"github.com/taalhach/Golang-Demo/internal/configs"
	"gorm.io/gorm"
	"os"
)

const (
	envKey = "VELOCITY_WORKS_SETTINGS"
)

var (
	MainConfigs *configs.MainConfig
	DB          *gorm.DB
)

func loadConfigs() (*configs.MainConfig, error) {

	file := os.Getenv(envKey)
	if file == "" {
		fmt.Printf("Missing env variable: %v", envKey)
		os.Exit(1)
	}

	dict, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	MainConfigs, err = configs.LoadMainConfig(dict)
	if err != nil {
		return nil, err
	}

	// make connection
	DB = MainConfigs.Database.MustGetDB()

	//var size string
	//DB.Raw("select pg_size_pretty(pg_database_size('velocity_works'))").Scan(&size)
	//fmt.Println(size)
	return MainConfigs, err
}
