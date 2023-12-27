package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opportunity struct {
	gorm.Model
	Role     string
	Company  string
	IsRemote bool
	Link     string
	Salary   string
}

type OpportunityResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"uptadedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	IsRemote  bool      `json:"isRemote"`
	Link      string    `json:"link"`
	Salary    string    `json:"salary"`
}
