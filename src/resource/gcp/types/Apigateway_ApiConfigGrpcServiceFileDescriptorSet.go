package types

type Apigateway_ApiConfigGrpcServiceFileDescriptorSet struct {
	// Base64 encoded content of the file.
	Contents string `json:"contents,omitempty" yaml:"contents,omitempty"`

	// The file path (full or relative path). This is typically the path of the file when it is uploaded.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
