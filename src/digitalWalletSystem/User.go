package main

type User struct {
	UserId   string
	UserName string
	EmailId  string
	Password string
}

func NewUser(userId string, userName string, emailId string, password string) User {
	return User{UserId: userId, UserName: userName, EmailId: emailId, Password: password}

}
