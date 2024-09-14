package types

type Dataplex_TaskNotebookInfrastructureSpecContainerImage struct {
	// Container image to use.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// A list of Java JARS to add to the classpath. Valid input includes Cloud Storage URIs to Jar binaries. For example, gs://bucket-name/my/path/to/file.jar
	JavaJars []string `json:"javaJars,omitempty" yaml:"javaJars,omitempty"`

	// Override to common configuration of open source components installed on the Dataproc cluster. The properties to set on daemon config files. Property keys are specified in prefix:property format, for example core:hadoop.tmp.dir. For more information, see Cluster properties.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// A list of python packages to be installed. Valid formats include Cloud Storage URI to a PIP installable library. For example, gs://bucket-name/my/path/to/lib.tar.gz
	PythonPackages []string `json:"pythonPackages,omitempty" yaml:"pythonPackages,omitempty"`
}
