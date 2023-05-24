package structs

type User struct {
	Token      string `json:"token"`
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	FacebookID string `json:"FacebookID"`
	GoogleID   string `json:"GoogleID"`
	Avatar     string `json:"avatar"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type RegisterSchema struct {
	FirstName string `json:"firstName" validate:"required,alpha,min=5,max=128"`
	LastName  string `json:"lastName" validate:"required,alpha,min=5,max=128"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,alphanum,min=6,max=128"`
}

type LoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=6,max=128"`
}
