package ssm

import types "libds/aws/types"

type Document struct {
	// Additional permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// A map of tags to assign to the object. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The target type which defines the kinds of resources the document can run on. For example, `/AWS::EC2::Instance`. For a list of valid resource types, see [AWS resource and property types reference](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	TargetType string `json:"targetType,omitempty" yaml:"targetType,omitempty"`

	// The version of the artifact associated with the document. For example, `12.6`. This value is unique across all versions of a document, and can't be changed.
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`

	// One or more configuration blocks describing attachments sources to a version of a document. See `attachments_source` block below for details.
	AttachmentsSources []types.Ssm_DocumentAttachmentsSource `json:"attachmentsSources,omitempty" yaml:"attachmentsSources,omitempty"`

	// The content for the SSM document in JSON or YAML format. The content of the document must not exceed 64KB. This quota also includes the content specified for input parameters at runtime. We recommend storing the contents for your new document in an external JSON or YAML file and referencing the file in a command.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The format of the document. Valid values: `JSON`, `TEXT`, `YAML`.
	DocumentFormat string `json:"documentFormat,omitempty" yaml:"documentFormat,omitempty"`

	// The type of the document. For a list of valid values, see the [API Reference](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreateDocument.html#systemsmanager-CreateDocument-request-DocumentType).
	DocumentType string `json:"documentType,omitempty" yaml:"documentType,omitempty"`
}
