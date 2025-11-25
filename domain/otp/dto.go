package otp

type InputRequestOTP struct {
	Email string `json:"email" validate:"required,email,endswith=@gmail.com"`
}
