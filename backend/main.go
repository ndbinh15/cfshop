package main

import (
	"log"
	"net/http"

	"cfshop/backend/db"
	"cfshop/backend/handler"
)

func main() {
	db.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../fontend/index.html")
	})

	http.HandleFunc("/login", handler.Login)

	http.Handle("/myapp/", http.StripPrefix("/myapp/", http.FileServer(http.Dir("../fontend"))))

	http.HandleFunc("/home/admin/products/add", handler.InsertProduct)
	http.HandleFunc("/home/admin/products/show", handler.GetAllProducts)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
