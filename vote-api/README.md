# Pollination - Vote API

REST API that receives votes from users and sends them to a queue to be processed later.

## Execution

### Running the API

You can run the API with the following command:

```shell
make run
```

## Dependencies

- Go 1.18 (https://go.dev)
- Logrus: structured logger (https://github.com/sirupsen/logrus)
- Fiber v2: web framework (https://gofiber.io)
- Go AMQP 0.9.1: RabbitMQ client library (https://github.com/rabbitmq/amqp091-go)
