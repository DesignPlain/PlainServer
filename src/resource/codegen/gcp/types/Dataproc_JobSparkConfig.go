package types

type Dataproc_JobSparkConfig struct {
	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty" yaml:"jarFileUris,omitempty"`

	// The runtime logging config of the job
	LoggingConfig Dataproc_JobSparkConfigLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	/*
	   The class containing the main method of the driver. Must be in a
	   provided jar or jar that is already on the classpath. Conflicts with `main_jar_file_uri`
	*/
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	/*
	   The HCFS URI of jar file containing
	   the driver jar. Conflicts with `main_class`
	*/
	MainJarFileUri string `json:"mainJarFileUri,omitempty" yaml:"mainJarFileUri,omitempty"`

	/*
	   A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in `/etc/spark/conf/spark-defaults.conf` and classes in user code.

	   - `logging_config.driver_log_levels`- (Required) The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	*/
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty" yaml:"archiveUris,omitempty"`

	// The arguments to pass to the driver.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty" yaml:"fileUris,omitempty"`
}
