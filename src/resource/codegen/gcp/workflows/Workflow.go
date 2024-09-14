package workflows

type Workflow struct {
	/*
	   Name of the service account associated with the latest workflow version. This service
	   account represents the identity of the workflow and determines what permissions the workflow has.
	   Format: projects/{project}/serviceAccounts/{account} or {account}.
	   Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	   The {account} value can be the email address or the unique_id of the service account.
	   If not provided, workflow will use the project's default service account.
	   Modifying this field for an existing workflow results in a new workflow revision.
	*/
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Workflow code to be executed. The size limit is 128KB.
	SourceContents string `json:"sourceContents,omitempty" yaml:"sourceContents,omitempty"`

	// User-defined environment variables associated with this workflow revision. This map has a maximum length of 20. Each string can take up to 4KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS".
	UserEnvVars map[string]string `json:"userEnvVars,omitempty" yaml:"userEnvVars,omitempty"`

	/*
	   A set of key/value label pairs to assign to this Workflow.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Name of the Workflow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the
	   specified prefix. If this and name are unspecified, a random value is chosen for the name.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the workflow.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Describes the level of platform logging to apply to calls and call responses during
	   executions of this workflow. If both the workflow and the execution specify a logging level,
	   the execution level takes precedence.
	   Possible values are: `CALL_LOG_LEVEL_UNSPECIFIED`, `LOG_ALL_CALLS`, `LOG_ERRORS_ONLY`, `LOG_NONE`.
	*/
	CallLogLevel string `json:"callLogLevel,omitempty" yaml:"callLogLevel,omitempty"`

	/*
	   The KMS key used to encrypt workflow and execution data.
	   Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	*/
	CryptoKeyName string `json:"cryptoKeyName,omitempty" yaml:"cryptoKeyName,omitempty"`

	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
