package Routes

import (
	controllers "cdn-project/Controllers"

	"github.com/gorilla/mux"
)

func FileRoutes(router *mux.Router) {
	router.HandleFunc("/upload", controllers.Upload()).Methods("POST")
	router.HandleFunc("/files", controllers.GetFiles()).Methods("GET")
}
