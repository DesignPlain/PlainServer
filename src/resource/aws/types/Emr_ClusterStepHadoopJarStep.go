package types

type Emr_ClusterStepHadoopJarStep struct {
	// List of command line arguments passed to the JAR file's main function when executed.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// Path to a JAR file run during the step.
	Jar string `json:"jar,omitempty" yaml:"jar,omitempty"`

	// Name of the main class in the specified Java file. If not specified, the JAR file should specify a Main-Class in its manifest file.
	MainClass string `json:"mainClass,omitempty" yaml:"mainClass,omitempty"`

	// Key-Value map of Java properties that are set when the step runs. You can use these properties to pass key value pairs to your main function.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
}
