package users

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	Role         string
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewUser(name, email, passwordHash, role string) User {
	if role == "" {
		role = "technician"
	}

	return User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}
}
