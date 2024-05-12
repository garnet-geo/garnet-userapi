package db

type UserModel struct {
	Id       UserId `json:"id"`
	Domain   DomainId `json:"domain"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserId string

type DomainModel struct {
	Id       DomainId `json:"id"`
	Name     string `json:"name"`
	LongName string `json:"long_name"`
}

type DomainId string

type RepositoryModel struct {
	Id          RepositoryId `json:"id"`
	Domain      DomainId `json:"domain"`
	Name        string `json:"name"`
	LongName    string `json:"long_name"`
	Description string `json:"description"`
}

type RepositoryId string
