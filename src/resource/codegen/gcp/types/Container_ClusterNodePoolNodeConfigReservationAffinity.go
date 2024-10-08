package types

type Container_ClusterNodePoolNodeConfigReservationAffinity struct {
	// The label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify "compute.googleapis.com/reservation-name" as the key and specify the name of your reservation as its value.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The list of label values of reservation resources. For example: the name of the specific reservation when using a key of "compute.googleapis.com/reservation-name"
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	/*
	   The type of reservation consumption
	   Accepted values are:

	   - `"UNSPECIFIED"`: Default value. This should not be used.
	   - `"NO_RESERVATION"`: Do not consume from any reserved capacity.
	   - `"ANY_RESERVATION"`: Consume any reservation available.
	   - `"SPECIFIC_RESERVATION"`: Must consume from a specific reservation. Must specify key value fields for specifying the reservations.
	*/
	ConsumeReservationType string `json:"consumeReservationType,omitempty" yaml:"consumeReservationType,omitempty"`
}
