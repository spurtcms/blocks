package postgres

import (
	"time"

	"gorm.io/gorm"
)

type TblBlock struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	BlockDescription string    `gorm:"type:text"`
	BlockContent     string    `gorm:"type:text"`
	BlockCss         string    `gorm:"type:text"`
	CoverImage       string    `gorm:"type:character varying"`
	IconImage        string    `gorm:"type:character varying"`
	Free             int       `gorm:"type:integer"`
	Prime            int       `gorm:"type:integer"`
	TenantId         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
}

type TblBlockTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	BlockId   int       `gorm:"type:integer"`
	TagId     int       `gorm:"type:integer"`
	TagName   string    `gorm:"type:character varying"`
	TenantId  int       `gorm:"type:integer"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
}

type TblBlockMstrTag struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name      string    `gorm:"type:character varying"`
	TenantId  int       `gorm:"type:integer"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
}

type TblBlockCollection struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId    int       `gorm:"type:integer"`
	BlockId   int       `gorm:"type:integer"`
	TenantId  int       `gorm:"type:integer"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}

func MigrationTables(db *gorm.DB) {

	db.AutoMigrate(
		&TblBlock{},
		&TblBlockCollection{},
		&TblBlockMstrTag{},
		&TblBlockTags{},
	)

}
