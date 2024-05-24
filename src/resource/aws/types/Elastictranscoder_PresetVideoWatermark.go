package types

type Elastictranscoder_PresetVideoWatermark struct {
	// A unique identifier for the settings for one watermark. The value of Id can be up to 40 characters long. You can specify settings for up to four watermarks.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The maximum height of the watermark.
	MaxHeight string `json:"maxHeight,omitempty" yaml:"maxHeight,omitempty"`

	// The maximum width of the watermark.
	MaxWidth string `json:"maxWidth,omitempty" yaml:"maxWidth,omitempty"`

	// A percentage that indicates how much you want a watermark to obscure the video in the location where it appears.
	Opacity string `json:"opacity,omitempty" yaml:"opacity,omitempty"`

	// A value that determines how Elastic Transcoder interprets values that you specified for `video_watermarks.horizontal_offset`, `video_watermarks.vertical_offset`, `video_watermarks.max_width`, and `video_watermarks.max_height`. Valid values are `Content` and `Frame`.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// The vertical position of the watermark unless you specify a nonzero value for `vertical_align`. Valid values are `Top`, `Bottom`, `Center`.
	VerticalAlign string `json:"verticalAlign,omitempty" yaml:"verticalAlign,omitempty"`

	// The horizontal position of the watermark unless you specify a nonzero value for `horzontal_offset`.
	HorizontalAlign string `json:"horizontalAlign,omitempty" yaml:"horizontalAlign,omitempty"`

	// A value that controls scaling of the watermark. Valid values are: `Fit`, `Stretch`, `ShrinkToFit`
	SizingPolicy string `json:"sizingPolicy,omitempty" yaml:"sizingPolicy,omitempty"`

	// The amount by which you want the vertical position of the watermark to be offset from the position specified by `vertical_align`
	VerticalOffset string `json:"verticalOffset,omitempty" yaml:"verticalOffset,omitempty"`

	// The amount by which you want the horizontal position of the watermark to be offset from the position specified by `horizontal_align`.
	HorizontalOffset string `json:"horizontalOffset,omitempty" yaml:"horizontalOffset,omitempty"`
}
