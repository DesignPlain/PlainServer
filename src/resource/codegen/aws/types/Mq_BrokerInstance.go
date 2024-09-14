package types

type Mq_BrokerInstance struct {
	// The URL of the [ActiveMQ Web Console](http://activemq.apache.org/web-console.html) or the [RabbitMQ Management UI](https://www.rabbitmq.com/management.html#external-monitoring) depending on `engine_type`.
	ConsoleUrl string `json:"consoleUrl,omitempty" yaml:"consoleUrl,omitempty"`

	/*
	   Broker's wire-level protocol endpoints in the following order & format referenceable e.g., as `instances.0.endpoints.0` (SSL):
	   - For `ActiveMQ`:
	   - `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
	   - `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
	   - `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
	   - `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
	   - `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
	   - For `RabbitMQ`:
	   - `amqps://broker-id.mq.us-west-2.amazonaws.com:5671`
	*/
	Endpoints []string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	// IP Address of the broker.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
