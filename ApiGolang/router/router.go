package router

import "github.com/gin-gonic/gin"


func Initialize(){
// Inicializa o Router utilizando as configuracoes default do gin
router  := gin.Default()
//Definido uma rota
router.GET("/ping", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
})
router.Run() // listen and serve on 0.0.0.0:8080
}