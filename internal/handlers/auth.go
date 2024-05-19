package handlers

import (
	"errors"

	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func AuthPostLogin(c *gin.Context) {
	var err error
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err = c.BindJSON(&loginData); err != nil {
		c.AbortWithError(400, err)
		return
	}

	log.Traceln("login - Got login data", loginData)

	hashedEmail, err := security.CreateHash(
		loginData.Email,
		env.GetSecurityHashParams(),
		env.GetSecurityHashSpecialSalt())
	if err != nil {
		c.JSON(500, err)
		return
	}
	row := db.ExecuteQueryRow(
		"SELECT id, password FROM users "+
			"WHERE email_hash = $1;",
		hashedEmail,
	)

	var userId string
	var userPassword string
	err = row.Scan(&userId, &userPassword)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, "User not found")
		return
	} else if err != nil {
		c.JSON(500, err)
		return
	}

	log.Traceln("login - Got user", userId)

	verified, err := security.VerifyHash(loginData.Password, userPassword)
	if err != nil {
		c.JSON(500, err)
		return
	}
	if !verified {
		c.JSON(401, "Invalid password")
		return
	}

	log.Traceln("login - Verified user password", userId)

	token := security.GenTokenForUser(string(userId), env.GetSecurityTokenSecret())

	log.Traceln("login - Created token for user", userId)

	c.JSON(200, gin.H{
		"token": token,
	})
}

func AuthPostUser(c *gin.Context) {
	var err error
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
		LongName string `json:"long_name"`
	}

	if err = c.BindJSON(&loginData); err != nil {
		c.AbortWithError(400, err)
		return
	}

	log.Traceln("registration - Got user data", loginData)

	if len(loginData.Password) < 4 {
		c.JSON(400, "Password must be at least 4 characters long")
		return
	}
	if len(loginData.Name) < 5 {
		c.JSON(400, "Name must be at least 5 characters long")
		return
	}

	log.Traceln("registration - Validated user data")

	hashedEmail, err := security.CreateHash(
		loginData.Email,
		env.GetSecurityHashParams(),
		env.GetSecurityHashSpecialSalt())
	if err != nil {
		c.JSON(500, err)
		return
	}
	row := db.ExecuteQueryRow(
		"SELECT COUNT(id) FROM users "+
			"WHERE email_hash = $1;",
		hashedEmail,
	)
	var count int
	err = row.Scan(&count)
	if err != nil {
		c.JSON(500, err)
		return
	}

	log.Traceln("registration - Users with same email:", count)

	if count > 0 {
		c.JSON(403, "User already exists")
		return
	}

	log.Traceln("registration - Checked user existence")

	// TODO(uSlashVlad): Make it in transaction
	row = db.ExecuteQueryRow(
		"INSERT INTO domains ( name, long_name ) "+
			"VALUES ( $1, $2 ) "+
			"RETURNING id;",
		loginData.Name,
		loginData.LongName,
	)

	var domainId string
	err = row.Scan(&domainId)
	if err != nil {
		c.JSON(500, err)
		return
	}

	log.Traceln("registration - Inserted domain", domainId)

	passwordHash, err := security.CreateHash(loginData.Password, env.GetSecurityHashParams())
	if err != nil {
		c.JSON(500, err)
		return
	}
	log.Traceln("registration - Hashed password")
	encryptedEmail := security.EncryptSymetric(loginData.Email, env.GetSecurityCryptoParams())
	log.Traceln("registration - Encrypted email")

	row = db.ExecuteQueryRow(
		"INSERT INTO users ( domain, email, email_hash, password ) "+
			"VALUES ( $1, $2, $3, $4 ) "+
			"RETURNING id;",
		domainId,
		encryptedEmail,
		hashedEmail,
		passwordHash,
	)

	var userId string
	err = row.Scan(&userId)
	if err != nil {
		c.JSON(500, err)
		return
	}

	log.Traceln("registration - Inserted user", userId)

	token := security.GenTokenForUser(string(userId), env.GetSecurityTokenSecret())

	log.Traceln("registration - Got token for user", userId)

	c.JSON(200, gin.H{
		"token": token,
		"user": UserInfoModel{
			Id:       UserId(userId),
			Name:     DomainName(loginData.Name),
			LongName: loginData.LongName,
			Email:    UserEmail(loginData.Email),
		},
	})
}

func AuthGetCheckToken(c *gin.Context) {
	userID := c.MustGet(consts.UserIDContextKey).(string)

	c.JSON(200, gin.H{
		"user_id": userID,
	})
}
