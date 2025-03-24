package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type PairConfig struct {
	Alias  string         `json:"alias"`
	Source DatabaseConfig `json:"source"`
	Target DatabaseConfig `json:"target"`
}

type Config struct {
	Pairs []PairConfig `json:"pairs"`
}

func LoadConfig() (Config, error) {
	const filename = "db.json"
	var config Config

	file, err := os.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("ошибка чтения файла конфигурации: %v", err)
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return config, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	if len(config.Pairs) == 0 {
		return config, fmt.Errorf("в файле конфигурации не найдено ни одной пары баз данных")
	}

	return config, nil
}
