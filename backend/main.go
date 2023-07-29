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
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("../img"))))

	http.HandleFunc("/home/admin/products/count", handler.GetProductCount)
	http.HandleFunc("/home/admin/products/add", handler.InsertProduct)
	http.HandleFunc("/home/admin/products/show", handler.GetAllProducts)

	http.HandleFunc("/products/categories/get", handler.GetAllCategory)

	log.Println("Server is running at http://localhost:1235")
	log.Fatal(http.ListenAndServe(":1235", nil))
}
