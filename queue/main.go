package main

import (
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/compress"
	//"github.com/gin-gonic/gin"
	"log"
	"sync"
	cons "london/consumer"
	//"os"
	//"fmt"
	//"hanoi/rabbitmq"
	//"hanoi/handler"
	//"hanoi/users"
	//"hanoi/route"
	//"hanoi/database"
	//"hanoi/models"
	//"hanoi/handler/njwt"
	// "gorm.io/gorm"
	// "time"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/fiber-swagger"
	//"github.com/swaggo/files"
	// _ "hanoi/docs" // สำหรับเอกสาร Swagger

	
)

// func loadDatabase() {
// 	if err := database.Connect(); err != nil {
// 		handleError(err)
// 	}

// }

// func DropTable () {

// 	database.Database.Migrator().DropTable(&models.TransactionSub{})
// 	database.Database.Migrator().DropTable(&models.BuyInOut{})

// }

// func migrateNormal(db *gorm.DB) {

// 	if err := db.AutoMigrate(&models.Product{},&models.BanksAccount{},&models.Users{},&models.TransactionSub{},&models.BankStatement{},&models.BuyInOut{}); err != nil {
// 		handleError(err)
// 	}
	 
// 	fmt.Println("Migrations Normal Tables executed successfully")
// }
// func migrateAdmin() {
 
// 	if err := database.Database.AutoMigrate(&models.TsxAdmin{},&models.Provider{}); err != nil {
// 		handleError(err)
// 	}
// 	fmt.Println("Migrations Admin Tables executed successfully")
// }



// @title Api Goteway in Go
// @version 1.0
// @description Api Goteway in Go.
// @host 167.71.100.123:3003
// @BasePath /api/v1


func handleError(err error) {
	log.Fatal(err)
}

//func main() {

	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
    //     AllowOrigins: "*", // อนุญาตทุกโดเมน (ในโปรดักชันให้ระบุโดเมนที่จำเป็นเท่านั้น)
    //     AllowMethods: "GET,POST,PUT,DELETE",
    //     AllowHeaders: "Origin, Content-Type, Accept",
    // }))
	// app.Use(compress.New())

 

	// app.Use(func(c *fiber.Ctx) error {
	// 	loc, _ := time.LoadLocation("Asia/Bangkok")
	// 	c.Locals("location", loc)
	// 	return c.Next()
	// })

	// app.Use(logger.New())
	
	
	//  //migrateAdmin()

	//  v1 := app.Group("/api/v1")
	//  route.SetupRoutes(v1)
 
    // // เรียกใช้ฟังก์ชันจาก efinity.go
	// log.Fatal(app.Listen(":8030"))
	 
	const (
		consumerCount = 5 // จำนวน Consumer ที่ต้องการ
	)
	
	func main() {
		var wg sync.WaitGroup
	
		// เริ่ม Consumer หลายตัวตามจำนวนที่กำหนด
		for i := 0; i < consumerCount; i++ {
			wg.Add(1)
			go func(consumerID int) {
				defer wg.Done()
				cons.StartConsumer(consumerID)
			}(i)
		}
	
		// รอให้ทุก Consumer ทำงานเสร็จ
		wg.Wait()
	}
//}