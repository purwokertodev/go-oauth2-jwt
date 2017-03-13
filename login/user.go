package login

import(
	"strings"
)

type UserLogin struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func (userLogin *UserLogin) IsValidUser() (bool){
	return strings.ToLower(userLogin.Username) == "wuriyanto" && userLogin.Password == "123456"
}
