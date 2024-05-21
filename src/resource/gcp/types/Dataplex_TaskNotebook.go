package types

type Dataplex_TaskNotebook struct {
	// Cloud Storage URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// Cloud Storage URIs of files to be placed in the working directory of each executor.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	/*
	   Infrastructure specification for the execution.
	   Structure is documented below.
	*/
	InfrastructureSpec Dataplex_TaskNotebookInfrastructureSpec `json:"infrastructureSpec,omitempty" yaml:"infrastructureSpec,omitempty"`

	// Path to input notebook. This can be the Cloud Storage URI of the notebook file or the path to a Notebook Content. The execution args are accessible as environment variables (TASK_key=value).
	Notebook string `json:"notebook,omitempty" yaml:"notebook,omitempty"`
}
