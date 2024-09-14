package dataflow

type Job struct {
	// The Service Account email used to create the job.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	// The GCS path to the Dataflow job template.
	TemplateGcsPath string `json:"templateGcsPath,omitempty" yaml:"templateGcsPath,omitempty"`

	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty" yaml:"enableStreamingEngine,omitempty"`

	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   User labels to be specified for the job. Keys and values should follow the restrictions
	   specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The machine type to use for the job.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A unique name for the resource, required by Dataflow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete string `json:"onDelete,omitempty" yaml:"onDelete,omitempty"`

	// The region in which the created job should run.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK". If the [subnetwork is located in a Shared VPC network](https://cloud.google.com/dataflow/docs/guides/specifying-networks#shared), you must use the complete URL. For example `"googleapis.com/compute/v1/projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET_NAME"`
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" yaml:"transformNameMapping,omitempty"`

	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// If set to `true`, Pulumi will treat `DRAINING` and `CANCELLING` as terminal states when deleting the resource, and will remove the resource from Pulumi state and move on.  See above note.
	SkipWaitOnJobTermination bool `json:"skipWaitOnJobTermination,omitempty" yaml:"skipWaitOnJobTermination,omitempty"`

	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// List of experiments that should be used by the job. An example value is `["enable_stackdriver_agent_metrics"]`.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" yaml:"additionalExperiments,omitempty"`

	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration string `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`

	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers int `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   A writeable location on GCS for the Dataflow job to dump its temporary data.

	   - - -
	*/
	TempGcsLocation string `json:"tempGcsLocation,omitempty" yaml:"tempGcsLocation,omitempty"`
}
