package types

type Sagemaker_UserProfileUserSettingsRStudioServerProAppSettings struct {
	// Indicates whether the current user has access to the RStudioServerPro app. Valid values are `ENABLED` and `DISABLED`.
	AccessStatus string `json:"accessStatus,omitempty" yaml:"accessStatus,omitempty"`

	// The level of permissions that the user has within the RStudioServerPro app. This value defaults to `R_STUDIO_USER`. The `R_STUDIO_ADMIN` value allows the user access to the RStudio Administrative Dashboard. Valid values are `R_STUDIO_USER` and `R_STUDIO_ADMIN`.
	UserGroup string `json:"userGroup,omitempty" yaml:"userGroup,omitempty"`
}
