package consts

// Environment variables for hashing settings
const SecurityHashMemoryEnv = "GARNET_SECURITY_HASH_MEMORY"
const SecurityHashIterationsEnv = "GARNET_SECURITY_HASH_ITERATIONS"
const SecurityHashParallelismEnv = "GARNET_SECURITY_HASH_PARALLELISM"
const SecurityHashSaltLengthEnv = "GARNET_SECURITY_HASH_SALT_LENGTH"
const SecurityHashKeyLengthEnv = "GARNET_SECURITY_HASH_KEY_LENGTH"

// Environment variables for encryption settings
const SecurityEncryptionKeyEnv = "GARNET_SECURITY_ENCRYPTION_KEY"
const SecurityEncryptionInitVectorEnv = "GARNET_SECURITY_ENCRYPTION_IV"
const SecurityEncryptionTokenSecretEnv = "GARNET_SECURITY_ENCRYPTION_TOKEN_SECRET"

// Environment variables for server configuration
const ServerHttpPortEnv = "GARNET_SERVER_HTTP_PORT"

// Environment variables for database configuration
const DatabaseUrlEnv = "GARNET_DATABASE_URL"
