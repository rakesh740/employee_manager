package data

type Employee struct {
	ID       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(255)"`
	Position string
	Salary   float64 // can be taken as uint64 to remove floating point error
}
