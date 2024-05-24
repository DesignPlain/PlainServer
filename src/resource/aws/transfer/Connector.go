package transfer

import types "DesignSphere_Server/src/resource/aws/types"

type Connector struct {
	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	As2Config types.Transfer_ConnectorAs2Config `json:"as2Config,omitempty" yaml:"as2Config,omitempty"`

	// The IAM Role which is required for allowing the connector to turn on CloudWatch logging for Amazon S3 events.
	LoggingRole string `json:"loggingRole,omitempty" yaml:"loggingRole,omitempty"`

	// Either SFTP or AS2 is configured.The parameters to configure for the connector object. Fields documented below.
	SftpConfig types.Transfer_ConnectorSftpConfig `json:"sftpConfig,omitempty" yaml:"sftpConfig,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The URL of the partners AS2 endpoint or SFTP endpoint.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The IAM Role which provides read and write access to the parent directory of the file location mentioned in the StartFileTransfer request.
	AccessRole string `json:"accessRole,omitempty" yaml:"accessRole,omitempty"`
}
