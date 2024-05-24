package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermevicente-hub/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupRotas() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetupRotas()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
	mockRes := `{"API diz:":"E ai gui, tudo beleza ? "}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody))
	fmt.Println(string(resBody))
	fmt.Println(string(mockRes))
}
