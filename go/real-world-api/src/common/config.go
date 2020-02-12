package common

import "os"

// Config godoc
var Config = map[string]string{}

// InitConfig godoc
func InitConfig() {
	Config["MYSQL_USERNAME"] = os.Getenv("MYSQL_USERNAME")
	Config["MYSQL_PASSWORD"] = os.Getenv("MYSQL_PASSWORD")
	Config["PORT"] = os.Getenv("PORT")
	Config["SECRET"] = os.Getenv("SECRET")
}
