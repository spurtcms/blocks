package blocks

import (
	"fmt"
	"strings"

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
// pass limit, offset get collectionlist
func (blocks *Block) CollectionList(filter Filter, tenantid string, channelid string) (collectionlists []TblBlock, count int64, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	collectionlist, blockcount, err := Blockmodel.CollectionLists(filter, blocks.DB, tenantid, channelid)

	if err != nil {

		fmt.Println(err)
	}
	return collectionlist, blockcount, nil
}

// Block list
func (blocks *Block) BlockList(limit, offset int, filter Filter, tenantid string) (blocklists []TblBlock, countblock int64, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	blocklist, _, err := Blockmodel.BlockLists(limit, offset, filter, blocks.DB, tenantid)

	_, count, _ := Blockmodel.BlockLists(0, 0, filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return blocklist, count, nil

}

// Default Block list
func (blocks *Block) DefaultBlockList(limit, offset int, filter Filter, tenantid string) (blocklists []TblBlock, countblock int64, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, 0, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	blocklist, _, err := Blockmodel.DefaultBlockLists(limit, offset, filter, blocks.DB, tenantid)

	_, count, _ := Blockmodel.DefaultBlockLists(0, 0, filter, blocks.DB, tenantid)

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
	block.IsActive = Bc.IsActive
	block.ChannelSlugname = Bc.ChannelName
	block.ChannelID = Bc.ChannelId
	block.SlugName = strings.ToLower(strings.ReplaceAll(Bc.Title, " ", "-"))

	createblock, err := Blockmodel.CreateBlocks(block, blocks.DB)

	if err != nil {

		fmt.Println(err)
	}

	return createblock, nil

}

// Check tag name is already exists
func (blocks *Block) CheckTagName(tagname string, tenantid string) (flg TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockMstrTag{}, AuthErr
	}

	var block TblBlockMstrTag

	tag, err1 := Blockmodel.TagNameCheck(tagname, blocks.DB, block, tenantid)

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
func (blocks *Block) TagList(filter Filter, tenantid string) (taglists []TblBlockMstrTag, err error) {

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
func (blocks *Block) RemoveBlock(id int, tenantid string) error {

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

	var collection TblBlockCollection

	collection.BlockId = id
	collection.TenantId = tenantid
	collection.DeletedBy = blocks.UserId
	collection.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	collection.IsDeleted = 1

	err1 := Blockmodel.DeleteBlockCollection(collection, blocks.DB)

	if err1 != nil {
		return err1
	}

	var tags TblBlockTags

	tags.BlockId = id
	tags.TenantId = tenantid
	tags.DeletedBy = blocks.UserId
	tags.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tags.IsDeleted = 1

	err2 := Blockmodel.BlockDeleteTag(tags, blocks.DB)

	if err2 != nil {
		return err2
	}

	return nil

}

// Check collection  already exists
func (blocks *Block) CheckCollection(blockid, user_id int, tenantid string) (flg bool, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return false, AuthErr
	}

	var block TblBlockCollection

	flg, err1 := Blockmodel.CheckCollectionById(block, blockid, user_id, tenantid, blocks.DB)

	if err1 != nil {
		return false, err
	}

	return flg, nil

}

func (blocks *Block) Addblocktomycollecton(id int, tenantid string, userid int) (bool, error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return false, AuthErr
	}

	var Block TblBlock

	err := Blockmodel.GetBlocks(id, blocks.DB, &Block, tenantid)

	if Block.Id != 0 {

		if err != nil {
			fmt.Println("Add to mycollection contain error,line 338", err)
		}
		currenttime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		myblock := TblBlock{
			Title:            Block.Title,
			BlockDescription: Block.BlockDescription,
			BlockContent:     Block.BlockContent,
			ChannelSlugname:  Block.ChannelSlugname,
			BlockCss:         Block.BlockCss,
			IconImage:        Block.IconImage,
			CoverImage:       Block.CoverImage,
			Prime:            Block.Prime,
			IsActive:         Block.IsActive,
			CreatedOn:        currenttime,
			CreatedBy:        userid,
			ModifiedBy:       Block.ModifiedBy,
			DeletedOn:        Block.DeletedOn,
			DeletedBy:        Block.DeletedBy,
			IsDeleted:        Block.IsDeleted,
			TenantId:         tenantid,
			ChannelID:        Block.ChannelID,
		}

		// Blockmodel.AddToMycollection(myblock, blocks.DB)

		_, err2 := Blockmodel.CreateBlocks(myblock, blocks.DB)

		if err2 != nil {
			fmt.Println("block err", err)
		}

		// var block TblBlockMstrTag

		// tag, err1 := Blockmodel.TagNameCheck("default", blocks.DB, block, tenantid)

		// if err1 != nil {
		// 	return false, err
		// }

		// TagCreate := CreateTag{
		// 	BlockId:   blockdata.Id,
		// 	TagId:     tag.Id,
		// 	TagName:   tag.TagTitle,
		// 	CreatedBy: userid,
		// }

		// var tags TblBlockTags
		// tags.BlockId = TagCreate.BlockId
		// tags.TagId = TagCreate.TagId
		// tags.TagName = TagCreate.TagName
		// tags.TenantId = tenantid
		// tags.CreatedBy = TagCreate.CreatedBy
		// tags.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		// Blockmodel.CreateBlockTag(tags, blocks.DB)

		// var collection TblBlockCollection

		// collection.BlockId = blockdata.Id
		// collection.UserId = userid
		// collection.TenantId = tenantid
		// collection.DeletedBy = userid
		// collection.IsDeleted = blockdata.IsDeleted

		// Blockmodel.CreateBlockCollection(collection, blocks.DB)

		return true, nil

	} else {
		return false, nil
	}

}

// check block title is alreay exists

func (blocks *Block) CheckTitleInBlock(title string, id int, tenantid string) (bool, error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return false, AuthErr
	}

	var tblblocks TblBlock

	err := Blockmodel.CheckTitleInBlock(&tblblocks, title, blocks.DB, id, tenantid)

	if err != nil {

		return false, err

	}

	return true, nil
}

// last 10 days la add pana block count
func (blocks *Block) DashBoardBlockCount(tenantid string) (Totalcount int, lcount int, err error) {

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

func (blocks *Block) BlockIsActive(id int, status int, modifiedby int, tenantid string) (bool, error) {

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
func (blocks *Block) BlockEdit(id int, tenantid string) (blockdata TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlock{}, AuthErr
	}

	blockdetails, err := Blockmodel.BlockEdit(id, blocks.DB, tenantid)

	if err != nil {

		return TblBlock{}, err

	}
	return blockdetails, nil

}

// Update Functionality
func (blocks *Block) UpdateBlock(id int, updateblock BlockCreation) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var block TblBlock

	block.BlockContent = updateblock.BlockContent
	block.Title = updateblock.Title
	block.BlockCss = updateblock.BlockCss
	block.CoverImage = updateblock.CoverImage
	block.ModifiedBy = updateblock.ModifiedBy
	block.IsActive = updateblock.IsActive
	block.Prime = updateblock.Prime
	block.TenantId = updateblock.TenantId
	block.ChannelSlugname = updateblock.ChannelName
	block.ChannelID = updateblock.ChannelId
	block.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Blockmodel.UpdateBlock(block, id, blocks.DB)

	if err != nil {

		return err

	}
	return nil

}

// Delete Collection

func (blocks *Block) DeleteCollection(id int, tenantid string) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var collection TblBlockCollection

	collection.BlockId = id
	collection.TenantId = tenantid
	collection.DeletedBy = blocks.UserId
	collection.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	collection.IsDeleted = 1

	err1 := Blockmodel.DeleteBlockCollection(collection, blocks.DB)

	if err1 != nil {
		return err1
	}
	return nil
}

// Delete tag in tbl_block_tags table

func (blocks *Block) DeleteTags(id int, name string, tenantid string) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var tag TblBlockTags

	tag.BlockId = id
	tag.TagName = name
	tag.TenantId = tenantid
	tag.DeletedBy = blocks.UserId
	tag.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tag.IsDeleted = 1

	err1 := Blockmodel.DeleteTag(tag, blocks.DB)

	if err1 != nil {
		return err1
	}
	return nil

}
