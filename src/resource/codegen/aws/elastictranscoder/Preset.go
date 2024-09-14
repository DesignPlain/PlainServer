package elastictranscoder

import types "libds/aws/types"

type Preset struct {
	// Codec options for the video parameters
	VideoCodecOptions map[string]string `json:"videoCodecOptions,omitempty" yaml:"videoCodecOptions,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Video parameters object (documented below)
	Video types.Elastictranscoder_PresetVideo `json:"video,omitempty" yaml:"video,omitempty"`

	// Audio parameters object (documented below).
	Audio types.Elastictranscoder_PresetAudio `json:"audio,omitempty" yaml:"audio,omitempty"`

	// Codec options for the audio parameters (documented below)
	AudioCodecOptions types.Elastictranscoder_PresetAudioCodecOptions `json:"audioCodecOptions,omitempty" yaml:"audioCodecOptions,omitempty"`

	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container string `json:"container,omitempty" yaml:"container,omitempty"`

	// A description of the preset (maximum 255 characters)
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the preset. (maximum 40 characters)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Thumbnail parameters object (documented below)
	Thumbnails types.Elastictranscoder_PresetThumbnails `json:"thumbnails,omitempty" yaml:"thumbnails,omitempty"`

	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks []types.Elastictranscoder_PresetVideoWatermark `json:"videoWatermarks,omitempty" yaml:"videoWatermarks,omitempty"`
}
