package types

type Dms_getEndpointElasticsearchSetting struct {
	//
	EndpointUri string `json:"endpointUri,omitempty" yaml:"endpointUri,omitempty"`

	//
	ErrorRetryDuration int `json:"errorRetryDuration,omitempty" yaml:"errorRetryDuration,omitempty"`

	//
	FullLoadErrorPercentage int `json:"fullLoadErrorPercentage,omitempty" yaml:"fullLoadErrorPercentage,omitempty"`

	//
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`
}
