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

	http.Handle("/myapp/", http.StripPrefix("/myapp/", http.FileServer(http.Dir("../frontend"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("../img"))))

	http.HandleFunc("/login/login", handler.Login)
	http.HandleFunc("/user/register", handler.CreateUser)

	http.HandleFunc("/products/categories/get", handler.GetAllCategory)
	http.HandleFunc("/products/count", handler.GetProductCount)
	http.HandleFunc("/products/add", handler.InsertProduct)
	http.HandleFunc("/products/addQuantity", handler.UpdateProductQuantity)
	http.HandleFunc("/products/delete", handler.DeleteProduct)
	http.HandleFunc("/products/update", handler.UpdateProductByID)
	http.HandleFunc("/products/get", handler.GetAllProducts)
	http.HandleFunc("/products", handler.GetProductById)

	http.HandleFunc("/users/count", handler.GetUserCount)
	http.HandleFunc("/users/update", handler.UpdateUserByID)
	http.HandleFunc("/users/get", handler.GetAllUsers)
	http.HandleFunc("/users", handler.GetUserById)
	// http.HandleFunc("/users/delete", handler.DeleteProduct)

	http.HandleFunc("/cart/add", handler.AddToCart)
	http.HandleFunc("/cart/get", handler.GetCartByUserID)

	http.HandleFunc("/orders/create", handler.CreateOrder)
	http.HandleFunc("/orders/count", handler.GetOrderCount)
	http.HandleFunc("/orders/countInprogress", handler.GetInProgressOrderCount)
	http.HandleFunc("/orders/countCompleted", handler.GetCompletedOrderCount)
	http.HandleFunc("/orders/getInprogress", handler.GetAllInProgressOrders)
	http.HandleFunc("/orders/getCompleted", handler.GetAllCompletedOrders)
	http.HandleFunc("/orders/get", handler.GetOrdersByUserID)
	http.HandleFunc("/orders/markAsDone", handler.MarkOrderAsDoneByID)

	log.Println("Server is running at http://localhost:1235")
	log.Fatal(http.ListenAndServe(":1235", nil))
}
