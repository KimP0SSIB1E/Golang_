package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data  string `json:"User Logged in Successfully"`
	Error string `json:"error,omitempty"`
}
