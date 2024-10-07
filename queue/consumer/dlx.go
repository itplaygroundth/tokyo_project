package consumer

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

func Dlx() {
    // Connect to RabbitMQ
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    defer ch.Close()

    // Declare the Dead Letter Exchange (DLX)
    err = ch.ExchangeDeclare(
        "dlx_exchange", // name
        "direct",       // type
        true,           // durable
        false,          // auto-deleted
        false,          // internal
        false,          // no-wait
        nil,            // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare the Dead Letter Exchange: %s", err)
    }

    // Declare the Dead Letter Queue (DLQ)
    _, err = ch.QueueDeclare(
        "dead_letter_queue", // name of the queue
        true,                // durable
        false,               // delete when unused
        false,               // exclusive
        false,               // no-wait
        nil,                 // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare the Dead Letter Queue: %s", err)
    }

    // Bind the DLQ to the DLX
    err = ch.QueueBind(
        "dead_letter_queue", // name of the queue
        "failed_message",    // routing key
        "dlx_exchange",      // exchange
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to bind the DLQ to DLX: %s", err)
    }

    // Declare the Main Queue with DLX argument
    args := amqp.Table{
        "x-dead-letter-exchange": "dlx_exchange", // specify the DLX
        "x-dead-letter-routing-key": "failed_message",
    }

    _, err = ch.QueueDeclare(
        "main_queue", // name of the main queue
        true,         // durable
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        args,         // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare the Main Queue: %s", err)
    }

    fmt.Println("Main Queue and Dead Letter Queue are set up successfully.")
}
