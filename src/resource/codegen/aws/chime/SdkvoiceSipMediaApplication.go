package chime

import types "libds/aws/types"

type SdkvoiceSipMediaApplication struct {
	// The AWS Region in which the AWS Chime SDK Voice Sip Media Application is created.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media application. Currently, only one endpoint is supported. See `endpoints`.
	Endpoints types.Chime_SdkvoiceSipMediaApplicationEndpoints `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	/*
	   The name of the AWS Chime SDK Voice Sip Media Application.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
