package config

import (
	"github.com/bahodurnazarov/wallet/pkg/models"
	"github.com/spf13/viper"
	"log"
)

const (
	configName = "conf"
	configType = "json"
	configPath = "config"
)

var Cnfg *models.Settings

func Init() (*models.Settings, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath("pkg/" + configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error in reading configs %w", err)
		return nil, err
	}
	var cfg *models.Settings
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	Cnfg = cfg
	return Cnfg, nil
}
