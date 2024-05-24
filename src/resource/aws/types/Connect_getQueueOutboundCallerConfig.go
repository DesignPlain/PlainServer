package types

type Connect_getQueueOutboundCallerConfig struct {
	// Outbound whisper flow to be used during an outbound call.
	OutboundFlowId string `json:"outboundFlowId,omitempty" yaml:"outboundFlowId,omitempty"`

	// Specifies the caller ID name.
	OutboundCallerIdName string `json:"outboundCallerIdName,omitempty" yaml:"outboundCallerIdName,omitempty"`

	// Specifies the caller ID number.
	OutboundCallerIdNumberId string `json:"outboundCallerIdNumberId,omitempty" yaml:"outboundCallerIdNumberId,omitempty"`
}
