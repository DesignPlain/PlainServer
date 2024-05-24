package types

type Batch_getJobDefinitionEksProperty struct {
	// The properties for the Kubernetes pod resources of a job.
	PodProperties []string `json:"podProperties,omitempty" yaml:"podProperties,omitempty"`
}
