package handlers 

import "real-world-api/src/common"

func checkPassword(pass string, modelPass string) bool {
	return common.Hash(pass) == modelPass
}