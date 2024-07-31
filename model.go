package blocks

import (
	"gorm.io/gorm"
)

// pass limit , offset get collectionlist
func (Blockmodel BlockModel) CollectionLists(filter Filter, DB *gorm.DB, tenantid int) (collection []TblBlock, err error) {

	query := DB.Debug().Table("tbl_blocks").Select("tbl_blocks.*").Joins("inner join tbl_block_collections on tbl_block_collections.block_id = tbl_blocks.id").Where("tbl_block_collections.tenant_id=? and tbl_block_collections.is_deleted = ? and tbl_block_collections.user_id = ?", tenantid, 0, Blockmodel.UserId).Order("tbl_blocks.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_collections.user_id=?", Blockmodel.UserId)
	}

	query.Find(&collection)

	return collection, err

}
