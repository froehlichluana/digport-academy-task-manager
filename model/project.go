package model

type Project struct{
	Name string
	ID string
	Member []Member
}

type Member struct{
	Name string
	ID string
	Role string
}

