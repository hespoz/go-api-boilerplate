package routes

import (
	"frenco-api/controllers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{postID}", GetPost)
	router.Post("/", NewPost)
	router.Get("/", GetPostList)
	return router
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	post := Post{
		ID:      postID,
		Content: "Content post",
	}
	render.JSON(w, r, post)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	post := Post{
		ID:      "AXDFTZ",
		Content: "Content post",
	}
	render.JSON(w, r, post)
}

func GetPostList(w http.ResponseWriter, r *http.Request) {
	postList := []Post{
		{
			ID:      "AXDFTZ",
			Content: controllers.Test(),
		},
		{
			ID:      "AXDFTZ",
			Content: "Content post 2",
		},
		{
			ID:      "AXDFTZ",
			Content: "Content post 3",
		},
	}
	render.JSON(w, r, postList)
}
