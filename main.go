package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context)  {
	c.JSON(200, gin.H{
		"id":"1",
		"nome":"Guilherme Vicente", 
	})
}

func main(){
	r := gin.Default()
	r.GET("/alunos", ExibeAlunos)
	r.Run()
}