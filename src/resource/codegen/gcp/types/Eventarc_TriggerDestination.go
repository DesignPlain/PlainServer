package types

type Eventarc_TriggerDestination struct {
	// An HTTP endpoint destination described by an URI.
	HttpEndpoint Eventarc_TriggerDestinationHttpEndpoint `json:"httpEndpoint,omitempty" yaml:"httpEndpoint,omitempty"`

	// Optional. Network config is used to configure how Eventarc resolves and connect to a destination. This should only be used with HttpEndpoint destination type.
	NetworkConfig Eventarc_TriggerDestinationNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	// The resource name of the Workflow whose Executions are triggered by the events. The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	Workflow string `json:"workflow,omitempty" yaml:"workflow,omitempty"`

	// The Cloud Function resource name. Only Cloud Functions V2 is supported. Format projects/{project}/locations/{location}/functions/{function} This is a read-only field. [WARNING] Creating Cloud Functions V2 triggers is only supported via the Cloud Functions product. An error will be returned if the user sets this value.
	CloudFunction string `json:"cloudFunction,omitempty" yaml:"cloudFunction,omitempty"`

	// Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.
	CloudRunService Eventarc_TriggerDestinationCloudRunService `json:"cloudRunService,omitempty" yaml:"cloudRunService,omitempty"`

	// A GKE service capable of receiving events. The service should be running in the same project as the trigger.
	Gke Eventarc_TriggerDestinationGke `json:"gke,omitempty" yaml:"gke,omitempty"`
}
