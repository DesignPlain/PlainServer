package types

type Location_MapConfiguration struct {
	// Specifies the map style selected from an available data provider. Valid values can be found in the [Location Service CreateMap API Reference](https://docs.aws.amazon.com/location/latest/APIReference/API_CreateMap.html).
	Style string `json:"style,omitempty" yaml:"style,omitempty"`
}
