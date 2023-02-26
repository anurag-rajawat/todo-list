package types

type Todos struct {
	Id          int `gorm:"primary_key"`
	Description string
	Completed   bool
}
