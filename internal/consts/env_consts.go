package consts

const GlobalDebugEnv = "GARNET_DEBUG"

// Environment variables for hashing settings
const (
	SecurityHashMemoryEnv      = "GARNET_SECURITY_HASH_MEMORY"
	SecurityHashIterationsEnv  = "GARNET_SECURITY_HASH_ITERATIONS"
	SecurityHashParallelismEnv = "GARNET_SECURITY_HASH_PARALLELISM"
	SecurityHashSaltLengthEnv  = "GARNET_SECURITY_HASH_SALT_LENGTH"
	SecurityHashKeyLengthEnv   = "GARNET_SECURITY_HASH_KEY_LENGTH"
	SecurityHashSpecialSalt    = "GARNET_SECURITY_HASH_SPECIAL_SALT"
)

// Environment variables for encryption settings
const (
	SecurityEncryptionKeyEnv         = "GARNET_SECURITY_ENCRYPTION_KEY"
	SecurityEncryptionInitVectorEnv  = "GARNET_SECURITY_ENCRYPTION_IV"
	SecurityEncryptionTokenSecretEnv = "GARNET_SECURITY_ENCRYPTION_TOKEN_SECRET"
)

// Environment variables for server configuration
const ServerHttpPortEnv = "GARNET_SERVER_HTTP_PORT"

// Environment variables for database configuration
const DatabaseUrlEnv = "GARNET_DATABASE_URL"
