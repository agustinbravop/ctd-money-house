package domain

type User struct {
	// ID es el UUID generado por Keycloak, y también es la Primary Key en la DB.
	ID        string `json:"id"`
	Name      string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dni       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
