package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/controllers"
	"go-api/entities"
	"go-api/types"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gopkg.in/go-playground/validator.v9"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Group(func(router chi.Router) {
		router.Get("/{userID}", GetUser)
		router.With(Validation(&entities.User{})).Post("/", NewUser)
		router.Get("/", GetUserList)
	})

	return router
}

var validate *validator.Validate

func Validation(objectType interface{}) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//var newObj1 = reflect.TypeOf(objectType).Elem()

			fmt.Println(reflect.Zero(reflect.TypeOf(objectType)))

			var payload *entities.User

			json.NewDecoder(r.Body).Decode(&payload)

			fmt.Println("Llega aca")
			fmt.Println(payload)

			validate = validator.New()

			err := validate.Struct(payload)
			if err != nil {
				fmt.Println("ererrw")
				fmt.Println(err)
				render.JSON(w, r, "Error validacion")
			}

			ctx := context.WithValue(r.Context(), "user", payload)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user := types.User{
		ID:       userID,
		Email:    "Content post",
		Password: "Password",
	}
	render.JSON(w, r, user)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	newUser, ok := ctx.Value("user").(*entities.User)
	if !ok {
		render.JSON(w, r, "err")
	}

	err := controllers.CreateUser(newUser)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, newUser)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	userList, err := controllers.FindAll()

	if err != nil {
		fmt.Println(err)
		render.JSON(w, r, nil)
	}

	render.JSON(w, r, userList)
}
