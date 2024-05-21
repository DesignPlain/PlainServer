package cloudbuild

import types "DesignSphere_Server/src/resource/gcp/types"

type WorkerPool struct {
	// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   User-defined name of the `WorkerPool`.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Network configuration for the `WorkerPool`. Structure is documented below.
	NetworkConfig types.Cloudbuild_WorkerPoolNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
	WorkerConfig types.Cloudbuild_WorkerPoolWorkerConfig `json:"workerConfig,omitempty" yaml:"workerConfig,omitempty"`

	/*
	   User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
	   limitations. --Note--: This field is non-authoritative, and will only manage the annotations present in your
	   configuration. Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
