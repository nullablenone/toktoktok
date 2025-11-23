package main

import (
	"net/http"

	"github.com/nullablenone/toktoktok/domain/otp"
	"github.com/nullablenone/toktoktok/routes"
)

func main() {
	OtpHandler := otp.NewHandler()

	router := routes.SetupRoutes(OtpHandler)

	http.ListenAndServe(":808", router)
}
