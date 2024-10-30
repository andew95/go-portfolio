package config

import "github.com/spf13/viper"

type Config struct {
	AppConfig  appConfig        `mapstructure:"app"`
	PostgreSQL postgreSQLConfig `mapstructure:"postgreSQL"`
}

type appConfig struct {
	SaveUploadPath string `mapstructure:"saveUploadPath"`
	Port           int    `mapstructure:"port"`
}

type postgreSQLConfig struct {
	BaseUrl  string `mapstructure:"baseUrl"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

func NewConfig() Config {
	// app config
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return Config{
		AppConfig: appConfig{
			SaveUploadPath: viper.GetString("app.saveUploadPath"),
			Port:           viper.GetInt("app.port"),
		},
		PostgreSQL: postgreSQLConfig{
			BaseUrl:  viper.GetString("postgreSQL.baseUrl"),
			Username: viper.GetString("postgreSQL.username"),
			Password: viper.GetString("postgreSQL.password"),
			DbName:   viper.GetString("postgreSQL.dbName"),
		},
	}
}
