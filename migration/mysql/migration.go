package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblBlocks struct {
	Id          int       `gorm:"primaryKey;auto_increment;"`
	BlogTitle   string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int;DEFAULT:0"`
	ImagePath   string    `gorm:"type:varchar(255)"`
	BlocksView  string    `gorm:"type:varchar(600)"`
}

func MigrationTables(db *gorm.DB) {

	db.AutoMigrate(
		&TblBlocks{},
	)

}
