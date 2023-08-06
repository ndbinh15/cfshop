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
		http.ServeFile(w, r, "../frontend/index.html")
	})

	http.HandleFunc("/login/login", handler.Login)

	http.Handle("/myapp/", http.StripPrefix("/myapp/", http.FileServer(http.Dir("../frontend"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("../img"))))

	http.HandleFunc("/products/count", handler.GetProductCount)
	http.HandleFunc("/products/add", handler.InsertProduct)
	http.HandleFunc("/products/addQuantity", handler.UpdateProductQuantity)
	http.HandleFunc("/products/get", handler.GetAllProducts)
	http.HandleFunc("/products/delete", handler.DeleteProduct)

	http.HandleFunc("/user/register", handler.CreateUser)

	http.HandleFunc("/products/categories/get", handler.GetAllCategory)

	log.Println("Server is running at http://localhost:1235")
	log.Fatal(http.ListenAndServe(":1235", nil))
}
