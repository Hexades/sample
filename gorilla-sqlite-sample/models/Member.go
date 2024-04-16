package models
type Member struct {
	//gorm.Model
	MemberId string `gorm:"primaryKey" json:"member_id"`
	First    string `gorm:"<-" json:"first"`
	Last     string `gorm:"<-" json:"last"`
}
