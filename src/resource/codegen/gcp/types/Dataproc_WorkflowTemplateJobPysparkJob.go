package types

type Dataproc_WorkflowTemplateJobPysparkJob struct {
	// HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.
	PythonFileUris []string `json:"pythonFileUris,omitempty" yaml:"pythonFileUris,omitempty"`

	// HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobPysparkJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file.
	MainPythonFileUri string `json:"mainPythonFileUri,omitempty" yaml:"mainPythonFileUri,omitempty"`

	// A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
