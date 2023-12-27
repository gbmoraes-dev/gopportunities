package schemas

import (
	"gorm.io/gorm"
)

type Opportunity struct {
	gorm.Model
	Role     string
	Company  string
	Logo     string
	IsRemote bool
	Link     string
	Salary   string
}
