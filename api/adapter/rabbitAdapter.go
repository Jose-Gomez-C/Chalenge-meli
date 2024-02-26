package adapter

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/Jose-Gomez-c/challenge/api/model"
// 	"github.com/streadway/amqp"
// )

// type RabbitAdpater interface {
// 	SendLine(product model.ProductId)
// }

// type RabbitAdpaterLayer struct {
// 	conn    *amqp.Connection
// 	channel *amqp.Channel
// 	queue   amqp.Queue
// }

// func NewRabbitAdpater(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue) RabbitAdpater {
// 	return &RabbitAdpaterLayer{conn: conn, channel: channel, queue: queue}
// }

// func (layer RabbitAdpaterLayer) GetChanel() (*amqp.Channel, error) {
// 	channel, err := layer.conn.Channel()
// 	if err != nil {
// 		fmt.Println("No se pudo conectar canal", err)
// 		return nil, err
// 	}
// 	defer channel.Close()
// 	return channel, nil

// }
// func (layer RabbitAdpaterLayer) GetQueue(channel *amqp.Channel) (amqp.Queue, error) {
// 	queue, err := channel.QueueDeclare(
// 		"fileQueue",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		fmt.Println("No se pudo declarar la cola ", err)
// 		return queue, err
// 	}
// 	return queue, nil
// }

// func (layer RabbitAdpaterLayer) SendLine(product model.ProductId) {
// 	body, err := json.Marshal(product)
// 	if err != nil {
// 		fmt.Println("Error al convertir el producto ", err)
// 		return
// 	}
// 	err = layer.channel.Publish(
// 		"",
// 		layer.queue.Name,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		})
// 	if err != nil {
// 		fmt.Println("Error al publicar el mensaje ", err)
// 		return
// 	}
// }

// func (layer RabbitAdpaterLayer) getMessage() {
// 	var product model.ProductId
// 	msgs, err := layer.channel.Consume(
// 		layer.queue.Name,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		log.Fatalf("Error al consumir mensajes: %s", err)
// 	}
// 	go func() {

// 		for msg := range msgs {
// 			err := json.Unmarshal(msg.Body, &product)
// 			if err != nil {
// 				fmt.Println("Error al convertir el mensaje", err)
// 				return
// 			}

// 		}
// 	}()
// }
