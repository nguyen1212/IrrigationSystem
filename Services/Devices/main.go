package main

import (
	"fmt"
	"log"

	"utils"
)

func main() {
	// Create new Websocket to client
	http.HandleFunc("/ws", handleWSConnections)

	fmt.Println("Server listening on port 8080")
	err	:=	http.ListenAndServe(":8080", nil)
	if err	!= nil{
		log.Fatal(err)
	}
}

