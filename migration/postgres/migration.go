package postgres

import (
	"time"

	"gorm.io/gorm"
)

type TblBlocks struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	BlockTitle  string    `gorm:"type:character varying"`
	Description string    `gorm:"type:text"`
	ImagePath   string    `gorm:"type:character varying"`
	BlockView   string    `gorm:"type:text"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
}

func MigrationTables(db *gorm.DB) {

	db.AutoMigrate(
		&TblBlocks{},
	)

}
