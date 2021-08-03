package dto

type UserInfo struct {
	Account  string
	UserName string
	NickName string
}

type UserAccount struct {
	Accout   string
	Password string
}

type UserSession struct {
	Account   string
	Sig       string
	LoginTime uint64
}
