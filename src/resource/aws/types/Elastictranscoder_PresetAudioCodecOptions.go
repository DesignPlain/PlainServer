package types

type Elastictranscoder_PresetAudioCodecOptions struct {
	// The order the bits of a PCM sample are stored in. The supported value is LittleEndian. (PCM Only)
	BitOrder string `json:"bitOrder,omitempty" yaml:"bitOrder,omitempty"`

	// If you specified AAC for Audio:Codec, choose the AAC profile for the output file.
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`

	// Whether audio samples are represented with negative and positive numbers (signed) or only positive numbers (unsigned). The supported value is Signed. (PCM Only)
	Signed string `json:"signed,omitempty" yaml:"signed,omitempty"`

	// The bit depth of a sample is how many bits of information are included in the audio samples. Valid values are `16` and `24`. (FLAC/PCM Only)
	BitDepth string `json:"bitDepth,omitempty" yaml:"bitDepth,omitempty"`
}
