package handlers

type UserInfoModel struct {
	Id       UserId     `json:"id,omitempty"`
	Name     DomainName `json:"name,omitempty"`
	LongName string     `json:"long_name,omitempty"`
	Email    UserEmail  `json:"email,omitempty"`
	Password string     `json:"password,omitempty"`
}

type UserId string

type DomainId string

type DomainName string

type UserEmail string
