package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermevicente-hub/api-go-gin/controllers"
	"github.com/guilhermevicente-hub/api-go-gin/database"
	"github.com/guilhermevicente-hub/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
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

func TestListHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotas()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	fmt.Println(res.Body)
}

func TestBuscaCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotas()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
