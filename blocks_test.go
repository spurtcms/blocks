package blocks

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "postgres",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "Nov16",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to database:", err)

	}
	if err != nil {

		return nil, err

	}

	return db, nil
}

func TestBlockList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    2,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Blocks", auth.CRUD, 1)

	Blocks := BlockSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permission {
		Blocks.UserId = 2
		blocklist, blockcount, err := Blocks.BlockList(10, 0, Filter{}, 1)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("blocklist:", blocklist, "blockcount:", blockcount)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestCreate(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    2,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Blocks", auth.CRUD, 1)

	Blocks := BlockSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permission {

		BlockCreate := BlockCreation{
			Title:        "Block",
			BlockContent: `<div class="card-6 p-[12px] bg-white rounded flex gap-[8px]"><div class="flex items-start gap-[6px]"><div class="flex flex-col gap-[4px]"><div class="w-[24px] h-[24px]"><img src="https://dev.spurtcms.com/public/img/block-6.png" alt="" class="h-full w-full object-cover"></div><div class="w-[24px] h-[24px] object-contain"><img src="https://dev.spurtcms.com/public/img/block-6.png" alt="" class="h-full w-full object-cover"></div></div><div class="w-[80px] h-full"><img src="https://dev.spurtcms.com/public/img/block-6.png" alt="" class="w-full rounded-[2px] h-full"></div></div><div class="flex flex-col gap-[8px] items-start w-full"><div class="flex flex-col items-start gap-[4px]"><h5 class="text-xs font-normal mb-0 text-[#262626]">Lorem ipsum dolor sit</h5><div class="flex gap-[1px] items-center"><img src="img/reviewstar-full.svg" alt=""><img src="https://dev.spurtcms.com/public/img/reviewstar-full.svg" alt=""><img src="https://dev.spurtcms.com/public/img/reviewstar-full.svg" alt=""><img src="https://dev.spurtcms.com/public/img/reviewstar-full.svg" alt=""><img src="https://dev.spurtcms.com/public/img/reviewstar.svg" alt=""></div></div><p class="m-0 text-[#262626] text-xs font-normal">$ 15.00</p><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Color</h5><div class="flex items-center gap-[1px]"><div class="w-[20px] h-[20px] rounded-full p-[2px] border border-[#000000]"><div class="bg-[#000000] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#004DFF] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#3B8620] h-full rounded-full"></div></div></div></div><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Size</h5><div class="flex items-center gap-[8px]"><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">s</div><div class="h-[16px] flex items-center justify-center px-[6px] border bg-[#EBEBEB] border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">m</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">l</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">xl</div></div></div></div></div>`,
			CoverImage:   "/blocks/IMG-1726551883.jpeg",
			TenantId:     1,
			Prime:        1,
			CreatedBy:    2,
			IsActive:     1,
		}

		createblock, err := Blocks.CreateBlock(BlockCreate,"",0)

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
		fmt.Println(createblock)
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestBlockUpdate(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    2,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Blocks", auth.CRUD, 1)

	Blocks := BlockSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permission {

		BlockCreate := BlockCreation{
			Title:      "Halo",
			ModifiedBy: 2,
			TenantId:   1,
		}

		err := Blocks.UpdateBlock(23, BlockCreate)

		if err != nil {
			log.Println("collection list", err)
		}
		fmt.Println("hello")
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestBlockDelete(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:    2,
		ExpiryFlg: false,
		SecretKey: "Secret123",
		DB:        db,
		RoleId:    2,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permission, _ := Auth.IsGranted("Blocks", auth.CRUD, 1)

	Blocks := BlockSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permission {

		err := Blocks.RemoveBlock(23, 1)
		if err != nil {
			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")

	}
}
