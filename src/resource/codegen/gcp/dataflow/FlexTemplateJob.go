package dataflow

type FlexTemplateJob struct {
	/*
	   Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the
	   corresponding name prefixes of the new job.
	*/
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" yaml:"transformNameMapping,omitempty"`

	/*
	   The GCS path to the Dataflow job Flex
	   Template.

	   - - -
	*/
	ContainerSpecGcsPath string `json:"containerSpecGcsPath,omitempty" yaml:"containerSpecGcsPath,omitempty"`

	/*
	   The name for the Cloud KMS key for the job. Key format is:
	   projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   User labels to be specified for the job. Keys and values
	   should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	   page. --Note--: This field is marked as deprecated as the API does not currently
	   support adding labels.
	   --NOTE--: Google-provided Dataflow templates often provide default labels
	   that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	   labels will be ignored to prevent diffs on re-apply.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The machine type to use for launching the job. The default is n1-standard-1.
	LauncherMachineType string `json:"launcherMachineType,omitempty" yaml:"launcherMachineType,omitempty"`

	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Docker registry location of container image to use for the 'worker harness. Default is the container for the version of
	   the SDK. Note this field is only valid for portable pipelines.
	*/
	SdkContainerImage string `json:"sdkContainerImage,omitempty" yaml:"sdkContainerImage,omitempty"`

	// The Cloud Storage path to use for staging files. Must be a valid Cloud Storage URL, beginning with gs://.
	StagingLocation string `json:"stagingLocation,omitempty" yaml:"stagingLocation,omitempty"`

	/*
	   If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from
	   terraform state and moving on. WARNING: this will lead to job name conflicts if you do not ensure that the job names are
	   different, e.g. by embedding a release ID or by using a random_id.
	*/
	SkipWaitOnJobTermination bool `json:"skipWaitOnJobTermination,omitempty" yaml:"skipWaitOnJobTermination,omitempty"`

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" yaml:"additionalExperiments,omitempty"`

	// Indicates if the job should use the streaming engine feature.
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty" yaml:"enableStreamingEngine,omitempty"`

	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	IpConfiguration string `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`

	// The machine type to use for the job.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// A unique name for the resource, required by Dataflow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   One of "drain" or "cancel". Specifies behavior of
	   deletion during `pulumi destroy`.  See above note.
	*/
	OnDelete string `json:"onDelete,omitempty" yaml:"onDelete,omitempty"`

	/*
	   The project in which the resource belongs. If it is not
	   provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.
	TempLocation string `json:"tempLocation,omitempty" yaml:"tempLocation,omitempty"`

	// The algorithm to use for autoscaling
	AutoscalingAlgorithm string `json:"autoscalingAlgorithm,omitempty" yaml:"autoscalingAlgorithm,omitempty"`

	/*
	   The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to
	   1000.
	*/
	MaxWorkers int `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`

	// The region in which the created job should run.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The initial number of Google Compute Engine instances for the job.
	NumWorkers int `json:"numWorkers,omitempty" yaml:"numWorkers,omitempty"`

	/*
	   Key/Value pairs to be passed to the Dataflow job (as
	   used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	   such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The Service Account email used to create the job.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
}
