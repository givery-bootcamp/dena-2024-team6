package config

import (
	"os"
	"strconv"
)

var HostName = "localhost"
var Port = 9000
var AppEnv = "prod"
var CorsAllowOrigin = "http://localhost:3000"
var DBHostName = "db"
var DBPort = 3306
var DBName = "training"
var DefaultTimeoutSecond = 3
var DBPassword = ""
var DBUser = "root"
var JwtKey = "my_secret_key"

func init() {
	if v := os.Getenv("HOSTNAME"); v != "" {
		HostName = v
	}
	if v, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64); err == nil {
		Port = int(v)
	}
	if v := os.Getenv("CORS_ALLOW_ORIGIN"); v != "" {
		CorsAllowOrigin = v
	}
	if v := os.Getenv("DB_HOSTNAME"); v != "" {
		DBHostName = v
	}
	if v, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64); err == nil {
		DBPort = int(v)
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DBName = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		DBPassword = v
	}
	if v := os.Getenv("DB_USERNAME"); v != "" {
		DBUser = v
	}
	if v := os.Getenv("JWT_KEY"); v != "" {
		JwtKey = v
	}

	// NOTE: 本番環境でEnvを追加するのがめんどいので、CORSのオリジンでDevかProdか判別する
	if CorsAllowOrigin == "http://localhost:3000" {
		AppEnv = "dev"
	}

}

// GetDomainName はAPIのドメインネームを返す
func GetDomainName() string {
	if AppEnv == "prod" {
		return HostName
	}
	return "localhost"
}

func GetIsSecured() bool {
	return AppEnv == "prod"
}
