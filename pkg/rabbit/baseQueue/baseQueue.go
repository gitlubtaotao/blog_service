package baseQueue

import (
	"blog_service/pkg/rabbit"
	"github.com/streadway/amqp"
)

type BaseQueue struct {
	QueueName string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Arguments  map[string]interface{}
}
// send 
func (b *BaseQueue) Send() error   {
	return nil
}

func (b *BaseQueue) Receive(callback func(content interface{}),) error  {
	conn, err := amqp.Dial(rabbit.RabbitUrl)
	defer conn.Close()
	if err != nil{
		return err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil{
		return err
	}
	q, err := ch.QueueDeclare(
		b.QueueName, // name
		b.Durable,   // durable
		b.AutoDelete,   // delete when unused
		b.Exclusive,   // exclusive
		b.NoWait,   // no-wait
		b.Arguments,     // arguments
	)
	if err != nil{
		return err
	}
	mags, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil{
		return err
	}
	forever := make(chan bool)
	go func() {
		for d := range mags {
			callback(d.Body)
		}
	}()
	<-forever
	return nil
}
