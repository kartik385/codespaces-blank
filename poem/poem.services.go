package poem

import "net/http"

func CreatePoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Poem Created!"))
}

func GetPoem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all poems!"))
}
