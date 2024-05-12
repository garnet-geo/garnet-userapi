package main

import (
	"fmt"
	"log"

	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/garnet-geo/garnet-userapi/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Starting database connection
	db.InitConnection()
	defer db.CloseConnection()

	server.InitServer()

	return

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	hash, err := security.CreateHash("password123", env.GetSecurityHashParams())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)

	verified, err := security.VerifyHash("password123", hash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Verified: ", verified)

	fmt.Println("-------")

	encryptedEmail := security.EncryptSymetric("u.slash.vlad@gmail.com", env.GetSecurityCryptoParams())
	fmt.Println(encryptedEmail)

	decryptedEmail := security.DecryptSymetric(encryptedEmail, env.GetSecurityCryptoParams())
	fmt.Println(decryptedEmail)
}
