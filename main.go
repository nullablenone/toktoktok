package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nullablenone/toktoktok/domain/otp"
	"github.com/nullablenone/toktoktok/routes"
)

func main() {
	validate := validator.New()

	OtpHandler := otp.NewHandler(validate)

	router := routes.SetupRoutes(OtpHandler)

	fmt.Println("run in port 808")
	if err := http.ListenAndServe(":808", router); err != nil {
		log.Println("error: ", err)
	}
}
