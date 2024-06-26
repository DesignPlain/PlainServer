package endpoints

type Service struct {
	/*
	   The full text of the OpenAPI YAML configuration as described [here](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/2.0.md).
	   Either this, or -both- of `grpc_config` and `protoc_output_base64` must be specified.
	*/
	OpenapiConfig string `json:"openapiConfig,omitempty" yaml:"openapiConfig,omitempty"`

	// The project ID that the service belongs to.  If not provided, provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The full contents of the Service Descriptor File generated by protoc.  This should be a compiled .pb file, base64-encoded.
	ProtocOutputBase64 string `json:"protocOutputBase64,omitempty" yaml:"protocOutputBase64,omitempty"`

	/*
	   The name of the service.  Usually of the form `$apiname.endpoints.$projectid.cloud.goog`.

	   - - -
	*/
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	/*
	   The full text of the Service Config YAML file (Example located [here](https://github.com/GoogleCloudPlatform/python-docs-samples/blob/main/endpoints/bookstore-grpc/api_config.yaml)).
	   If provided, must also provide `protoc_output_base64`.  `open_api` config must -not- be provided.
	*/
	GrpcConfig string `json:"grpcConfig,omitempty" yaml:"grpcConfig,omitempty"`
}
