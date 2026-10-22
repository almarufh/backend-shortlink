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

type LinksControl struct {
	ls *services.LinksService
}

func NewLinksControl(ls *services.LinksService) *LinksControl {
	return &LinksControl{
		ls: ls,
	}
}

// Create ShortURL
//
// @Summary         Create New Short Link
// @Description     Create new short link for authenticated user
// @Tags            Links
// @Accept          json
// @Produce         json
// @Security		Bearer
// @Router          /api/links  [post]
// @Param           body    body        dto.LinksRequest true "Body data create short link"
// @Success         201     {object}    dto.ResponseSuccess
// @Failure         400     {object}    dto.ResponseError
// @Failure         409     {object}    dto.ResponseError
// @Failure         500     {object}    dto.ResponseError
func (lc *LinksControl) CreateShortURLcontrol(c *gin.Context) {
	var body dto.LinksRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Printf("[CreateShortURLcontrol.ShouldBindWith] Bind error: %v\n", err.Error())
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

	res, exist := c.Get("ID")
	if !exist {
		log.Printf("[CreateShortURLcontrol] Error get token from redis")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	userID, ok := res.(int32)
	if !ok {
		log.Printf("[CreateShortURLcontrol] Error change to string")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}
	body.UserID = userID

	if err := lc.ls.CreateShortURLservice(c.Request.Context(), body); err != nil {
		log.Printf("[CreateShortURLcontrol] Error change to string")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccess{
		Success: true,
		Message: "create shortlink successfully",
	})
}

// Redirect
//
// @Summary         Redirect Short Link
// @Description     Redirect user to the original URL by short_url slug
// @Tags            Redirect
// @Param           slug    path        string    true    "Short URL slug"
// @Router          /{slug} [get]
// @Success         302          "Redirect to original URL"
// @Failure         404          {object}    dto.ResponseError
// @Failure         500          {object}    dto.ResponseError
func (lc *LinksControl) RedirectControl(c *gin.Context) {
	shortURL := c.Param("slug")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, dto.ResponseError{
			Success: false,
			Message: "Slug short URL tidak boleh kosong",
		})
		return
	}

	originalURL, err := lc.ls.GetOriginalURLService(c.Request.Context(), shortURL)
	if err != nil {
		log.Printf("[RedirectControl.GetOriginalURLService] Error: %v\n", err)

		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, *originalURL)
}

// Get MyLinks
//
// @Summary         Get My Short Link
// @Description     Get My short link for authenticated user
// @Tags            Links
// @Produce         json
// @Security		Bearer
// @Router          /api/links  [get]
// @Success         200     {object}    dto.ResponseSuccess
// @Failure         400     {object}    dto.ResponseError
// @Failure         500     {object}    dto.ResponseError
func (lc *LinksControl) GetMyLinksControl(c *gin.Context) {
	res, exist := c.Get("ID")
	if !exist {
		log.Printf("[GetMyLinksControl] Error get token from redis")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	userID, ok := res.(int32)
	if !ok {
		log.Printf("[LogoutAC.res] Error change to string")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	res, err := lc.ls.GetMyLinksService(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccess{
		Success: true,
		Message: "get my links successfully",
		Data:    res,
	})
}

// Delete Link
//
// @Summary         Delete Link
// @Description     Delete Link for authenticated user
// @Tags            Links
// @Produce         json
// @Security		Bearer
// @Param           slug    path        string    true    "Short URL slug"
// @Router          /api/links/{slug}   [delete]
// @Success         200     {object}    dto.ResponseSuccess
// @Failure         400     {object}    dto.ResponseError
// @Failure         500     {object}    dto.ResponseError
func (lc *LinksControl) DelLinkRControl(c *gin.Context) {
	shortURL := c.Param("slug")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, dto.ResponseError{
			Success: false,
			Message: "Slug short URL tidak boleh kosong",
		})
		return
	}

	res, exist := c.Get("ID")
	if !exist {
		log.Printf("[GetMyLinksControl] Error get token from redis")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	userID, ok := res.(int32)
	if !ok {
		log.Printf("[LogoutAC.res] Error change to string")
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: message.ErrInternal.Error(),
		})
		return
	}

	err := lc.ls.DelLinkService(c.Request.Context(), userID, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccess{
		Success: true,
		Message: "delete links successfully",
	})
}
