package types

type Dataproc_JobHadoopConfig struct {
	// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime logging config of the job
	LoggingConfig Dataproc_JobHadoopConfigLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	// The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
	MainJarFileUri string `json:"mainJarFileUri,omitempty" yaml:"mainJarFileUri,omitempty"`

	/*
	   A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in `/etc/hadoop/conf/--site` and classes in user code..

	   - `logging_config.driver_log_levels`- (Required) The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	*/
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
