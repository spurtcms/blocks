package blocks

import (
	"gorm.io/gorm"
)

// pass limit , offset get collectionlist
func (Blockmodel BlockModel) CollectionList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int) (collection []TblBlock, totalcollection int64, err error) {

	query := DB.Table("tbl_blocks").Select("tbl_blocks.*").Joins("inner join tbl_block_collections on tbl_block_collections.block_id = tbl_block.id").Where("tbl_block_collection.tenant_id=? and tbl_block_collections.is_deleted = ? and tbl_block_collection.user_id ", tenantid, 0, Blockmodel.UserId).Order("tbl_block.id desc")

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(title)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	if Blockmodel.DataAccess == 1 {

		query = query.Where("tbl_block_collection.user_id=?", Blockmodel.UserId)
	}

	if limit != 0 {

		query.Offset(offset).Limit(limit).Order("id desc").Find(&collection)

		return collection, 0, err

	} else {

		query.Find(&collection).Count(&totalcollection)

		return collection, totalcollection, err
	}
}
