package config

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/redis"
	"encoding/json"
)

const Key = "datatable:config"

func Get() (models.Config, error) {
	var config models.Config
	var err error

	res, err := redis.Exec("get", Key)
	if err != nil {
		return config, err
	}
	if res != nil {
		configStr, err := redis.ToString(res, err)
		if err != nil {
			return config, err
		}
		err = json.Unmarshal([]byte(configStr), &config)
		return config, err
	}
	config, err = new(services.ConfigService).GetOne()
	if err != nil {
		return config, err
	}
	return config, Set(config)
}

//设置config值
func Set(config models.Config) error {
	jsonByte, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return redis.Set(Key, string(jsonByte), 0)
}
