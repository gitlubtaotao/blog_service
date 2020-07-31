package rabbit

type IRabbit interface{
	Send()
	Receive()
}

const RabbitUrl = "amqp://guest:guest@localhost:5672/"