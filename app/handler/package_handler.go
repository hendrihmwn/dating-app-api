package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	"github.com/hendrihmwn/dating-app-api/app/utils"
	"net/http"
)

type PackageHandlerImpl struct {
	packageService service.PackageService
}

type GetPackageRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type PackageHandler interface {
	List(c *gin.Context)
	Get(c *gin.Context)
}

func (u *PackageHandlerImpl) List(c *gin.Context) {
	res, err := u.packageService.List(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res.Data,
	})
}

func (u *PackageHandlerImpl) Get(c *gin.Context) {
	var req GetPackageRequest

	if err := c.ShouldBindUri(&req); err != nil {
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

	res, err := u.packageService.Get(c, service.GetPackageParams{
		ID: req.ID,
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

func NewPackageHandler(packageService service.PackageService) PackageHandler {
	return &PackageHandlerImpl{
		packageService: packageService,
	}
}
