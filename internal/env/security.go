package env

import (
	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/security"
)

func GetSecurityHashParams() *security.HashParams {
	return &security.HashParams{
		Memory:      uint32(GetIntegerEnv(consts.SecurityHashMemoryEnv)),
		Iterations:  uint32(GetIntegerEnv(consts.SecurityHashIterationsEnv)),
		Parallelism: uint8(GetIntegerEnv(consts.SecurityHashParallelismEnv)),
		SaltLength:  uint32(GetIntegerEnv(consts.SecurityHashSaltLengthEnv)),
		KeyLength:   uint32(GetIntegerEnv(consts.SecurityHashKeyLengthEnv)),
	}
}

func GetSecurityCryptoParams() *security.CryptoParams {
	return &security.CryptoParams{
		Secret: GetStringEnv(consts.SecurityEncryptionKeyEnv),
		Iv:     GetStringEnv(consts.SecurityEncryptionInitVectorEnv),
	}
}
