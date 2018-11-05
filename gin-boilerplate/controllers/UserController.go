package controllers

import (
	"net/http"

	"gin-boilerplate/models"
	"gin-boilerplate/types"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	return string(hash)
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return false
	}
	return true
}

// Register register a user
func Register(c *gin.Context) {
	var registerRequest types.RegisterRequest
	err := c.BindJSON(&registerRequest)
	if err != nil {
		response := types.APIErrResponse{Msg: "Please check your data", Success: false, Err: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate register request struct
	_, err = govalidator.ValidateStruct(registerRequest)
	if err != nil {
		errMap := govalidator.ErrorsByField(err)
		response := types.APIErrResponse{Msg: "Please check your data", Success: false, Err: errMap}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Maybe add same tag in govalidator
	if registerRequest.Password != registerRequest.PasswordAgain {
		errMap := make(map[string]string)
		errMap["password_again"] = "Password again must be equal to password"
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please check your data", "err": errMap})
		return
	}

	// hash password
	bytePassword := []byte(registerRequest.Password)
	hashedPassword := hashPassword(bytePassword)

	// Save user
	tx, err := models.DB.Begin()
	defer tx.Rollback()

	user := models.User{}
	user.Email = registerRequest.Email
	user.Password = hashedPassword
	user.IsAdmin = 0
	if err = user.Save(tx); err != nil {
		response := types.APIErrResponse{Msg: "Please check your data", Success: false, Err: err.Error()}
		c.JSON(http.StatusNotFound, response)
	} else {
		tx.Commit()
		response := types.APIResponse{Msg: "Register user successfully", Success: true}
		c.JSON(http.StatusOK, response)
	}
}
