package system

import (
	"context"
	b64 "encoding/base64"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type contextKey string

const (
	ClaimsKey       contextKey = "claims"
	ActiveUserKey   contextKey = "activeUser"
	LoggedInUserKey contextKey = "loggedInUser"
)

var (
	Path, _      = os.Getwd()
	AgaContext   context.Context
	AllowOrigins = configEnv("aga.AllowOrigins")
	AllowHeaders = configEnv("aga.AllowHeaders")
	Port         = configEnv("aga.Port")
	Language     = configEnv("aga.Language")

	// -- MYSQL
	mysqlPassB64, _  = b64.StdEncoding.DecodeString(configEnv("aga.mysql.Password"))
	mysqlPassword    = string(mysqlPassB64)
	mysqlHost        = configEnv("aga.mysql.Host")
	mysqlPort        = configEnv("aga.mysql.Port")
	mysqlUser        = configEnv("aga.mysql.Username")
	mysqlDB          = configEnv("aga.mysql.DB")
	MysqlCredentials = mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8&parseTime=True&loc=Local"

	DomainAGA    = configEnv("aga.Domain")
	SecretKeyAGA = configEnv("SecretKeyAGA")
	UserKeyAGA   = configEnv("UserKeyAGA")
	AdminUserId  = configEnv("aga.AdminUserId")

	PortServerRPC = configEnv("aga.rpc.Port")

	Log    = make(chan interface{})
	Notify = make(chan interface{})

	OrcIdClientID                  = configEnv("aga.OrcId.ClientID")
	OrcIdClientSecret              = configEnv("aga.OrcId.ClientSecret")
	PasswordExpirationThreshold, _ = strconv.Atoi(configEnv("aga.PasswordExpirationThreshold"))
  AssetsSigenuApiUrl             = configEnv("aga.AssetsSigenu.ApiURL")

	// -- EMAIL
	EmailName     = configEnv("aga.mail.Name")
	EmailUser     = configEnv("aga.mail.Username")
	EmailPassword = configEnv("aga.mail.Password")

)

func configEnv(data string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	return os.Getenv(data)
}
