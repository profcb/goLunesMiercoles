package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pedido struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	ComboUno int    `json:"combo_uno"`
	ComboDos int    `json:"combo_dos"`
	Postre   int    `json:"postre"`
}

var pedidos []Pedido
var nextID = 1

// Pedido inicial
func seedData() {
	p := Pedido{
		ID:       nextID,
		Nombre:   "Pablo",
		ComboUno: 2,
		ComboDos: 1,
		Postre:   1,
	}
	nextID++
	pedidos = append(pedidos, p)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	seedData() // carga un pedido inicial

	r.GET("/pedidos", listarPedidos)
	r.GET("/pedidos/:id", obtenerPedido)
	r.POST("/pedidos", crearPedido)
	r.PUT("/pedidos/:id", actualizarPedido)
	r.DELETE("/pedidos/:id", eliminarPedido)

	r.Run(":8080")
}

func listarPedidos(c *gin.Context) {
	c.JSON(http.StatusOK, pedidos)
}

func obtenerPedido(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, p := range pedidos {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
}

func crearPedido(c *gin.Context) {
	var p Pedido
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	p.ID = nextID
	nextID++

	pedidos = append(pedidos, p)
	c.JSON(http.StatusOK, p)
}

func actualizarPedido(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input Pedido
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for i, p := range pedidos {
		if p.ID == id {
			pedidos[i].Nombre = input.Nombre
			pedidos[i].ComboUno = input.ComboUno
			pedidos[i].ComboDos = input.ComboDos
			pedidos[i].Postre = input.Postre

			c.JSON(http.StatusOK, pedidos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
}

func eliminarPedido(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, p := range pedidos {
		if p.ID == id {
			pedidos = append(pedidos[:i], pedidos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
}
