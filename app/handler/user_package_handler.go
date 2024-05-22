package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	"github.com/hendrihmwn/dating-app-api/app/utils"
	"net/http"
)

type UserPackageHandlerImpl struct {
	userPackageService service.UserPackageService
}

type CreateUserPackageRequest struct {
	PackageId int64 `json:"package_id" binding:"required"`
}

func (u *UserPackageHandlerImpl) Create(c *gin.Context) {
	var req CreateUserPackageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]utils.ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = utils.ErrorMessage{Field: fe.Field(), Message: utils.GetErrorMessage(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	res, err := u.userPackageService.Create(c, service.CreateUserPackageParams{
		UserId:    int64(c.GetInt("user_id")),
		PackageId: req.PackageId,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (u *UserPackageHandlerImpl) List(c *gin.Context) {
	res, err := u.userPackageService.List(c, service.ListUserPackageParams{
		UserId: c.GetInt64("user_id"),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

type UserPackageHandler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
}

func NewUserPackageHandler(userPackageService service.UserPackageService) UserPackageHandler {
	return &UserPackageHandlerImpl{userPackageService: userPackageService}
}
