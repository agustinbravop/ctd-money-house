package domain

type User struct {
	// ID es la Primary Key del User en la DB.
	ID int
	// KeycloakID es autogenerado por Keycloak. Es único para cada User. No debe ir en las requests.
	// Sirve para que cuando el usuario haga login y keycloak devuelva su KeycloakID, se pueda buscar sus datos en la DB.
	KeycloakID string
	Name       string
	LastName   string
	Dni        string
	Email      string
	Telephone  string
	Cvu        int
	Alias      string
}
