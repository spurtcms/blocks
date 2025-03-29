package blocks

import (
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword   string
	Channelid string
}

type BlockCreation struct {
	Title            string
	BlockDescription string
	BlockContent     string
	BlockCss         string
	CoverImage       string
	IconImage        string
	Prime            int
	TenantId         string
	CreatedBy        int
	ModifiedBy       int
	IsActive         int
	ChannelName      string
	ChannelId        string
}

type MasterTagCreate struct {
	TagTitle  string
	TenantId  string
	CreatedBy int
}

type CreateTag struct {
	BlockId   int
	TagId     int
	TagName   string
	TenantId  string
	CreatedBy int
}

type CreateCollection struct {
	UserId   int
	BlockId  int
	TenantId string
}

type TblBlock struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	SlugName         string    `gorm:"type:character varying"`
	ChannelSlugname  string    `gorm:"type:character varying"`
	BlockDescription string    `gorm:"type:text"`
	BlockContent     string    `gorm:"type:text"`
	BlockCss         string    `gorm:"type:text"`
	CoverImage       string    `gorm:"type:character varying"`
	IconImage        string    `gorm:"type:character varying"`
	TenantId         string    `gorm:"type:character varying"`
	Prime            int       `gorm:"type:integer"`
	IsActive         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedBy       int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	ProfileImagePath string    `gorm:"<-:false"`
	FirstName        string    `gorm:"<-:false"`
	LastName         string    `gorm:"<-:false"`
	Username         string    `gorm:"<-:false"`
	TagValueArr      []string  `gorm:"-"`
	TagValue         string    `gorm:"<-:false;"`
	NameString       string    `gorm:"-"`
	Actions          string    `gorm:"<-:false"`
	CreatedDate      string    `gorm:"-:migration;<-:false"`
	ModifiedDate     string    `gorm:"-:migration;<-:false"`
	ChannelID        string    `gorm:"type:character varying"`
	ChannelNames     []string  `gorm:"-"`
}

type TblBlockTags struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	BlockId   int       `gorm:"type:integer"`
	TagId     int       `gorm:"type:integer"`
	TagName   string    `gorm:"type:character varying"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}

type TblBlockMstrTag struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	TagTitle  string    `gorm:"type:character varying"`
	TenantId  string    `gorm:"type:character varying"`
	CreatedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy int       `gorm:"type:integer"`
}

type TblBlockCollection struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	UserId    int       `gorm:"type:integer"`
	BlockId   int       `gorm:"type:integer"`
	TenantId  string    `gorm:"type:character varying"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}

// get collectionlist
func (Blockmodel BlockModel) CollectionLists(filter Filter, DB *gorm.DB, tenantid string, chid string) (collection []TblBlock, count int64, err error) {

	query := DB.Debug().Table("tbl_blocks").Select("tbl_blocks.id,tbl_blocks.is_active,tbl_blocks.title,tbl_blocks.block_description,tbl_blocks.block_content,tbl_blocks.block_css,tbl_blocks.cover_image,tbl_blocks.created_by,tbl_users.profile_image_path as profile_image_path").Joins("inner join tbl_users on tbl_users.id = tbl_blocks.created_by").Where("tbl_blocks.tenant_id=? and tbl_blocks.is_deleted=0", tenantid).Group("tbl_blocks.id").Group("tbl_users.profile_image_path").Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_blocks.title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	// if chid != "" {

	// 	query = query.Where("string_to_array((tbl_blocks.channel_id), ',') && string_to_array(?, ',')", (chid))
	// }

	query.Find(&collection).Count(&count)

	return collection, count, err

}
// // get blocklist
// func (Blockmodel BlockModel) BlockLists(limit, offset int, filter Filter, DB *gorm.DB, tenantid string) (block []TblBlock, Totalblock int64, err error) {

// 	dbName := DB.Dialector.Name()

// 	var query *gorm.DB

// 	if dbName == "postgres" {

// 		query = DB.Select("tbl_blocks.*,max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username, STRING_AGG(tbl_block_tags.tag_name, ', ') as tag_value ,(case when (select id from tbl_block_collections where tbl_block_collections.block_id = tbl_blocks.id and is_deleted = 0 limit 1) is not null then 'true' else 'false' end ) as actions ").Table("tbl_blocks")

// 	} else if dbName == "mysql" {

// 		query = DB.Select("tbl_blocks.*,max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username, GROUP_CONCAT(tbl_block_tags.tag_name ORDER BY tbl_block_tags.tag_name SEPARATOR ', ') AS tag_value ,(case when (select id from tbl_block_collections where tbl_block_collections.block_id = tbl_blocks.id and is_deleted = 0 limit 1) is not null then 'true' else 'false' end ) as actions ").Table("tbl_blocks")
// 	}

// 	query = query.Joins("inner join tbl_block_tags on tbl_block_tags.block_id = tbl_blocks.id").Joins("left join tbl_block_collections on tbl_block_collections.block_id = tbl_blocks.id").Joins("inner join tbl_users on case when tbl_block_collections.id is not null then tbl_users.id = tbl_block_collections.user_id else tbl_users.id = tbl_blocks.created_by end")

// 	if dbName == "postgres" {

// 		query = query.Where("tbl_blocks.is_deleted = ? and (tbl_block_collections.is_deleted = ? or tbl_block_collections is NULL ) and tbl_blocks.tenant_id = ?  and (tbl_blocks.created_by = ? or tbl_block_collections.user_id = ?)and tbl_block_tags.is_deleted = ?  ", 0, 0, tenantid, Blockmodel.UserId, Blockmodel.UserId, 0)

// 	} else if dbName == "mysql" {

// 		query = query.Where("tbl_blocks.is_deleted = ? and (tbl_block_collections.is_deleted = ? or tbl_block_collections.id is NULL ) and tbl_blocks.tenant_id = ?  and (tbl_blocks.created_by = ? or tbl_block_collections.user_id = ?)and tbl_block_tags.is_deleted = ?  ", 0, 0, tenantid, Blockmodel.UserId, Blockmodel.UserId, 0)
// 	}

// 	query = query.Group("tbl_blocks.id").Order("tbl_blocks.id desc")

// 	if filter.Keyword != "" {

// 		query = query.Where("LOWER(TRIM(tbl_blocks.title)) LIKE LOWER(TRIM(?))  ", "%"+filter.Keyword+"%")

// 	}

// 	if filter.Channelid != "" {
// 		query = query.Where("string_to_array(tbl_blocks.channel_id::TEXT, ',', '') @> ARRAY[?]::TEXT[]", filter.Channelid)
// 	}

// 	if Blockmodel.DataAccess == 1 {

// 		query = query.Where("tbl_blocks.created_by=?", Blockmodel.UserId)
// 	}

// 	if limit != 0 {

// 		query.Limit(limit).Offset(offset).Debug().Find(&block)

// 		return block, 0, err

// 	}

// 	query.Find(&block).Count(&Totalblock)

// 	return block, Totalblock, err

// }
// get blocklist
func (Blockmodel BlockModel) BlockLists(limit, offset int, filter Filter, DB *gorm.DB, tenantid string) (block []TblBlock, Totalblock int64, err error) {

	var query *gorm.DB

	query = DB.Select("tbl_blocks.*, max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username").
		Table("tbl_blocks").
		Joins("INNER JOIN tbl_users ON tbl_users.id = tbl_blocks.created_by")

	query = query.Where("tbl_blocks.is_deleted = ? AND tbl_blocks.tenant_id = ?", 0, tenantid)

	if Blockmodel.DataAccess == 1 {
		query = query.Where("tbl_blocks.created_by = ?", Blockmodel.UserId)
	}

	if filter.Keyword != "" {
		query = query.Where("LOWER(TRIM(tbl_blocks.title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")
	}

	if filter.Channelid != "" {
		query = query.Where("tbl_blocks.channel_id = ?", filter.Channelid)
	}

	query = query.Group("tbl_blocks.id").Order("tbl_blocks.id DESC")

	if limit != 0 {
		err = query.Limit(limit).Offset(offset).Find(&block).Error
		return block, 0, err
	}

	err = query.Find(&block).Count(&Totalblock).Error
	return block, Totalblock, err
}

// get dafault blocklist
func (Blockmodel BlockModel) DefaultBlockLists(limit, offset int, filter Filter, DB *gorm.DB, tenantid string) (dafaultblock []TblBlock, Totaldefaultblock int64, err error) {

	dbName := DB.Dialector.Name()

	var query *gorm.DB

	if dbName == "postgres" {

		query = DB.Select("tbl_blocks.*,max(tbl_users.id),max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username, STRING_AGG(tbl_block_tags.tag_name, ', ') as tag_value")

	} else if dbName == "mysql" {

		query = DB.Select("tbl_blocks.*,GROUP_CONCAT(tbl_block_tags.tag_name ORDER BY tbl_block_tags.tag_name SEPARATOR ', ') AS tag_value ,max(tbl_users.id),max(tbl_users.first_name) as first_name,max(tbl_users.last_name)  as last_name, max(tbl_users.profile_image_path) as profile_image_path, max(tbl_users.username)  as username")
	}

	query = query.Table("tbl_blocks").Joins("inner join tbl_users on tbl_users.id = tbl_blocks.created_by").Joins("inner join tbl_block_tags ON tbl_block_tags.block_id = tbl_blocks.id").Where("tbl_blocks.tenant_id is NULL and tbl_blocks.is_deleted=0 and block_description='spurtcms'").Group("tbl_blocks.id").Order("tbl_blocks.id desc")
	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_blocks.title)) LIKE LOWER(TRIM(?))  ", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_blocks.created_by=?", Blockmodel.UserId)
	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Debug().Find(&dafaultblock)

		return dafaultblock, 0, err

	}

	query.Find(&dafaultblock).Count(&Totaldefaultblock)

	return dafaultblock, Totaldefaultblock, err

}

// Create blocks
func (Blockmodel BlockModel) CreateBlocks(block TblBlock, DB *gorm.DB) (cblock TblBlock, err error) {

	if err := DB.Table("tbl_blocks").Create(&block).Error; err != nil {

		return TblBlock{}, err
	}
	return block, nil
}

// Check tag name is already exists
func (Blockmodel BlockModel) TagNameCheck(tagname string, DB *gorm.DB, tags TblBlockMstrTag, tenantid string) (tag TblBlockMstrTag, err error) {

	if err := DB.Table("tbl_block_mstr_tags").Where("LOWER(TRIM(tag_title))=LOWER(TRIM(?)) and tenant_id = ? ", tagname, tenantid).First(&tags).Error; err != nil {

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

	if err := DB.Table("tbl_block_collections").Create(&collection).Error; err != nil {
		return err

	}

	return nil
}

// get taglist
func (Blockmodel BlockModel) TagLists(filter Filter, DB *gorm.DB, tenantid string) (tags []TblBlockMstrTag, err error) {

	query := DB.Table("tbl_block_mstr_tags").Where("tbl_block_mstr_tags.tenant_id = ? ", tenantid).Order("tbl_block_mstr_tags.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_block_mstr_tags.tag_title)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_mstr_tags.created_by=?", Blockmodel.UserId)
	}

	query.Find(&tags)

	return tags, err

}

// Delete Collection
func (Blockmodel BlockModel) DeleteBlock(block TblBlock, DB *gorm.DB) error {

	if err := DB.Table("tbl_blocks").Where("id = ? and tenant_id = ?  ", block.Id, block.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": block.IsDeleted, "deleted_by": block.DeletedBy, "deleted_on": block.DeletedOn}).Error; err != nil {

		return err

	}

	return nil

}

// Check collection already exists

func (Blockmodel BlockModel) CheckCollectionById(collections TblBlockCollection, blockid, userid int, tenantid string, DB *gorm.DB) (flg bool, err error) {

	if err := DB.Table("tbl_block_collections").Where("block_id = ? and user_id = ? and tenant_id = ?  and is_deleted = 0  ", blockid, userid, tenantid).First(&collections).Error; err != nil {

		return false, err

	}

	return true, nil

}

// func (Blockmodel BlockModel) GetBlocks(id int, DB *gorm.DB, Blocks *TblBlock) error {

// 	if err := DB.Table("tbl_blocks").Where("id=? and tenant_id is NULL", id).Debug().First(&Blocks).Error; err != nil {

// 		return err
// 	}
// 	return nil
// }

func (Blockmodel BlockModel) GetBlocks(id int, DB *gorm.DB, Blocks *TblBlock, tenantid string) error {

	if err := DB.Table("tbl_blocks").
		Select("tbl_blocks.*, tbl_channels.id as channel_id").
		Joins("INNER JOIN tbl_channels ON tbl_blocks.channel_Slugname = tbl_channels.channel_name").
		Where("tbl_blocks.id = ? and tbl_blocks.tenant_id IS NULL and tbl_channels.tenant_id=?", id, tenantid).
		Debug().Find(&Blocks).Error; err != nil {
		return err
	}

	return nil
}

func (Blockmodel BlockModel) AddToMycollection(Block TblBlock, DB *gorm.DB) error {

	if err := DB.Table("tbl_blocks").Create(&Block).Error; err != nil {

		return err
	}
	return nil

}

// check block title is alreay exists

func (Blockmodel BlockModel) CheckTitleInBlock(block *TblBlock, title string, DB *gorm.DB, id int, tenantid string) error {

	if id == 0 {
		if err := DB.Table("tbl_blocks").Where("LOWER(TRIM(title))=LOWER(TRIM(?)) and tenant_id = ?  and is_deleted = 0", title, tenantid).First(&block).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_blocks").Where("LOWER(TRIM(title))=LOWER(TRIM(?)) and tenant_id = ? and is_deleted = 0 and tbl_blocks.id not in (?)", title, tenantid, id).First(&block).Error; err != nil {

			return err
		}
	}

	return nil
}

// get delete collection blockid

func (Blockmodel BlockModel) GetUserBlocks(blocks []TblBlock, tenantid string, DB *gorm.DB) (block []TblBlock, err error) {

	if err := DB.Table("tbl_blocks").Where("tbl_blocks.is_deleted = 0 and tenant_id = ? ", tenantid).Find(&blocks).Error; err != nil {

		return []TblBlock{}, err

	}

	return blocks, nil

}

//Total Block Count

func (Blockmodel BlockModel) AllBlockCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_blocks").Where("tbl_blocks.is_deleted = 0 and  tenant_id = ?", tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil

}

// Last 10 days block Count
func (Blockmodel BlockModel) NewBlockCount(DB *gorm.DB, tenantid string) (count int64, err error) {

	if err := DB.Table("tbl_blocks").Where("created_on >=? and  tenant_id=? and is_deleted = 0", time.Now().AddDate(0, 0, -10), tenantid).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil

}

// Block is active

func (Blockmodel BlockModel) BlcokIsActive(blockstatus TblBlock, id int, status int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_blocks").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_active": status, "modified_by": blockstatus.ModifiedBy, "modified_on": blockstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// Edit functionality
func (Blockmodel BlockModel) BlockEdit(id int, DB *gorm.DB, tenantid string) (blockdata TblBlock, err error) {

	query := DB.Select("tbl_blocks.*").Table("tbl_blocks")

	if err := query.Where("tbl_blocks.id = ? AND tbl_blocks.tenant_id = ?", id, tenantid).First(&blockdata).Error; err != nil {
		return TblBlock{}, err
	}

	return blockdata, nil
}

// Update Functionality

func (Blockmodel BlockModel) UpdateBlock(block TblBlock, id int, DB *gorm.DB) error {

	if block.CoverImage != "" {
		if err := DB.Table("tbl_blocks").Where("tbl_blocks.id=? and  tbl_blocks.tenant_id=?", id, block.TenantId).UpdateColumns(map[string]interface{}{"title": block.Title, "block_content": block.BlockContent, "channel_slugname": block.ChannelSlugname, "channel_id": block.ChannelID, "is_active": block.IsActive, "modified_by": block.ModifiedBy, "modified_on": block.ModifiedOn, "prime": block.Prime, "cover_image": block.CoverImage}).Error; err != nil {

			return err
		}

	} else {
		if err := DB.Table("tbl_blocks").Where("tbl_blocks.id=? and tbl_blocks.tenant_id=?", id, block.TenantId).UpdateColumns(map[string]interface{}{"title": block.Title, "block_content": block.BlockContent, "channel_slugname": block.ChannelSlugname, "channel_id": block.ChannelID, "is_active": block.IsActive, "modified_by": block.ModifiedBy, "modified_on": block.ModifiedOn, "prime": block.Prime}).Error; err != nil {

			return err
		}

	}

	return nil
}

// Delete Functionality

func (Blockmodel BlockModel) DeleteBlockCollection(blockcollection TblBlockCollection, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_collections").Where("block_id = ? and tenant_id = ? ", blockcollection.BlockId, blockcollection.TenantId).Delete(&blockcollection).Error; err != nil {

		return err

	}

	return nil

}

// Tag Delete functionality

func (Blockmodel BlockModel) DeleteTag(tag TblBlockTags, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_tags").Where("block_id = ? and tag_name = ? and tenant_id = ?  ", tag.BlockId, tag.TagName, tag.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": tag.IsDeleted, "deleted_by": tag.DeletedBy, "deleted_on": tag.DeletedOn}).Error; err != nil {

		return err

	}

	return nil

}

// Tag Delete based on block delete functionality

func (Blockmodel BlockModel) BlockDeleteTag(tags TblBlockTags, DB *gorm.DB) error {

	if err := DB.Table("tbl_block_tags").Where("block_id = ? and tenant_id = ?", tags.BlockId, tags.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": tags.IsDeleted, "deleted_by": tags.DeletedBy, "deleted_on": tags.DeletedOn}).Error; err != nil {

		return err

	}

	return nil

}
