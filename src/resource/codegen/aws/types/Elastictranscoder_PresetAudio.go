package types

type Elastictranscoder_PresetAudio struct {
	// The audio codec for the output file. Valid values are `AAC`, `flac`, `mp2`, `mp3`, `pcm`, and `vorbis`.
	Codec string `json:"codec,omitempty" yaml:"codec,omitempty"`

	// The sample rate of the audio stream in the output file, in hertz. Valid values are: `auto`, `22050`, `32000`, `44100`, `48000`, `96000`
	SampleRate string `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`

	// The method of organizing audio channels and tracks. Use Audio:Channels to specify the number of channels in your output, and Audio:AudioPackingMode to specify the number of tracks and their relation to the channels. If you do not specify an Audio:AudioPackingMode, Elastic Transcoder uses SingleTrack.
	AudioPackingMode string `json:"audioPackingMode,omitempty" yaml:"audioPackingMode,omitempty"`

	// The bit rate of the audio stream in the output file, in kilobits/second. Enter an integer between 64 and 320, inclusive.
	BitRate string `json:"bitRate,omitempty" yaml:"bitRate,omitempty"`

	// The number of audio channels in the output file
	Channels string `json:"channels,omitempty" yaml:"channels,omitempty"`
}
