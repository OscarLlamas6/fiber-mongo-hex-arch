package settings

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		AppEnv                 string `mapstructure:"APP_ENV"`
		ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
		APIPort                string `mapstructure:"API_PORT"`
		DBHost                 string `mapstructure:"MONGODB_HOST"`
		DBPort                 string `mapstructure:"MONGODB_PORT"`
		DBUser                 string `mapstructure:"MONGODB_USER"`
		DBPass                 string `mapstructure:"MONGODB_PASS"`
		DBName                 string `mapstructure:"MONGODB_NAME"`
		DBCollection           string `mapstructure:"MONGODB_COLL"`
		AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
		RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
		AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
		RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	}
)

var (
	AppConfig   Config
	onceEnvLoad sync.Once
)

func ConvertToInt(stringNumber string) int {

	if stringNumber == "" {
		return 2
	}

	i, err := strconv.Atoi(stringNumber)
	if err != nil {
		log.Fatal(err)
		return 2
	}

	return i
}

func AskForEnv() {

	viper.AddConfigPath(".")

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if AppConfig.AppEnv == "development" || AppConfig.AppEnv == "dev" {
		log.Println("The App is running in development env")
	}
}

func SetConfig() {

	onceEnvLoad.Do(func() {
		if os.Getenv("IS_CONTAINER") != "TRUE" {
			AskForEnv()
		} else {
			AppConfig.AppEnv = os.Getenv("APP_ENV")
			AppConfig.ServerAddress = os.Getenv("SERVER_ADDRESS")
			AppConfig.APIPort = os.Getenv("API_PORT")
			AppConfig.DBHost = os.Getenv("MONGODB_HOST")
			AppConfig.DBPort = os.Getenv("MONGODB_PORT")
			AppConfig.DBUser = os.Getenv("MONGODB_USER")
			AppConfig.DBPass = os.Getenv("MONGODB_PASS")
			AppConfig.DBName = os.Getenv("MONGODB_NAME")
			AppConfig.DBCollection = os.Getenv("MONGODB_COLL")
			AppConfig.AccessTokenExpiryHour = ConvertToInt(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
			AppConfig.RefreshTokenExpiryHour = ConvertToInt(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))
			AppConfig.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
			AppConfig.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
		}

	})

}
