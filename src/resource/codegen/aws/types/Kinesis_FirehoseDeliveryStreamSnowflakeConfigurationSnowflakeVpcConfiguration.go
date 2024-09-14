package types

type Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration struct {
	// The VPCE ID for Firehose to privately connect with Snowflake.
	PrivateLinkVpceId string `json:"privateLinkVpceId,omitempty" yaml:"privateLinkVpceId,omitempty"`
}
