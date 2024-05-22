package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hendrihmwn/dating-app-api/app/domain/service"
	"github.com/hendrihmwn/dating-app-api/app/utils"
	"net/http"
)

type UserHandlerImpl struct {
	userService service.UserService
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,max=50"`
	Email    string `json:"email" binding:"required,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type ListRequest struct {
	Limit *int `form:"limit"`
}

func (u *UserHandlerImpl) Login(c *gin.Context) {
	var req LoginRequest

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

	res, err := u.userService.Login(c, service.UserLoginParams{
		Email:    req.Email,
		Password: req.Password,
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

func (u *UserHandlerImpl) Register(c *gin.Context) {
	var req RegisterRequest

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

	res, err := u.userService.Register(c, service.UserRegisterParams{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
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

func (u *UserHandlerImpl) List(c *gin.Context) {
	var req ListRequest

	if err := c.ShouldBindQuery(&req); err != nil {
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
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	userListParam := service.UserListParams{
		UserId: int64(c.GetInt("user_id")),
		Limit:  limit,
	}
	res, err := u.userService.List(c, userListParam)
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

type UserHandler interface {
	List(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{userService: userService}
}
