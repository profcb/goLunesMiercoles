package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Response string `json:"response"`
}

func main() {
	router := gin.Default()

	// POST /ping
	router.POST("/ping", func(c *gin.Context) {
		var req PingRequest

		// Bind JSON
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON inválido",
			})
			return
		}

		resp := PingResponse{
			Response: "pong",
		}

		// Lógica dinámica igual que tu ejemplo
		if req.Message != "ping" {
			resp.Response = "mensaje desconocido"
		}

		c.JSON(http.StatusOK, resp)
	})

	router.Run(":8080")
}

/*


GET -> PIDE
RESPONSE<-

POST -> ENVIA
RESPONSE <-

PUT-> MODIFICA
RESPONSE<-

DELETE-> BORRA
RESPONSE<-


*/