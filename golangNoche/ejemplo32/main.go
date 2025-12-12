package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Alumno struct{
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Cursos int `json:"cursos"`
}

var alumnos []Alumno
var nextID = 1

func main(){
	fmt.Println("Arranca nuestra aplicacion")
	r := gin.Default()
	//r.SetTrustedProxies(nil)

	a := Alumno{
		ID: nextID,
		Nombre: "Base",
		Cursos: 2,
	}
	
	nextID++

	alumnos = append(alumnos, a)

	r.GET("/alumnos",listarAlumnos)
	r.GET("/alumnos/:id",obtenerAlumno)
    r.POST("/alumnos", crearAlumno)
    r.PUT("/alumnos/:id", actualizarAlumno)
    r.DELETE("/alumnos/:id", eliminarAlumno)

    r.Run(":8080")
}

func listarAlumnos(c *gin.Context) {
    c.JSON(http.StatusOK, alumnos)
}

func obtenerAlumno(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, a := range alumnos {
        if a.ID == id {
            c.JSON(http.StatusOK, a)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
}

func crearAlumno(c *gin.Context) {
    var a Alumno
    if err := c.BindJSON(&a); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }

    a.ID = nextID
    nextID++

    alumnos = append(alumnos, a)
    c.JSON(http.StatusOK, a)
}

func actualizarAlumno(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var input Alumno

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }

    for i, a := range alumnos {
        if a.ID == id {
            alumnos[i].Nombre = input.Nombre
            alumnos[i].Cursos = input.Cursos
            c.JSON(http.StatusOK, alumnos[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
}

func eliminarAlumno(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, a := range alumnos {
        if a.ID == id {
            alumnos = append(alumnos[:i], alumnos[i+1:]...)
            c.Status(http.StatusNoContent)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
}