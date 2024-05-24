package networkmanager

type AttachmentAccepter struct {
	// The ID of the attachment.
	AttachmentId string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`

	// The type of attachment. Valid values can be found in the [AWS Documentation](https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_ListAttachments.html#API_ListAttachments_RequestSyntax)
	AttachmentType string `json:"attachmentType,omitempty" yaml:"attachmentType,omitempty"`
}
