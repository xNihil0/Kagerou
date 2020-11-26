package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mtdn.io/Kagerou/database"
	"mtdn.io/Kagerou/internal"
	"net/http"
	"strconv"
)

const (
	InvalidType   = 40001
	ExistingUser  = 40002
	UserNotFound  = 40401
	DatabaseError = 50001
)

func CreateUser(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.PostForm("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	i, err = strconv.ParseInt(c.PostForm("u2_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid u2 id", InvalidType, nil, 1)
	}
	U2ID := int32(i)
	inputUser := &database.InputUser{
		TelegramID: TelegramID,
		U2ID:       U2ID,
	}
	_, err = inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		}
	} else {
		return MakeErrorReturn(http.StatusConflict, "existing user", ExistingUser, nil, 1)
	}
	//goland:noinspection GoNilness
	err = inputUser.CreateUser()
	if err != nil {
		return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
	}
	user, err := inputUser.GetUser()
	if err != nil {
		return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
	}
	return MakeSuccessReturn(database.ToOutputUser(user))
}

func GetUser(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.Param("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	inputUser := database.InputUser{
		TelegramID: TelegramID,
	}
	user, err := inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		} else {
			return MakeErrorReturn(http.StatusNotFound, "user not found", UserNotFound, nil, 1)
		}
	}
	return MakeSuccessReturn(database.ToOutputUser(user))
}

func VerifyUser(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.Param("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	inputUser := database.InputUser{
		TelegramID: TelegramID,
	}
	user, err := inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		} else {
			return MakeErrorReturn(http.StatusNotFound, "user not found", UserNotFound, nil, 1)
		}
	}
	if user.Verified == true {
		return MakeErrorReturn(http.StatusOK, "user already verified", 20000, database.ToOutputUser(user), 1)
	}
	if internal.CheckVerificationCode(user.U2ID, user.VerificationCode) == false {
		return MakeErrorReturn(http.StatusOK, "verification failed", 20001, database.ToOutputUser(user), 1)
	}
	user.Verified = true
	user.UpdateUser()
	return MakeSuccessReturn(database.ToOutputUser(user))
}

func UpdateUser(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.PostForm("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	i, err = strconv.ParseInt(c.PostForm("u2_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid u2 id", InvalidType, nil, 1)
	}
	U2ID := int32(i)
	inputUser := &database.InputUser{
		TelegramID: TelegramID,
		U2ID:       U2ID,
	}
	user, err := inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		} else {
			return MakeErrorReturn(http.StatusNotFound, "user not found", UserNotFound, nil, 1)
		}
	}
	user.U2ID = U2ID
	user.VerificationCode = internal.GenerateVerificationCode()
	user.Verified = false
	err = user.UpdateUser()
	if err != nil {
		return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
	}
	return MakeSuccessReturn(database.ToOutputUser(user))
}

func RemoveUser(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.Param("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	inputUser := database.InputUser{
		TelegramID: TelegramID,
	}
	user, err := inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		} else {
			return MakeErrorReturn(http.StatusNotFound, "user not found", UserNotFound, nil, 1)
		}
	}
	err = user.RemoveUser()
	if err != nil {
		return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
	}
	return MakeSuccessReturn(nil)
}

func ResetUserVerificationCode(c *gin.Context) (int, interface{}) {
	var i int64
	var err error
	i, err = strconv.ParseInt(c.Param("telegram_id"), 10, 32)
	if err != nil {
		return MakeErrorReturn(http.StatusBadRequest, "invalid telegram id", InvalidType, nil, 1)
	}
	TelegramID := int32(i)
	inputUser := database.InputUser{
		TelegramID: TelegramID,
	}
	user, err := inputUser.GetUser()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return MakeErrorReturn(http.StatusInternalServerError, err.Error(), DatabaseError, nil, 1)
		} else {
			return MakeErrorReturn(http.StatusNotFound, "user not found", UserNotFound, nil, 1)
		}
	}
	if user.Verified == true {
		return MakeErrorReturn(200, "already verified", 20003, nil, 1)
	}
	user.VerificationCode = internal.GenerateVerificationCode()
	user.UpdateUser()
	return MakeSuccessReturn(database.ToOutputUser(user))
}
