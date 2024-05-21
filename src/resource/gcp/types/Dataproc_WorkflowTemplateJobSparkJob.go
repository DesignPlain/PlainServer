package types

type Dataproc_WorkflowTemplateJobSparkJob struct {
	// The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in `jar_file_uris`.
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	// The HCFS URI of the jar file that contains the main class.
	MainJarFileUri string `json:"mainJarFileUri,omitempty" yaml:"mainJarFileUri,omitempty"`

	// A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobSparkJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`
}
