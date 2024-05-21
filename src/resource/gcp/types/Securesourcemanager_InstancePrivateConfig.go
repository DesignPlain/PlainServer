package types

type Securesourcemanager_InstancePrivateConfig struct {
	/*
	   (Output)
	   Service Attachment for SSH, resource is in the format of `projects/{project}/regions/{region}/serviceAttachments/{service_attachment}`.
	*/
	SshServiceAttachment string `json:"sshServiceAttachment,omitempty" yaml:"sshServiceAttachment,omitempty"`

	// CA pool resource, resource must in the format of `projects/{project}/locations/{location}/caPools/{ca_pool}`.
	CaPool string `json:"caPool,omitempty" yaml:"caPool,omitempty"`

	/*
	   (Output)
	   Service Attachment for HTTP, resource is in the format of `projects/{project}/regions/{region}/serviceAttachments/{service_attachment}`.
	*/
	HttpServiceAttachment string `json:"httpServiceAttachment,omitempty" yaml:"httpServiceAttachment,omitempty"`

	// 'Indicate if it's private instance.'
	IsPrivate bool `json:"isPrivate,omitempty" yaml:"isPrivate,omitempty"`
}
