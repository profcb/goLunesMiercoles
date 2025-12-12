package main


import (
	"encoding/json"
	"log"
	"net/http"
)

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Response string `json:"response"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Aceptamos solo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req PingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Lógica “ping → pong”
	resp := PingResponse{
		Response: "pong",
	}

	// Opcional: comportamiento dinámico
	if req.Message != "ping" {
		resp.Response = "mensaje desconocido"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}