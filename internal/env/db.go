package env

import "github.com/garnet-geo/garnet-userapi/internal/consts"

func GetDatabaseUrl() string {
	return GetStringEnv(consts.DatabaseUrlEnv)
}
