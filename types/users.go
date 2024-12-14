package types

type User struct {
	PrivateKey string
	Publickey  string
	Hash       string
}

type UserFetch struct {
	Publickey string
}
