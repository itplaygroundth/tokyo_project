package consumer

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// ฟังก์ชันที่ใช้เริ่ม Consumer และดึงข้อมูลจาก Queue
func StartConsumer(consumerID int) {
	conn, err := amqp.Dial("amqp://guest:guest@128.199.92.45:5672/")
	if err != nil {
		log.Fatalf("เชื่อมต่อ RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("เปิด channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"transactions_queue", // ชื่อ Queue
		true,                 // Durable
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No-wait
		nil,                  // Arguments
	)
	if err != nil {
		log.Fatalf("ติดต่อ queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // ชื่อ Queue
		"",     // Consumer
		true,   // Auto Acknowledge
		false,  // Exclusive
		false,  // No Local
		false,  // No Wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("ลงทะเบียน consumer: %s", err)
	}

	// รับและประมวลผลข้อความในแบบขนาน
	for msg := range msgs {
		go processMessage(consumerID, msg)
	}
}

// ฟังก์ชันสำหรับประมวลผลข้อความ
func processMessage(consumerID int, msg amqp.Delivery) {
	fmt.Printf("Consumer %d รับข้อมูล: %s\n", consumerID, msg.Body)
	// ทำการประมวลผลข้อมูล เช่น บันทึกข้อมูลลง Redis
	time.Sleep(1 * time.Second) // จำลองการประมวลผล
	fmt.Printf("Consumer %d ประมวลผลสำเร็จ: %s\n", consumerID, msg.Body)
}