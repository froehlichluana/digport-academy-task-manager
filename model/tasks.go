package model

type Tasks struct{
	Description string
	Checklist []Checklist
	Comment string
	StartDate int
	EndDate int
}

type Checklist struct{
	Title string
	Task string
	Done bool
}
