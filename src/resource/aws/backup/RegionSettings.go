package backup

type RegionSettings struct {
	// A map of services along with the management preferences for the Region. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_UpdateRegionSettings.html#API_UpdateRegionSettings_RequestSyntax).
	ResourceTypeManagementPreference map[string]bool `json:"resourceTypeManagementPreference,omitempty" yaml:"resourceTypeManagementPreference,omitempty"`

	// A map of services along with the opt-in preferences for the Region.
	ResourceTypeOptInPreference map[string]bool `json:"resourceTypeOptInPreference,omitempty" yaml:"resourceTypeOptInPreference,omitempty"`
}
