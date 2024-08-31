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
func (blocks *Block) BlockList(limit, offset int, filter Filter, tenantid int) (blocklists []TblBlock, countblock int64, defaultlists []TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, []TblBlock{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	var tenantbasedblock []TblBlock

	var DefaultCollectionList int

	tenantblock, _ := Blockmodel.GetUserBlocks(tenantbasedblock, tenantid, blocks.DB)

	for _, val := range tenantblock {
		if val.Id == 0 {
			DefaultCollectionList = 0
		} else {
			DefaultCollectionList = 1
		}
	}

	blocklist, _, err := Blockmodel.BlockLists(limit, offset, filter, blocks.DB, tenantid, DefaultCollectionList)

	_, count, _ := Blockmodel.BlockLists(0, 0, filter, blocks.DB, tenantid, DefaultCollectionList)

	var deblock []TblBlock

	defaultlist, _ := Blockmodel.GetBlocks(deblock, filter, blocks.DB)

	if err != nil {

		fmt.Println(err)
	}

	return blocklist, count, defaultlist, nil

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
	block.IsActive = Bc.IsActive
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
func (blocks *Block) RemoveBlock(id int, tenantid int) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	Blockmodel.UserId = blocks.UserId

	var block TblBlock

	block.Id = id
	block.TenantId = tenantid
	block.DeletedBy = blocks.UserId
	block.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	block.IsDeleted = 1

	err := Blockmodel.DeleteBlock(block, blocks.DB)

	if err != nil {

		return err
	}

	return nil

}

// Check collection  already exists
func (blocks *Block) CheckCollection(blockid, user_id, tenantid int) (flg TblBlockCollection, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockCollection{}, AuthErr
	}

	var block TblBlockCollection

	tag, err1 := Blockmodel.CheckCollectionById(block, blockid, user_id, tenantid, blocks.DB)

	if err1 != nil {
		return TblBlockCollection{}, err
	}

	return tag, nil

}

// check block title is alreay exists

func (blocks *Block) CheckTitleInBlock(title string, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return false, AuthErr
	}

	var tblblocks TblBlock

	err := Blockmodel.CheckTitleInBlock(&tblblocks, title, blocks.DB, tenantid)

	if err != nil {

		return false, err

	}

	return true, nil
}

// last 10 days la add pana block count
func (blocks *Block) DashBoardBlockCount(tenantid int) (Totalcount int, lcount int, err error) {

	autherr := AuthandPermission(blocks)

	if autherr != nil {

		return 0, 0, autherr
	}

	allblockcount, err := Blockmodel.AllBlockCount(blocks.DB, tenantid)

	if err != nil {

		return 0, 0, err
	}

	lblockcount, err := Blockmodel.NewBlockCount(blocks.DB, tenantid)

	if err != nil {

		return 0, 0, err
	}

	return int(allblockcount), int(lblockcount), nil
}

// IsActive functionality in block

func (blocks *Block) BlockIsActive(id int, status int, modifiedby int, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return false, AuthErr
	}

	var block TblBlock

	block.ModifiedBy = modifiedby

	block.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Blockmodel.BlcokIsActive(block, id, status, blocks.DB, tenantid)

	if err != nil {

		return false, err

	}
	return true, nil

}

// Edit Functionality
func (blocks *Block) BlockEdit(id int, tenantid int) (blockdata TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlock{}, AuthErr
	}

	var block TblBlock
	blockdetails, err := Blockmodel.BlockEdit(block, id, blocks.DB, tenantid)

	if err != nil {

		return TblBlock{}, err

	}
	return blockdetails, nil

}
