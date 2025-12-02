package main

import (
	"fmt"
	"encoding/json"
	"os"
)



type Config struct{
	URL string `json:"url"`
	Usuario string `json:"usuario"`
	Puerto int `json:"puerto"`
}

func main(){
	data,err:= os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	var c Config
	err = json.Unmarshal(data,&c)
	if err != nil{
		fmt.Println("Error",err)
		return
	}
	fmt.Println("URL",c.URL)
	fmt.Println("Usuario",c.Usuario)
	fmt.Println("Puerto",c.Puerto)
}


/*
func main() {
    jsonStr := `{
        "url":"prod-ejemplo.com.ar",
        "usuario":"admin",
        "puerto":8000,
        "extra": {
            "modo":"produccion",
            "debug":false
        }
    }`

    var data map[string]interface{}

    json.Unmarshal([]byte(jsonStr), &data)

    fmt.Println(data)
}
	*/