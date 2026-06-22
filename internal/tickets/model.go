package tickets

type Ticket struct {
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

func NewUser(name, email, passwordHash, role string) Ticket {
	if role == "" {
		role = "technician"
	}

	return Ticket{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}
}
