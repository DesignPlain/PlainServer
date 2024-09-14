package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroup struct {
	// The key of the application execution property key-value map.
	PropertyGroupId string `json:"propertyGroupId,omitempty" yaml:"propertyGroupId,omitempty"`

	// Application execution property key-value map.
	PropertyMap map[string]string `json:"propertyMap,omitempty" yaml:"propertyMap,omitempty"`
}
