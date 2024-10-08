package types

type Dataflow_PipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment struct {
	// Additional experiment flags for the job.
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" yaml:"additionalExperiments,omitempty"`

	/*
	   Configuration for VM IPs.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration
	   Possible values are: `WORKER_IP_UNSPECIFIED`, `WORKER_IP_PUBLIC`, `WORKER_IP_PRIVATE`.
	*/
	IpConfiguration string `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`

	// 'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	// The machine type to use for the job. Defaults to the value from the template if not specified.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.
	MaxWorkers int `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`

	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.
	TempLocation string `json:"tempLocation,omitempty" yaml:"tempLocation,omitempty"`

	// The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.
	WorkerRegion string `json:"workerRegion,omitempty" yaml:"workerRegion,omitempty"`

	// The email address of the service account to run the job as.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   Additional user labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.
	   'Example: { "name": "wrench", "mass": "1kg", "count": "3" }.'
	   'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	*/
	AdditionalUserLabels map[string]string `json:"additionalUserLabels,omitempty" yaml:"additionalUserLabels,omitempty"`

	// Whether to enable Streaming Engine for the job.
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty" yaml:"enableStreamingEngine,omitempty"`

	// Network to which VMs will be assigned. If empty or unspecified, the service will use the network "default".
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The initial number of Compute Engine instances for the job.
	NumWorkers int `json:"numWorkers,omitempty" yaml:"numWorkers,omitempty"`

	// Subnetwork to which VMs will be assigned, if desired. You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK" or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1-a". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.
	WorkerZone string `json:"workerZone,omitempty" yaml:"workerZone,omitempty"`

	/*
	   Set FlexRS goal for the job. https://cloud.google.com/dataflow/docs/guides/flexrs
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#FlexResourceSchedulingGoal
	   Possible values are: `FLEXRS_UNSPECIFIED`, `FLEXRS_SPEED_OPTIMIZED`, `FLEXRS_COST_OPTIMIZED`.
	*/
	FlexrsGoal string `json:"flexrsGoal,omitempty" yaml:"flexrsGoal,omitempty"`

	// The Compute Engine availability zone for launching worker instances to run your pipeline. In the future, workerZone will take precedence.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
