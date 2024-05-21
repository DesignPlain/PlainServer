package datastream

import types "DesignSphere_Server/src/resource/gcp/types"

type Stream struct {
	/*
	   Source connection profile configuration.
	   Structure is documented below.
	*/
	SourceConfig types.Datastream_StreamSourceConfig `json:"sourceConfig,omitempty" yaml:"sourceConfig,omitempty"`

	// The stream identifier.
	StreamId string `json:"streamId,omitempty" yaml:"streamId,omitempty"`

	// Desired state of the Stream. Set this field to `RUNNING` to start the stream, and `PAUSED` to pause the stream.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	// Display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Destination connection profile configuration.
	   Structure is documented below.
	*/
	DestinationConfig types.Datastream_StreamDestinationConfig `json:"destinationConfig,omitempty" yaml:"destinationConfig,omitempty"`

	/*
	   Labels.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of the location this stream is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
	   Structure is documented below.
	*/
	BackfillAll types.Datastream_StreamBackfillAll `json:"backfillAll,omitempty" yaml:"backfillAll,omitempty"`

	// Backfill strategy to disable automatic backfill for the Stream's objects.
	BackfillNone types.Datastream_StreamBackfillNone `json:"backfillNone,omitempty" yaml:"backfillNone,omitempty"`

	/*
	   A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
	   will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	*/
	CustomerManagedEncryptionKey string `json:"customerManagedEncryptionKey,omitempty" yaml:"customerManagedEncryptionKey,omitempty"`
}
