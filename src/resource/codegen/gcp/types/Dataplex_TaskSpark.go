package types

type Dataplex_TaskSpark struct {
	// The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in jar_file_uris. The execution args are passed in as a sequence of named process arguments (--key=value).
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	// The Cloud Storage URI of the jar file that contains the main class. The execution args are passed in as a sequence of named process arguments (--key=value).
	MainJarFileUri string `json:"mainJarFileUri,omitempty" yaml:"mainJarFileUri,omitempty"`

	// The Gcloud Storage URI of the main Python file to use as the driver. Must be a .py file. The execution args are passed in as a sequence of named process arguments (--key=value).
	PythonScriptFile string `json:"pythonScriptFile,omitempty" yaml:"pythonScriptFile,omitempty"`

	// The query text. The execution args are used to declare a set of script variables (set key='value';).
	SqlScript string `json:"sqlScript,omitempty" yaml:"sqlScript,omitempty"`

	// A reference to a query file. This can be the Cloud Storage URI of the query file or it can the path to a SqlScript Content. The execution args are used to declare a set of script variables (set key='value';).
	SqlScriptFile string `json:"sqlScriptFile,omitempty" yaml:"sqlScriptFile,omitempty"`

	// Cloud Storage URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// Cloud Storage URIs of files to be placed in the working directory of each executor.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	/*
	   Infrastructure specification for the execution.
	   Structure is documented below.
	*/
	InfrastructureSpec Dataplex_TaskSparkInfrastructureSpec `json:"infrastructureSpec,omitempty" yaml:"infrastructureSpec,omitempty"`
}
