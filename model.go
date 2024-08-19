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
	Prime            int
	TenantId         int
	CreatedBy        int
}

type MasterTagCreate struct {
	TagTitle  string
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

type CreateCollection struct {
	UserId   int
	BlockId  int
	TenantId int
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
	Prime            int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ProfileImagePath string    `gorm:"<-:false"`
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
	TagTitle  string    `gorm:"type:character varying"`
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

// get collectionlist
func (Blockmodel BlockModel) CollectionLists(filter Filter, DB *gorm.DB, tenantid int, blockid []int) (collection []TblBlock, err error) {

	query := DB.Debug().Table("tbl_blocks").Select("tbl_blocks.*").Joins("inner join tbl_block_collections on tbl_block_collections.block_id = tbl_blocks.id").Joins("inner join tbl_block_tags on tbl_block_tags.block_id = tbl_blocks.id").Where("tbl_block_collections.is_deleted = ?", 0).Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) or LOWER(TRIM(tbl_block_tags.tag_name)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if blockid != nil {

		query = query.Where("tbl_block_collections.user_id = ?", Blockmodel.UserId)
	} else {
		query = query.Where("tbl_block_collections.tenant_id = ? or tbl_block_collections is NULL", tenantid)
	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_collections.user_id=?", Blockmodel.UserId)
	}

	query.Find(&collection)

	return collection, err

}

// get blocklist
func (Blockmodel BlockModel) BlockLists(Limit, Offset int, filter Filter, DB *gorm.DB, tenantid int, work string) (block []TblBlock, Totalblock int64, err error) {

	query := DB.Select("tbl_blocks.*,tbl_users.profile_image_path as profile_image_path").Debug().Table("tbl_blocks").Joins("inner join tbl_users on tbl_users.id = tbl_blocks.created_by")

	if work != "" {
		query = query.Where("tbl_blocks.created_by =? and (tbl_blocks.tenant_id=? or tbl_blocks.tenant_id is Null  ) ", Blockmodel.UserId, tenantid).Order("tbl_blocks.id desc")

	} else {
		query = query.Where("tbl_blocks.tenant_id=? or tbl_blocks.tenant_id is NULL ", tenantid).Order("tbl_blocks.id desc")

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?))  ", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_blocks.created_by=?", Blockmodel.UserId)
	}

	if Limit != 0 {

		query.Limit(Limit).Offset(Offset).Find(&block)

		return block, 0, err

	}

	query.Find(&block).Count(&Totalblock)

	return block, Totalblock, err

}

// Create blocks
func (Blockmodel BlockModel) CreateBlocks(block TblBlock, DB *gorm.DB) (cblock TblBlock, err error) {

	if err := DB.Debug().Table("tbl_blocks").Create(&block).Error; err != nil {
		return TblBlock{}, err
	}
	return block, nil
}

// Check tag name is already exists
func (Blockmodel BlockModel) TagNameCheck(tagname string, DB *gorm.DB, tags TblBlockMstrTag) (tag TblBlockMstrTag, err error) {

	if err := DB.Table("tbl_block_mstr_tags").Where("LOWER(TRIM(tag_title))=LOWER(TRIM(?))", tagname).First(&tags).Error; err != nil {

		return TblBlockMstrTag{}, err

	}

	return tags, nil
}

// Create master tag
func (Blockmodel BlockModel) CreateMasterTag(mstrtags TblBlockMstrTag, DB *gorm.DB) (mstrtag TblBlockMstrTag, err error) {

	if err := DB.Table("tbl_block_mstr_tags").Create(&mstrtags).Error; err != nil {
		return TblBlockMstrTag{}, err

	}

	return mstrtags, nil
}

// Mapping tags in block
func (Blockmodel BlockModel) CreateBlockTag(tags TblBlockTags, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_tags").Create(&tags).Error; err != nil {
		return err

	}

	return nil
}

// Create block collection
func (Blockmodel BlockModel) CreateBlockCollection(collection TblBlockCollection, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_block_collections").Create(&collection).Error; err != nil {
		return err

	}

	return nil
}

// get taglist
func (Blockmodel BlockModel) TagLists(filter Filter, DB *gorm.DB, tenantid int) (tags []TblBlockMstrTag, err error) {

	// query := DB.Table("tbl_block_mstr_tags").Joins("inner join tbl_block_tags on tbl_block_tags.tag_id =tbl_block_mstr_tags.id").Joins("inner join tbl_blocks on tbl_blocks.id = tbl_block_tags.block_id").Where("tbl_block_mstr_tags.tenant_id=? ", tenantid).Order("tbl_block_mstr_tags.id desc")

	query := DB.Table("tbl_block_mstr_tags").Where("tbl_block_mstr_tags.tenant_id=? or tbl_block_mstr_tags.tenant_id is NULL ", tenantid).Order("tbl_block_mstr_tags.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_block_mstr_tags.name)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_mstr_tags.created_by=?", Blockmodel.UserId)
	}

	query.Find(&tags)

	return tags, err

}

// Delete Collection
func (Blockmodel BlockModel) DeleteCollection(collection TblBlockCollection, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_collections").Where("block_id = ? and (tenant_id = ? or tenant_id is NULL ) ", collection.BlockId, collection.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": collection.IsDeleted, "deleted_by": collection.DeletedBy, "deleted_on": collection.DeletedOn}).Error; err != nil {

		return err

	}

	return nil

}

// Get collection based on userid

func (Blockmodel BlockModel) GetCollectionByUserId(collections []TblBlockCollection, userid int, DB *gorm.DB) (collection []TblBlockCollection, err error) {

	if err := DB.Debug().Table("tbl_block_collections").Where("user_id = ? ", userid).Find(&collections).Error; err != nil {

		return []TblBlockCollection{}, err

	}

	return collections, nil

}
