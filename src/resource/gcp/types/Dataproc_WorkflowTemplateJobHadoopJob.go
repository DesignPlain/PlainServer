package types

type Dataproc_WorkflowTemplateJobHadoopJob struct {
	// HCFS URIs of archives to be extracted in the working directory of Hadoop drivers and tasks. Supported file types: .jar, .tar, .tar.gz, .tgz, or .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// The arguments to pass to the driver. Do not include arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// HCFS (Hadoop Compatible Filesystem) URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	// Jar file URIs to add to the CLASSPATHs of the Hadoop driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime log config for job execution.
	LoggingConfig Dataproc_WorkflowTemplateJobHadoopJobLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`.
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	// The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'
	MainJarFileUri string `json:"mainJarFileUri,omitempty" yaml:"mainJarFileUri,omitempty"`

	// A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/--site and classes in user code.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
