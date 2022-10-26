package domain

type User struct {
	// ID es la Primary Key del User en la DB.
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Dni       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Cvu       string `json:"cvu"`
	Alias     string `json:"alias"`
}
