package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	"github.com/hendrihmwn/dating-app-api/app/utils"
	"net/http"
)

type UserPairHandlerImpl struct {
	userPairService service.UserPairService
}

type CreateUserPairRequest struct {
	PairUserId int64 `json:"pair_user_id" binding:"required"`
	Status     int   `json:"status" binding:"required,max=1"`
}

func (u *UserPairHandlerImpl) Create(c *gin.Context) {
	var req CreateUserPairRequest

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

	res, err := u.userPairService.Create(c, service.CreateUserPairParams{
		UserId:     int64(c.GetInt("user_id")),
		PairUserId: req.PairUserId,
		Status:     req.Status,
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

type UserPairHandler interface {
	Create(c *gin.Context)
}

func NewUserPairHandler(userPairService service.UserPairService) UserPairHandler {
	return &UserPairHandlerImpl{userPairService: userPairService}
}
