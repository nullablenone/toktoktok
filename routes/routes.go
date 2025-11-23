package routes

import (
	"github.com/gorilla/mux"
	"github.com/nullablenone/toktoktok/domain/otp"
)

func SetupRoutes(otpHandler *otp.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/request-otp", otpHandler.RequestOTP()).Methods("POST")

	return router
}
