package blocks

import (
	"gorm.io/gorm"
)

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

	query := DB.Debug().Table("tbl_blocks").Joins("inner join tbl_block_tags on tbl_block_tags.id=tbl_block.id").Where("tbl_blocks.tenant_id=? ", tenantid).Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?)) or LOWER(TRIM(tag_name)) LIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_blocks.created_by=?", Blockmodel.UserId)
	}

	query.Find(&block)

	return block, err

}
