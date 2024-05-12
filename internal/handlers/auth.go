package handlers

import (
	"errors"
	"fmt"

	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func AuthPostLogin(c *gin.Context) {
	var err error
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.AbortWithError(400, err)
		return
	}
	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(400, "Missing email or password")
		return
	}

	encryptedEmail := security.EncryptSymetric(loginData.Email, env.GetSecurityCryptoParams())
	row := db.ExecuteQueryRow(
		"SELECT id, domain, email, password FROM users "+
			"WHERE email = $1;",
		encryptedEmail,
	)

	var userData db.UserModel
	err = row.Scan(&userData.Id, &userData.Domain, &userData.Email, &userData.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, "User not found")
		return
	}
	if err != nil {
		c.JSON(500, fmt.Sprint(err))
		return
	}

	verified, err := security.VerifyHash(loginData.Password, userData.Password)
	if err != nil {
		c.JSON(500, fmt.Sprint(err))
		return
	}
	if !verified {
		c.JSON(401, "Invalid password")
		return
	}

	row = db.ExecuteQueryRow(
		"SELECT name, long_name FROM domains "+
			"WHERE id = $1;",
		userData.Domain,
	)

	var result UserInfoModel
	err = row.Scan(&result.Name, &result.LongName)
	if err != nil {
		c.JSON(500, fmt.Sprint(err))
		return
	}
	result.Id = UserId(userData.Id)
	decryptedEmail := security.DecryptSymetric(userData.Email, env.GetSecurityCryptoParams())
	result.Email = UserEmail(decryptedEmail)

	c.JSON(200, result)
}

func AuthPostUser(c *gin.Context) {

}

func AuthGetCheckToken(c *gin.Context) {

}
