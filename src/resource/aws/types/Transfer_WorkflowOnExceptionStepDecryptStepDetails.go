package types

type Transfer_WorkflowOnExceptionStepDecryptStepDetails struct {
	// Specifies the location for the file being copied. Use ${Transfer:username} in this field to parametrize the destination prefix by username.
	DestinationFileLocation Transfer_WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocation `json:"destinationFileLocation,omitempty" yaml:"destinationFileLocation,omitempty"`

	// The name of the step, used as an identifier.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A flag that indicates whether or not to overwrite an existing file of the same name. The default is `FALSE`. Valid values are `TRUE` and `FALSE`.
	OverwriteExisting string `json:"overwriteExisting,omitempty" yaml:"overwriteExisting,omitempty"`

	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow. Enter ${previous.file} to use the previous file as the input. In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value. Enter ${original.file} to use the originally-uploaded file location as input for this step.
	SourceFileLocation string `json:"sourceFileLocation,omitempty" yaml:"sourceFileLocation,omitempty"`

	// The type of encryption used. Currently, this value must be `"PGP"`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
