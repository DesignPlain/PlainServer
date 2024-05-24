package auditmanager

type FrameworkShare struct {
	/*
	   Unique identifier for the shared custom framework.

	   The following arguments are optional:
	*/
	FrameworkId string `json:"frameworkId,omitempty" yaml:"frameworkId,omitempty"`

	// Comment from the sender about the share request.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Amazon Web Services account of the recipient.
	DestinationAccount string `json:"destinationAccount,omitempty" yaml:"destinationAccount,omitempty"`

	// Amazon Web Services region of the recipient.
	DestinationRegion string `json:"destinationRegion,omitempty" yaml:"destinationRegion,omitempty"`
}
