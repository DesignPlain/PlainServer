package ssm

import types "DesignSphere_Server/src/resource/aws/types"

type Document struct {
	// The format of the document. Valid document types include: `JSON` and `YAML`
	DocumentFormat string `json:"documentFormat,omitempty" yaml:"documentFormat,omitempty"`

	// The type of the document. Valid document types include: `Automation`, `Command`, `Package`, `Policy`, and `Session`
	DocumentType string `json:"documentType,omitempty" yaml:"documentType,omitempty"`

	// A field specifying the version of the artifact you are creating with the document. For example, "Release 12, Update 6". This value is unique across all versions of a document and cannot be changed for an existing document version.
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`

	// Additional Permissions to attach to the document. See Permissions below for details.
	Permissions map[string]string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// A map of tags to assign to the object. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance. For a list of valid resource types, see AWS Resource Types Reference (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	TargetType string `json:"targetType,omitempty" yaml:"targetType,omitempty"`

	// One or more configuration blocks describing attachments sources to a version of a document. Defined below.
	AttachmentsSources []types.Ssm_DocumentAttachmentsSource `json:"attachmentsSources,omitempty" yaml:"attachmentsSources,omitempty"`

	// The JSON or YAML content of the document.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The name of the document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
