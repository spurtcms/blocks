package blocks

import (
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword string
}

type BlockCreation struct {
	Title            string
	BlockDescription string
	BlockContent     string
	BlockCss         string
	CoverImage       string
	IconImage        string
	TenantId         int
	CreatedBy        int
}

type MasterTagCreate struct {
	Name      string
	TenantId  int
	CreatedBy int
}

type CreateTag struct {
	BlockId   int
	TagId     int
	TagName   string
	TenantId  int
	CreatedBy int
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

// pass limit , offset get collectionlist
func (Blockmodel BlockModel) CollectionLists(filter Filter, DB *gorm.DB, tenantid int) (collection []TblBlock, err error) {

	query := DB.Debug().Table("tbl_blocks").Select("tbl_blocks.*").Joins("inner join tbl_block_collections on tbl_block_collections.block_id = tbl_blocks.id").Where("tbl_block_collections.tenant_id=? and tbl_block_collections.is_deleted = ? and tbl_block_collections.user_id = ?", tenantid, 0, Blockmodel.UserId).Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) or LOWER(TRIM(name)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_collections.user_id=?", Blockmodel.UserId)
	}

	query.Find(&collection)

	return collection, err

}

// pass limit , offset get blocklist
func (Blockmodel BlockModel) BlockLists(filter Filter, DB *gorm.DB, tenantid int) (block []TblBlock, err error) {

	query := DB.Debug().Table("tbl_blocks").Joins("inner join tbl_block_tags on tbl_block_tags.id=tbl_blocks.id").Where("tbl_blocks.tenant_id=? ", tenantid).Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) or LOWER(TRIM(tag_name)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_blocks.created_by=?", Blockmodel.UserId)
	}

	query.Find(&block)

	return block, err

}

func (Blockmodel BlockModel) CreateBlocks(block TblBlock, DB *gorm.DB) (cblock TblBlock, err error) {

	if err := DB.Table("tbl_blocks").Create(block).Error; err != nil {
		return TblBlock{}, err
	}
	return cblock, nil
}

func (Blockmodel BlockModel) TagNameCheck(tagname string, DB *gorm.DB) (tags TblBlockMstrTag, err error) {

	if err := DB.Table("tbl_block_mstr_tags").Where("LOWER(TRIM(name))=LOWER(TRIM(?))", tagname).First(&tags).Error; err != nil {

		return TblBlockMstrTag{}, err

	}

	return tags, nil
}

func (Blockmodel BlockModel) CreateMasterTag(mstrtags TblBlockMstrTag, DB *gorm.DB) (mstrtag TblBlockMstrTag, err error) {

	if err := DB.Table("tbl_block_mstr_tags").Create(&mstrtags).Error; err != nil {
		return TblBlockMstrTag{}, err

	}

	return mstrtags, nil
}

func (Blockmodel BlockModel) CreateBlockTag(tags TblBlockTags, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_tags").Create(&tags).Error; err != nil {
		return err

	}

	return nil
}
