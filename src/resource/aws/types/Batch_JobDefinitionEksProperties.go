package types

type Batch_JobDefinitionEksProperties struct {
	// The properties for the Kubernetes pod resources of a job. See `pod_properties` below.
	PodProperties Batch_JobDefinitionEksPropertiesPodProperties `json:"podProperties,omitempty" yaml:"podProperties,omitempty"`
}
