package poem

import (
	"github.com/go-chi/chi/v5"
)

func PoemRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Get("/create", CreatePoem)
	r.Get("/getAll", GetPoem)

	return r
}
