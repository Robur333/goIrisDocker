package utils

import "os"

var ApiPaths = ApiPathsStruct{
	"/login",
	"/register",
	"/products",
	"/protected",
	"/updateUser",
	"/deleteUser",
	"/getAllUsers"}

var (
	SignKey = []byte(os.Getenv("SIGN_KEY"))
	EncKey  = []byte(os.Getenv("ENC_KEY"))
)