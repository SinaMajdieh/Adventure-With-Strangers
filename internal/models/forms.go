package models

// Forms used to transfer data from front-end to api/back-end

// Register Form model
type RegisterForm struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Confirmation string `json:"confirmation"`
}

// Login Form model
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
