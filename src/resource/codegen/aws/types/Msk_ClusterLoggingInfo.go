package types

type Msk_ClusterLoggingInfo struct {
	// Configuration block for Broker Logs settings for logging info. See below.
	BrokerLogs Msk_ClusterLoggingInfoBrokerLogs `json:"brokerLogs,omitempty" yaml:"brokerLogs,omitempty"`
}
