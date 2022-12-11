package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"example.com/m/poem"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middle)
	r.Route("/api", func(r chi.Router) {
		r.Mount("/poem", poem.PoemRouter())
	})
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()
}

func middle(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "Name", "Kartik")
		r = r.WithContext(ctx)
		nxt.ServeHTTP(w, r)

	})
}

func coc(w http.ResponseWriter, r *http.Request) {
	b := r.Context().Value("Name")
	a := fmt.Sprintf("Heelo World %d", b)
	w.Write([]byte(a))
}
