package blocks

import (
	"fmt"
	"time"

	"github.com/spurtcms/blocks/migration"
)

// role&permission setup config
func BlockSetup(config Config) *Block {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Block{
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		DB:               config.DB,
		Auth:             config.Auth,
	}

}

type Filter struct {
	Keyword string
}

type TblBlock struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	BlockDescription string    `gorm:"type:text"`
	BlockContent     string    `gorm:"type:text"`
	BlockCss         string    `gorm:"type:text"`
	CoverImage       string    `gorm:"type:character varying"`
	IconImage        string    `gorm:"type:character varying"`
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

/* Collection List*/
// pass limit , offset get collectionlist
func (blocks *Block) CollectionList(filter Filter, tenantid int) (collectionlists []TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	collectionlist, err := Blockmodel.CollectionLists(filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return collectionlist, nil

}

// Block list
func (blocks *Block) BlockList(filter Filter, tenantid int) (blocklists []TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	blocklist, err := Blockmodel.BlockLists(filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return blocklist, nil

}
