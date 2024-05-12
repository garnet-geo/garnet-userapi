package handlers

type UserInfoModel struct {
	Id       UserId     `json:"id"`
	Name     DomainName `json:"name"`
	LongName string     `json:"long_name"`
	Email    UserEmail  `json:"email"`
}

type UserId string

type DomainId string

type DomainName string

type UserEmail string
