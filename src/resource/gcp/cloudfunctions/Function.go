package cloudfunctions

import types "DesignSphere_Server/src/resource/gcp/types"

type Function struct {
	// Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
	AvailableMemoryMb int `json:"availableMemoryMb,omitempty" yaml:"availableMemoryMb,omitempty"`

	// A set of key/value environment variable pairs available during build time.
	BuildEnvironmentVariables map[string]string `json:"buildEnvironmentVariables,omitempty" yaml:"buildEnvironmentVariables,omitempty"`

	/*
	   The runtime in which the function is going to run.
	   Eg. `"nodejs16"`, `"python39"`, `"dotnet3"`, `"go116"`, `"java11"`, `"ruby30"`, `"php74"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.

	   - - -
	*/
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// Secret environment variables configuration. Structure is documented below.
	SecretEnvironmentVariables []types.Cloudfunctions_FunctionSecretEnvironmentVariable `json:"secretEnvironmentVariables,omitempty" yaml:"secretEnvironmentVariables,omitempty"`

	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	BuildWorkerPool string `json:"buildWorkerPool,omitempty" yaml:"buildWorkerPool,omitempty"`

	// The security level for the function. The following options are available:
	HttpsTriggerSecurityLevel string `json:"httpsTriggerSecurityLevel,omitempty" yaml:"httpsTriggerSecurityLevel,omitempty"`

	// Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `event_trigger`.
	TriggerHttp bool `json:"triggerHttp,omitempty" yaml:"triggerHttp,omitempty"`

	// The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/-/locations/-/connectors/-`.
	VpcConnector string `json:"vpcConnector,omitempty" yaml:"vpcConnector,omitempty"`

	// A set of key/value environment variable pairs to assign to the function.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// The limit on the maximum number of function instances that may coexist at a given time.
	MaxInstances int `json:"maxInstances,omitempty" yaml:"maxInstances,omitempty"`

	// Project of the function. If it is not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Secret volumes configuration. Structure is documented below.
	SecretVolumes []types.Cloudfunctions_FunctionSecretVolume `json:"secretVolumes,omitempty" yaml:"secretVolumes,omitempty"`

	// The source archive object (file) in archive bucket.
	SourceArchiveObject string `json:"sourceArchiveObject,omitempty" yaml:"sourceArchiveObject,omitempty"`

	// Description of the function.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
	DockerRepository string `json:"dockerRepository,omitempty" yaml:"dockerRepository,omitempty"`

	// The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
	VpcConnectorEgressSettings string `json:"vpcConnectorEgressSettings,omitempty" yaml:"vpcConnectorEgressSettings,omitempty"`

	/*
	   A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// If provided, the self-provided service account to run the function with.
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	// URL which triggers function execution. Returned only if `trigger_http` is used.
	HttpsTriggerUrl string `json:"httpsTriggerUrl,omitempty" yaml:"httpsTriggerUrl,omitempty"`

	/*
	   Represents parameters related to source repository where a function is hosted.
	   Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.-
	*/
	SourceRepository types.Cloudfunctions_FunctionSourceRepository `json:"sourceRepository,omitempty" yaml:"sourceRepository,omitempty"`

	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	   If specified, you must also provide an artifact registry repository using the `docker_repository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	// A user-defined name of the function. Function names must be unique globally.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Region of function. If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket string `json:"sourceArchiveBucket,omitempty" yaml:"sourceArchiveBucket,omitempty"`

	// Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
	DockerRegistry string `json:"dockerRegistry,omitempty" yaml:"dockerRegistry,omitempty"`

	// Name of the function that will be executed when the Google Cloud Function is triggered.
	EntryPoint string `json:"entryPoint,omitempty" yaml:"entryPoint,omitempty"`

	// A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
	EventTrigger types.Cloudfunctions_FunctionEventTrigger `json:"eventTrigger,omitempty" yaml:"eventTrigger,omitempty"`

	// String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
	IngressSettings string `json:"ingressSettings,omitempty" yaml:"ingressSettings,omitempty"`

	// The limit on the minimum number of function instances that may coexist at a given time.
	MinInstances int `json:"minInstances,omitempty" yaml:"minInstances,omitempty"`
}
