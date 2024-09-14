package types

type Bedrock_AgentAgentActionGroupApiSchema struct {
	/*
	   JSON or YAML-formatted payload defining the OpenAPI schema for the action group.
	   Only one of `payload` or `s3` can be specified.
	*/
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`

	/*
	   Details about the S3 object containing the OpenAPI schema for the action group. See `s3` Block for details.
	   Only one of `s3` or `payload` can be specified.
	*/
	S3 Bedrock_AgentAgentActionGroupApiSchemaS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
