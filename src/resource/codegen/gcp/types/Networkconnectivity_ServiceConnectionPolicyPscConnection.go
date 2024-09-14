package types

type Networkconnectivity_ServiceConnectionPolicyPscConnection struct {
	// The resource reference of the consumer address.
	ConsumerAddress string `json:"consumerAddress,omitempty" yaml:"consumerAddress,omitempty"`

	/*
	   The most recent error during operating this connection.
	   Structure is documented below.
	*/
	Error Networkconnectivity_ServiceConnectionPolicyPscConnectionError `json:"error,omitempty" yaml:"error,omitempty"`

	/*
	   The error type indicates whether the error is consumer facing, producer
	   facing or system internal.
	   Possible values are: `CONNECTION_ERROR_TYPE_UNSPECIFIED`, `ERROR_INTERNAL`, `ERROR_CONSUMER_SIDE`, `ERROR_PRODUCER_SIDE`.
	*/
	ErrorType string `json:"errorType,omitempty" yaml:"errorType,omitempty"`

	/*
	   The state of the PSC connection.
	   Possible values are: `STATE_UNSPECIFIED`, `ACTIVE`, `CREATING`, `DELETING`, `FAILED`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The resource reference of the PSC Forwarding Rule within the consumer VPC.
	ConsumerForwardingRule string `json:"consumerForwardingRule,omitempty" yaml:"consumerForwardingRule,omitempty"`

	// The project where the PSC connection is created.
	ConsumerTargetProject string `json:"consumerTargetProject,omitempty" yaml:"consumerTargetProject,omitempty"`

	/*
	   The error info for the latest error during operating this connection.
	   Structure is documented below.
	*/
	ErrorInfo Networkconnectivity_ServiceConnectionPolicyPscConnectionErrorInfo `json:"errorInfo,omitempty" yaml:"errorInfo,omitempty"`

	// The last Compute Engine operation to setup PSC connection.
	GceOperation string `json:"gceOperation,omitempty" yaml:"gceOperation,omitempty"`

	// The PSC connection id of the PSC forwarding rule.
	PscConnectionId string `json:"pscConnectionId,omitempty" yaml:"pscConnectionId,omitempty"`
}
