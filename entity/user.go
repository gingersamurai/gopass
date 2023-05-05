package entity

type Account struct {
	Id        int64
	UserId    int64
	ServiceId int64
	Login     string
	Password  string
}

type User struct {
	Id       int64
	Login    string
	Password string
}

type Service struct {
	Id   int64
	Name string
}
