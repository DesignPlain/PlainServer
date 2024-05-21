package eventarc

import types "DesignSphere_Server/src/resource/gcp/types"

type Trigger struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`

	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport types.Eventarc_TriggerTransport `json:"transport,omitempty" yaml:"transport,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias []types.Eventarc_TriggerMatchingCriteria `json:"matchingCriterias,omitempty" yaml:"matchingCriterias,omitempty"`

	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Required. Destination specifies where the events should be sent to.
	Destination types.Eventarc_TriggerDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType string `json:"eventDataContentType,omitempty" yaml:"eventDataContentType,omitempty"`

	/*
	   Optional. User labels attached to the triggers that can be used to group resources.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
