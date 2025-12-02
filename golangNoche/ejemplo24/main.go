package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Conectar o crear la base de datos (archivo local)
	db, err := sql.Open("sqlite3", "./usuarios.db")
	if err != nil {
		log.Fatal("Error", err)
		return
	}
	defer db.Close()

	// Crear tabla si no existe
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		edad INTEGER
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabla creada o verificada")

	// --- CREATE ---
	insertarUsuario(db, "Carlos", 32)
	insertarUsuario(db, "Lucía", 28)
	insertarUsuario(db, "Juan", 45)

	// --- READ ---
	fmt.Println("Listado de usuarios:")
	listarUsuarios(db)

	// --- UPDATE ---
	fmt.Println("Actualizando edad de Carlos a 33...")
	actualizarEdad(db, 1, 33)
	listarUsuarios(db)

	// --- DELETE ---
	fmt.Println("Borrando usuario con ID 2...")
	eliminarUsuario(db, 2)
	listarUsuarios(db)
}

// --- FUNCIONES CRUD ---

func insertarUsuario(db *sql.DB, nombre string, edad int) {
	stmt, err := db.Prepare("INSERT INTO usuarios(nombre, edad) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombre, edad)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuario agregado: %s (%d años)\n", nombre, edad)
}

func listarUsuarios(db *sql.DB) {
	rows, err := db.Query("SELECT id, nombre, edad FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, edad int
		var nombre string
		err = rows.Scan(&id, &nombre, &edad)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Nombre: %s | Edad: %d\n", id, nombre, edad)
	}
}

func actualizarEdad(db *sql.DB, id int, nuevaEdad int) {
	stmt, err := db.Prepare("UPDATE usuarios SET edad = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(nuevaEdad, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Edad actualizada correctamente")
}

func eliminarUsuario(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Usuario eliminado correctamente")
}
