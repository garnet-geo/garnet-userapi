package handlers

import (
	"errors"

	// "github.com/garnet-geo/garnet-userapi/internal/consts"
	"github.com/garnet-geo/garnet-userapi/internal/db"
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
	userInfo.Name = DomainName(domainId)

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

func UserInfoPatchUserById(c *gin.Context) {}

func UserInfoDeleteUserById(c *gin.Context) {}
