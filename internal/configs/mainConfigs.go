package configs

import (
	"fmt"
	ini "github.com/nanitor/goini"
)

type MainConfig struct {
	TemplatesDirectory string
	Database           *DatabaseConfig
}

func LoadMainConfig(dict ini.Dict) (*MainConfig, error) {
	var err error

	ret := &MainConfig{}

	ret.TemplatesDirectory = dict.GetStringDef("main", "templates_directory", "")
	if ret.TemplatesDirectory == "" {
		return nil, fmt.Errorf("templates_directory is empty")
	}

	ret.Database, err = DatabaseConfigsFromDict(dict)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
