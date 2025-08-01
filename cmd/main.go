package main

import (
	"fmt"
	"go-clean-architecture-playground/infrastructure/persistence"
	"go-clean-architecture-playground/interface/controller"
	"go-clean-architecture-playground/usecase"
	"log"
	"net/http"
	"strings"
)

func main() {
	userRepo := persistence.NewMemoryUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodGet {
			if r.Method == http.MethodPost {
				userController.CreateUser(w, r)
			} else {
				userController.ListUsers(w, r)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/users/") {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			userController.GetUser(w, r)
		case http.MethodPut:
			userController.UpdateUser(w, r)
		case http.MethodDelete:
			userController.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("サーバーを :8080 で起動しています...")
	fmt.Println("API エンドポイント:")
	fmt.Println("  POST   /users      - ユーザー作成")
	fmt.Println("  GET    /users      - ユーザー一覧取得")
	fmt.Println("  GET    /users/{id} - ユーザー取得")
	fmt.Println("  PUT    /users/{id} - ユーザー更新")
	fmt.Println("  DELETE /users/{id} - ユーザー削除")

	log.Fatal(http.ListenAndServe(":8080", nil))
}