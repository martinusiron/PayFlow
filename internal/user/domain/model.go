package domain

type User struct {
	ID        int
	Username  string
	Password  string
	Salary    float64
	Role      string // "admin" or "employee"
	CreatedAt string
	UpdatedAt string
	CreatedBy *int
	UpdatedBy *int
	IPAddress string
	RequestID string
}
