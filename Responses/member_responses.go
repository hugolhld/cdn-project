package routes

import (
	controllers "github.com/hugolhld/cdn-project/Controllers"

	"github.com/gorilla/mux"
)

func MemberRoutes(router *mux.Router) {

	router.HandleFunc("/member", controllers.CreateMember()).Methods("POST")
	router.HandleFunc("/member/{id}", controllers.GetMember()).Methods("GET")
	router.HandleFunc("/members", controllers.GetAllMembers()).Methods("GET")
	router.HandleFunc("/member/{id}", controllers.UpdateMember()).Methods("PUT")
	router.HandleFunc("/member/{id}", controllers.DeleteMember()).Methods("DELETE")

}
