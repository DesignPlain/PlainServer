package types

type Transfer_WorkflowOnExceptionStep struct {
	// Details for a step that performs a file copy. See Copy Step Details below.
	CopyStepDetails Transfer_WorkflowOnExceptionStepCopyStepDetails `json:"copyStepDetails,omitempty" yaml:"copyStepDetails,omitempty"`

	// Details for a step that invokes a lambda function.
	CustomStepDetails Transfer_WorkflowOnExceptionStepCustomStepDetails `json:"customStepDetails,omitempty" yaml:"customStepDetails,omitempty"`

	// Details for a step that decrypts the file.
	DecryptStepDetails Transfer_WorkflowOnExceptionStepDecryptStepDetails `json:"decryptStepDetails,omitempty" yaml:"decryptStepDetails,omitempty"`

	// Details for a step that deletes the file.
	DeleteStepDetails Transfer_WorkflowOnExceptionStepDeleteStepDetails `json:"deleteStepDetails,omitempty" yaml:"deleteStepDetails,omitempty"`

	// Details for a step that creates one or more tags.
	TagStepDetails Transfer_WorkflowOnExceptionStepTagStepDetails `json:"tagStepDetails,omitempty" yaml:"tagStepDetails,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
