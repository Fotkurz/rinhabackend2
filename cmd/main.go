package main

import (
	"net/http"

	"github.com/Fotkurz/rinhabackend2/internal/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	db.Connect()
	defer db.CONN.Close()

	http.ListenAndServe(":3000", r)
}
