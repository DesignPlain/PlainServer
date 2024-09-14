package pubsub

type LiteReservation struct {
	/*
	   Name of the reservation.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the pubsub lite reservation.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The reserved throughput capacity. Every unit of throughput capacity is
	   equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
	   messages.
	*/
	ThroughputCapacity int `json:"throughputCapacity,omitempty" yaml:"throughputCapacity,omitempty"`
}
