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

/* Collection List*/
// pass limit , offset get collectionlist
func (blocks *Block) CollectionList(limit, offset int, filter Filter, tenantid int) (collectionlists []TblBlock, totalcollection int64, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	var collection []TblBlockCollection

	usercollection, _ := Blockmodel.GetCollectionByUserId(collection, Blockmodel.UserId, blocks.DB)

	var id []int

	for _, val := range usercollection {

		id = append(id, val.BlockId)

	}

	collectionlist, _, err := Blockmodel.CollectionLists(limit, offset, filter, blocks.DB, tenantid, id)

	_, totalcount, _ := Blockmodel.CollectionLists(0, 0, filter, blocks.DB, tenantid, id)

	if err != nil {

		fmt.Println(err)
	}
	return collectionlist, totalcount, nil
}

// Block list
func (blocks *Block) BlockList(limit, offset int, filter Filter, tenantid int, work string) (blocklists []TblBlock, countblock int64, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	blocklist, _, err := Blockmodel.BlockLists(limit, offset, filter, blocks.DB, tenantid, work)

	_, count, _ := Blockmodel.BlockLists(0, 0, filter, blocks.DB, tenantid, work)

	if err != nil {

		fmt.Println(err)
	}

	return blocklist, count, nil

}

// Create Blog
func (blocks *Block) CreateBlock(Bc BlockCreation) (createblocks TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlock{}, AuthErr
	}

	var block TblBlock

	block.TenantId = Bc.TenantId
	block.BlockContent = Bc.BlockContent
	block.BlockCss = Bc.BlockCss
	block.BlockDescription = Bc.BlockDescription
	block.CoverImage = Bc.CoverImage
	block.Title = Bc.Title
	block.Prime = Bc.Prime
	block.CreatedBy = Bc.CreatedBy
	block.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	createblock, err := Blockmodel.CreateBlocks(block, blocks.DB)

	if err != nil {

		fmt.Println(err)
	}

	return createblock, nil

}

// Check tag name is already exists
func (blocks *Block) CheckTagName(tagname string) (flg TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockMstrTag{}, AuthErr
	}

	var block TblBlockMstrTag

	tag, err1 := Blockmodel.TagNameCheck(tagname, blocks.DB, block)

	if err1 != nil {
		return TblBlockMstrTag{}, err
	}

	return tag, nil

}

// Create Master tag
func (blocks *Block) CreateMasterTag(MstrTag MasterTagCreate) (createtags TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockMstrTag{}, AuthErr
	}

	var tags TblBlockMstrTag

	tags.TagTitle = MstrTag.TagTitle
	tags.TenantId = MstrTag.TenantId
	tags.CreatedBy = MstrTag.CreatedBy
	tags.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	createtag, err := Blockmodel.CreateMasterTag(tags, blocks.DB)

	if err != nil {
		return TblBlockMstrTag{}, err
	}

	return createtag, nil

}

// Create tag
func (blocks *Block) CreateTag(Tag CreateTag) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var tags TblBlockTags
	tags.BlockId = Tag.BlockId
	tags.TagId = Tag.TagId
	tags.TagName = Tag.TagName
	tags.TenantId = Tag.TenantId
	tags.CreatedBy = Tag.CreatedBy
	tags.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Blockmodel.CreateBlockTag(tags, blocks.DB)

	if err != nil {
		return err
	}

	return nil

}

// Create collection
func (blocks *Block) BlockCollection(Collections CreateCollection) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var collection TblBlockCollection

	collection.UserId = Collections.UserId

	collection.BlockId = Collections.BlockId

	collection.TenantId = Collections.TenantId

	err := Blockmodel.CreateBlockCollection(collection, blocks.DB)

	if err != nil {
		return err
	}

	return nil

}

// Get tag list
func (blocks *Block) TagList(filter Filter, tenantid int) (taglists []TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlockMstrTag{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	taglist, err := Blockmodel.TagLists(filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return taglist, nil

}

// Remove Collection
func (blocks *Block) RemoveCollection(id int, tenantid int) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	Blockmodel.UserId = blocks.UserId

	var collection TblBlockCollection

	collection.BlockId = id
	collection.TenantId = tenantid
	collection.DeletedBy = blocks.UserId
	collection.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	collection.IsDeleted = 1

	err := Blockmodel.DeleteCollection(collection, blocks.DB)

	if err != nil {

		return err
	}

	return nil

}
