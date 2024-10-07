package logger

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// สร้าง Redis Client
var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

// ฟังก์ชันสำหรับบันทึกข้อผิดพลาด
func logError(transactionID string, err error) {
	log.Printf("Error processing transaction %s: %v", transactionID, err)
}

// ฟังก์ชันตรวจสอบและประมวลผลธุรกรรม
func processTransaction(transactionID string, transactionData map[string]interface{}) {
	ctx := context.Background()
	idempotencyKey := "processed_transaction:" + transactionID

	// ตรวจสอบว่า Transaction ID นี้ได้ถูกประมวลผลไปแล้วหรือไม่
	exists, err := rdb.Exists(ctx, idempotencyKey).Result()
	if err != nil {
		logError(transactionID, err)
		return
	}

	if exists > 0 {
		log.Printf("Transaction %s has already been processed, skipping.", transactionID)
		return // ถ้าข้อมูลซ้ำ ไม่ต้องทำอะไรเพิ่มเติม
	}

	// ประมวลผลธุรกรรม
	log.Printf("Processing transaction %s...", transactionID)
	time.Sleep(1 * time.Second) // จำลองการประมวลผล

	// บันทึก Transaction ID ลง Redis
	err = rdb.Set(ctx, idempotencyKey, "processed", 0).Err()
	if err != nil {
		logError(transactionID, err)
		return
	}

	log.Printf("Transaction %s processed and saved in Redis.", transactionID)
}

// ฟังก์ชันตัวอย่างการประมวลผลธุรกรรม
func Logger(done chan struct{}) {
	transactionID := "273365"
	transactionData := map[string]interface{}{
		"accountno":         "3690815150",
		"bankname":          "KTB",
		"transactionamount": "0.00",
		"beforebalance":     "86.20",
		"balance":           "86.20",
		"status":            "101",
	}

	// ประมวลผลธุรกรรม
	processTransaction(transactionID, transactionData)

	// ทดลองประมวลผลธุรกรรมอีกครั้งเพื่อทดสอบการป้องกันข้อมูลซ้ำ
	processTransaction(transactionID, transactionData)
}

// ฟังก์ชันสำหรับเก็บ log ลง Redis โดยมีระยะเวลาหมดอายุ 72 ชั่วโมง
func logToRedis(logID string, logData map[string]interface{}) error {
	ctx := context.Background()

	// แปลงข้อมูล log ให้เป็น JSON string สำหรับบันทึกลงใน Redis
	logString, err := json.Marshal(logData)
	if err != nil {
		return err
	}

	// กำหนดระยะเวลาหมดอายุเป็น 72 ชั่วโมง
	expiration := 72 * time.Hour

	// บันทึกข้อมูลลง Redis พร้อมตั้งค่า TTL (Expiration)
	err = rdb.Set(ctx, logID, logString, expiration).Err()
	if err != nil {
		return err
	}

	log.Printf("Log %s has been saved to Redis with a TTL of 72 hours.", logID)
	return nil
}