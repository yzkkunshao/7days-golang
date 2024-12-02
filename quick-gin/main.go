// geektutu.com
// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string
	Age  int8
}

func main() {
	stu1 := &student{"Tom", 18}
	stu2 := &student{"Picky", 18}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Run(":9998") // listen and serve on 0.0.0.0:8080
}
