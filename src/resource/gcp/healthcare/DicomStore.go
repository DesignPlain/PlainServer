package healthcare

import types "DesignSphere_Server/src/resource/gcp/types"

type DicomStore struct {
	/*
	   Identifies the dataset addressed by this request. Must be in the format
	   'projects/{project}/locations/{location}/datasets/{dataset}'


	   - - -
	*/
	Dataset string `json:"dataset,omitempty" yaml:"dataset,omitempty"`

	/*
	   User-supplied key-value pairs used to organize DICOM stores.
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
	   The resource name for the DicomStore.
	   -- Changing this property may recreate the Dicom store (removing all data) --
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	NotificationConfig types.Healthcare_DicomStoreNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	/*
	   To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
	   streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.
	   Structure is documented below.
	*/
	StreamConfigs []types.Healthcare_DicomStoreStreamConfig `json:"streamConfigs,omitempty" yaml:"streamConfigs,omitempty"`
}
