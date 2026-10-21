package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/almarufh/backend-shortlink/internal/services"
	"github.com/almarufh/backend-shortlink/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type AuthContorller struct {
	as *services.AuthService
}

func NewAuthContorller(as *services.AuthService) *AuthContorller {
	return &AuthContorller{
		as: as,
	}
}

// Register
//
// @Summary			Create New Account
// @Description		Create new account for full access
// @Tags			Auth
// @Accept			json
// @Produce			json
// @Router			/auth/register	[post]
// @Param			body	body 		dto.RegisterRequest true "Body data new user"
// @Success			200		{object}	dto.ResponseSuccess
// @Failure			400		{object}	dto.ResponseError
// @Failure			409		{object}	dto.ResponseError
// @Failure			500		{object}	dto.ResponseError
func (ac *AuthContorller) RegisterAC(c *gin.Context) {
	var body dto.RegisterRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Printf("[AuthHandler.Register] ShouldBindWith: %v\n", err)
		found, isErr := errors.AsType[validator.ValidationErrors](err)
		if isErr {
			err := utils.ValidationError(found)
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Success: false,
				Message: err,
			})
			return
		}

		c.JSON(http.StatusBadRequest, dto.ResponseError{
			Success: false,
			Message: "invalid request payload",
		})
		return
	}

	if err := ac.as.RegisterAS(c.Request.Context(), body); err != nil {
		if errors.Is(err, message.ErrNoRowAffected) {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Success: false,
				Message: "Invalid input data",
			})
			return
		}

		if errors.Is(err, message.ErrEmailExist) {
			c.JSON(http.StatusConflict, dto.ResponseError{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseSuccess{
		Success: true,
		Message: "register successfully",
	})
}

// Login
//
// @Summary			Log In
// @Description		Log in for get full access
// @Tags			Auth
// @Accept			json
// @Produce			json
// @Router			/auth/login	[post]
// @Param			body	body 		dto.LoginRequest true "Body data to access ewalet"
// @Success			200		{object}	dto.ResponseSuccess
// @Failure			400		{object}	dto.ResponseError
func (ac *AuthContorller) LoginAC(c *gin.Context) {
	var body dto.LoginRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Printf("[AuthContorller.LoginAC] Bind error: %v\n", err.Error())
		found, isErr := errors.AsType[validator.ValidationErrors](err)
		if isErr {
			err := utils.ValidationError(found)
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Success: false,
				Message: err,
			})
			return
		}
	}

	res, err := ac.as.LoginAS(c.Request.Context(), body)
	if err != nil {
		log.Printf("[AuthContorller.LoginAS] Bind error: %v\n", err.Error())
		if errors.Is(err, message.ErrEmailExist) {
			c.JSON(http.StatusBadRequest, dto.ResponseError{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccess{
		Success: true,
		Message: "login successfully",
		Data:    res,
	})
}
