package env

import "github.com/garnet-geo/garnet-userapi/internal/consts"

func GetServerHttpPort() int {
	return GetIntegerEnv(consts.ServerHttpPortEnv)
}
