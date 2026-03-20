package user

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleArtist Role = "artist"
)

type User struct {
	ID           string `json:"id"`
	UserName     string `json:"username"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role"`
}
