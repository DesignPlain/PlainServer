package ivs

import types "DesignSphere_Server/src/resource/aws/types"

type RecordingConfiguration struct {
	// Object containing destination configuration for where recorded video will be stored.
	DestinationConfiguration types.Ivs_RecordingConfigurationDestinationConfiguration `json:"destinationConfiguration,omitempty" yaml:"destinationConfiguration,omitempty"`

	// Recording Configuration name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	RecordingReconnectWindowSeconds int `json:"recordingReconnectWindowSeconds,omitempty" yaml:"recordingReconnectWindowSeconds,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
	ThumbnailConfiguration types.Ivs_RecordingConfigurationThumbnailConfiguration `json:"thumbnailConfiguration,omitempty" yaml:"thumbnailConfiguration,omitempty"`
}
