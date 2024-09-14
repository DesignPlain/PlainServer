package sesv2

type DedicatedIpAssignment struct {
	// Dedicated IP address.
	DestinationPoolName string `json:"destinationPoolName,omitempty" yaml:"destinationPoolName,omitempty"`

	// Dedicated IP address.
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`
}
