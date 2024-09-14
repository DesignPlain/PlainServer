package sagemaker

type NotebookInstanceLifecycleConfiguration struct {
	// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
	OnCreate string `json:"onCreate,omitempty" yaml:"onCreate,omitempty"`

	// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
	OnStart string `json:"onStart,omitempty" yaml:"onStart,omitempty"`
}
