package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name  string
	Email string
	ID    int
}

func userRouter(w http.ResponseWriter, r *http.Request) {
	ourUser := User{}
	ourUser.Name = "Keiji Jin"
	ourUser.Email = "keijijin@gmail.com"
	ourUser.ID = 100

	output, _ := goyaml.Marchal(&ourUser)
	fmt.Fprintln(w, string(output))
}

func main() {
	fmt.Println("Starting YAML server")
	http.HandleFunc("/user", userRouter)
	http.ListenAndServe(":8080", nil)
}
