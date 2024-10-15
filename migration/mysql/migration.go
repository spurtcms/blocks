package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblBlock struct {
	Id           int       `gorm:"primaryKey;auto_increment;"`
	Title        string    `gorm:"type:varchar(255)"`
	BlockContent string    `gorm:"type:text"`
	CoverImage   string    `gorm:"type:varchar(255)"`
	TenantId     int       `gorm:"type:int"`
	Prime        int       `gorm:"type:int"`
	IsActive     int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy    int       `gorm:"type:int"`
	ModifiedBy   int       `gorm:"type:integer"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer;DEFAULT:0"`
}

type TblBlockTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;"`
	BlockId   int       `gorm:"type:int"`
	TagId     int       `gorm:"type:int"`
	TagName   string    `gorm:"type:varchar(255)"`
	TenantId  int       `gorm:"type:int"`
	CreatedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:int"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:int;DEFAULT:0"`
}

type TblBlockMstrTag struct {
	Id        int       `gorm:"primaryKey;auto_increment;"`
	TagTitle  string    `gorm:"type:varchar(255)"`
	TenantId  int       `gorm:"type:int"`
	CreatedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:int"`
}

type TblBlockCollection struct {
	Id        int       `gorm:"primaryKey;auto_increment;"`
	UserId    int       `gorm:"type:int"`
	BlockId   int       `gorm:"type:int"`
	TenantId  int       `gorm:"type:int"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:int;DEFAULT:0"`
}

func MigrationTables(db *gorm.DB) {

	db.AutoMigrate(
		&TblBlock{},
		&TblBlockCollection{},
		&TblBlockMstrTag{},
		&TblBlockTags{},
	)

}
