package healthcare

import types "DesignSphere_Server/src/resource/gcp/types"

type Hl7Store struct {
	/*
	   User-supplied key-value pairs used to organize HL7v2 stores.
	   Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	   conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	   Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	   bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	   No more than 64 labels can be associated with a given store.
	   An object containing a list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The resource name for the Hl7V2Store.
	   -- Changing this property may recreate the Hl7v2 store (removing all data) --
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   (Optional, Deprecated)
	   A nested object resource
	   Structure is documented below.

	   > --Warning:-- `notification_config` is deprecated and will be removed in a future major release. Use `notification_configs` instead.
	*/
	NotificationConfig types.Healthcare_Hl7StoreNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	/*
	   A list of notification configs. Each configuration uses a filter to determine whether to publish a
	   message (both Ingest & Create) on the corresponding notification destination. Only the message name
	   is sent as part of the notification. Supplied by the client.
	   Structure is documented below.
	*/
	NotificationConfigs []types.Healthcare_Hl7StoreNotificationConfigs `json:"notificationConfigs,omitempty" yaml:"notificationConfigs,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	ParserConfig types.Healthcare_Hl7StoreParserConfig `json:"parserConfig,omitempty" yaml:"parserConfig,omitempty"`

	// Determines whether duplicate messages are allowed.
	RejectDuplicateMessage bool `json:"rejectDuplicateMessage,omitempty" yaml:"rejectDuplicateMessage,omitempty"`

	/*
	   Identifies the dataset addressed by this request. Must be in the format
	   'projects/{project}/locations/{location}/datasets/{dataset}'


	   - - -
	*/
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`
}
