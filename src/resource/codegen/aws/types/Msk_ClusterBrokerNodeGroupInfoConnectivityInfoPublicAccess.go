package types

type Msk_ClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess struct {
	// Public access type. Valid values: `DISABLED`, `SERVICE_PROVIDED_EIPS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
