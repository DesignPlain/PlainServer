package types

type Compute_ServiceAttachmentConnectedEndpoint struct {
	/*
	   (Output)
	   The URL of the consumer forwarding rule.
	*/
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	/*
	   (Output)
	   The status of the connection from the consumer forwarding rule to
	   this service attachment.
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
