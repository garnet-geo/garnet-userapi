package handlers

import (
	"errors"

	"github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/db"
	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/garnet-geo/garnet-userapi/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func UserInfoGetUserById(c *gin.Context) {
	// All user profiles are public currently
	// userID := c.MustGet(consts.UserIDContextKey).(string)
	requestedUserId := c.Param("id")
	if requestedUserId == "" {
		c.JSON(400, "Invalid user ID")
		return
	}

	userInfo := UserInfoModel{
		Id: UserId(requestedUserId),
	}

	row := db.ExecuteQueryRow(
		"SELECT domain, email FROM users "+
			"WHERE id = $1;",
		requestedUserId,
	)
	var domainId string
	err := row.Scan(&domainId, &userInfo.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, "User not found")
	} else if err != nil {
		c.JSON(500, err)
		return
	}
	decryptedEmail := security.DecryptSymetric(string(userInfo.Email), env.GetSecurityCryptoParams())
	userInfo.Email = UserEmail(decryptedEmail)

	row = db.ExecuteQueryRow(
		"SELECT name, long_name FROM domains "+
			"WHERE id = $1;",
		domainId,
	)
	err = row.Scan(&userInfo.Name, &userInfo.LongName)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, userInfo)
}

func UserInfoPatchUserById(c *gin.Context) {
	// All user profiles are public currently
	userID := c.MustGet(consts.UserIDContextKey).(string)
	requestedUserId := c.Param("id")
	if requestedUserId == "" {
		c.JSON(400, "Invalid user ID")
		return
	}

	if userID != requestedUserId {
		c.JSON(403, "Forbidden")
		return
	}

	var err error
	var userEditData struct {
		Name     string `json:"name"`
		LongName string `json:"long_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err = c.BindJSON(&userEditData); err != nil {
		c.AbortWithError(400, err)
		return
	}

	if userEditData.Name == "" &&
		userEditData.LongName == "" &&
		userEditData.Email == "" &&
		userEditData.Password == "" {
		c.JSON(400, "Nothing to update")
		return
	}

	userInfo := UserInfoModel{
		Id: UserId(requestedUserId),
	}

	row := db.ExecuteQueryRow(
		"SELECT domain, email, password FROM users "+
			"WHERE id = $1;",
		requestedUserId,
	)
	var domainId string
	err = row.Scan(&domainId, &userInfo.Email, &userInfo.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(404, "User not found")
	} else if err != nil {
		c.JSON(500, err)
		return
	}

	row = db.ExecuteQueryRow(
		"SELECT name, long_name FROM domains "+
			"WHERE id = $1;",
		domainId,
	)
	err = row.Scan(&userInfo.Name, &userInfo.LongName)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if userEditData.Name == "" {
		userEditData.Name = string(userInfo.Name)
	}
	if userEditData.LongName == "" {
		userEditData.LongName = string(userInfo.LongName)
	}
	if userEditData.Email == "" {
		userEditData.Email = string(userInfo.Email)
	} else {
		encryptedEmail := security.EncryptSymetric(userEditData.Email, env.GetSecurityCryptoParams())
		userEditData.Email = encryptedEmail
	}
	if userEditData.Password == "" {
		userEditData.Password = string(userInfo.Password)
	} else {
		hashedPassword, err := security.CreateHash(userEditData.Password, env.GetSecurityHashParams())
		if err != nil {
			c.JSON(500, err)
			return
		}
		userEditData.Password = hashedPassword
	}

	transaction, err := db.ExecuteTransaction()
	if err != nil {
		c.JSON(500, err)
		return
	}

	transaction.Exec(db.Context(),
		"UPDATE users "+
			"SET email = $2, password = $3 "+
			"WHERE id = $1;",
		requestedUserId,
		userEditData.Email, userEditData.Password,
	)
	transaction.Exec(db.Context(),
		"UPDATE domains "+
			"SET name = $2, long_name = $3 "+
			"WHERE id = $1;",
		domainId,
		userEditData.Name, userEditData.LongName,
	)

	err = transaction.Commit(db.Context())
	if err != nil {
		c.JSON(500, err)
		return
	}

	decryptedEmail := security.DecryptSymetric(userEditData.Email, env.GetSecurityCryptoParams())
	userInfo.Email = UserEmail(decryptedEmail)
	userInfo.Password = ""
	userInfo.Name = DomainName(userEditData.Name)
	userInfo.LongName = userEditData.LongName

	c.JSON(200, userInfo)
}

func UserInfoDeleteUserById(c *gin.Context) {}
