package models

type TempUserTask struct {
	UserID   string //`gorm:"primaryKey;autoIncrement:true"`
	UserName string //`gorm:"primaryKey;"`
}

type TempCompanyTask struct {
	CompanyID    string //`gorm:"primaryKey;autoIncrement:true"`
	CompanyPMSID string
	CompanyName  string //`gorm:"primaryKey;"`
	Address      string
	Phone        string
}

type TempPicTask struct {
	PicID            string
	PicName          string
	GuestIDPMS       string
	Phone            string
	CompanyID        string
	IdentificationID string
}

type TempSales struct {
	SalesID   string
	SalesName string
}

type TempTask struct {
	TaskID string
}

type LoginUser struct {
	UserID    string
	UserName  string
	LevelUser string
}
