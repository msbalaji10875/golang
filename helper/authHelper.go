package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Invalid userType")
		return err
	}
	return err
}

func CheckUserType1(ctx gin.Context, role string) (err error){
	userType := ctx.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorzed user")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("Unaithorized to access this resorce")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
