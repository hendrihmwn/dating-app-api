package handler_test

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/dating-app-api/app/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandlerImpl_ListUserServer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	router := gin.Default()

	userHandler := handler.NewUserServer()

	// Define your handler
	router.GET("/books", userHandler.ListUserServer)

	t.Run("Success", func(t *testing.T) {

		// a response recorder for getting written http response
		request, err := http.NewRequest(http.MethodGet, "/books", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusOK, rr.Code)

		// use a middleware to set context for test
		// the only claims we care about in this test
		// is the UID
		//router := gin.Default()
		//
		//request, err := http.NewRequest(http.MethodGet, "/books", nil)
		//assert.NoError(t, err)
		//
		//router.ServeHTTP(rr, request)
		//
		//assert.Equal(t, 200, rr.Code)
		//mockUserHandler.AssertExpectations(t) // assert that UserService.Get was called
	})
}
