# Blocks

# forms-builders

# Installation

``` bash
go get github.com/spurtcms/blocks 
```


# Usage Example

``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB: &gorm.DB{},
		RoleId: 1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Forms Builder", auth.CRUD, 1)

	Blocks := BlockSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})


	if permisison {

		//list Blocks
		blocklist, blockcount, err := Blocks.BlockList(10, 0, Filter{}, 1)
		if err != nil {
			fmt.Println(err)
		}

		//create Blocks
        BlockCreate := BlockCreation{
			Title:        "TestBlock",
			BlockContent: `<p>Hello world</p>`,
			CoverImage:   "/blocks/IMG-1726551883.jpeg",
			TenantId:     1,
			Prime:        1,
			CreatedBy:    2,
			IsActive:     1,
		}

		createblock, err := Blocks.CreateBlock(BlockCreate)

		if err != nil {
			log.Println("collection list", err)
		}

		MstrTag := MasterTagCreate{

			TagTitle:  "Hello",
			TenantId:  1,
			CreatedBy: 2,
		}

		createtags, err := Blocks.CreateMasterTag(MstrTag)

		if err != nil {
			fmt.Println("block err", err)
		}

		TagCreate := CreateTag{
			BlockId:   createblock.Id,
			TagId:     createtags.Id,
			TagName:   createtags.TagTitle,
			TenantId:  1,
			CreatedBy: 2,
		}

		err2 := Blocks.CreateTag(TagCreate)

		if err2 != nil {
			fmt.Println("block err", err)
		}

		//update Blocks
        BlockUpdate := BlockCreation{
			Title:      "Halo",
			ModifiedBy: 2,
			TenantId:   1,
		}

		err := Blocks.UpdateBlock(23, BlockUpdate)

		if err != nil {
			log.Println("collection list", err)
		}

		// delete Blocks
		err := Blocks.RemoveBlock(23, 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

```
# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/blocks/issues]. 
