package domain

type User struct {
	// ID es la Primary Key del User en la DB.
	ID        int
	Name      string
	LastName  string
	Dni       string
	Email     string
	Telephone string
	Cvu       int
	Alias     string
}
