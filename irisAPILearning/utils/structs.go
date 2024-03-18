package utils

type UserClaims struct {
	User_Login string
}

type ApiPathsStruct struct {
	Login       string
	Register    string
	Products    string
	Protected   string
	UpdateUser  string
	DeleteUser  string
	GetAllUsers string
}
