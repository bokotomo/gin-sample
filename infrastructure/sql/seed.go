package main

import (
	"gin-sample/driver"
	. "gin-sample/model"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(".env file cannot be load.")
	}
	driver.Init()
	driver.DB.Unscoped().Delete(&Design{}, &Token{}, &User{})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 1}, Title: "design1", Text: "Hello1", Good: 1, Comments: 2, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 2}, Title: "design2", Text: "Hello2", Good: 2, Comments: 8, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 3}, Title: "design3", Text: "Hello3", Good: 3, Comments: 12, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 4}, Title: "design4", Text: "Hello4", Good: 4, Comments: 12, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 5}, Title: "design5", Text: "Hello5", Good: 5, Comments: 5, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 6}, Title: "design6", Text: "Hello6", Good: 6, Comments: 42, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 7}, Title: "design7", Text: "Hello7", Good: 7, Comments: 2, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 8}, Title: "design8", Text: "Hello8", Good: 8, Comments: 1, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 9}, Title: "design9", Text: "Hello7", Good: 7, Comments: 2, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 10}, Title: "design10", Text: "Hello8", Good: 18, Comments: 11, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 11}, Title: "design11", Text: "Hello7", Good: 71, Comments: 2, Thumbnail: "image"})
	driver.DB.Create(&Design{Model: gorm.Model{ID: 12}, Title: "design12", Text: "Hello8", Good: 18, Comments: 1, Thumbnail: "image"})
	defer driver.DB.Close()
}
