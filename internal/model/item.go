package model

type Item struct {
	CommonFields
	Name string
}

type NewItem struct {
	Name string
}
