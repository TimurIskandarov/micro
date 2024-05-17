package model

type RegisterIn struct {
	Name     string
	Email    string
	Password string
}

type RegisterOut struct {
	Message string
	Status  int
}

type AuthorizedIn struct {
	Name     string
	Email    string
	Password string
}

type AuthorizedOut struct {
	Name        string
	AccessToken string
	Message     string
	Status      int
}
