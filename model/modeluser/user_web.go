package modeluser

import "time"

type Request struct {
	ID       uint   `json:"id,omitempty" swaggerignore:"true"`
	Fullname string `json:"full_name" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
}

type Response struct {
	ID        uint       `json:"id"  example:"1"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	CreatedAt *time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	Fullname  string     `json:"full_name" example:"jhondoe"`
	Email     string     `json:"email" example:"test@example.com"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJxd2Vxd2..."`
}

// ExampleRequestUpdate only for example swaggo docs
type ExampleRequestUpdate struct {
	Fullname string `json:"full_name" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"your account has been successfully deleted"`
}
