package otp

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Validate *validator.Validate
}

func NewHandler(v *validator.Validate) *Handler {
	return &Handler{Validate: v}
}

func (h *Handler) RequestOTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var input InputRequestOTP

		// validasi format (json)
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "JSON formatnya acak-acakan bro",
			})

			return
		}

		// validasi lanjutan, bagian body (email)
		if err := h.Validate.Struct(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			// note: nanti nampilim pretelin errornya biar pesannya spesifik, sekarang bad request aja dulu awokawoiadaddeheheheheh
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Format salah! Harus email valid dan pake @gmail.com",
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Mantap, OTP dikirim ke " + input.Email,
			// tapi boong hehehe, belum bikin fiturnya
		})

	}
}
