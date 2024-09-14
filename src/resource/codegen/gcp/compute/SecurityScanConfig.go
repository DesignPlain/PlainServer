package compute

import types "libds/gcp/types"

type SecurityScanConfig struct {
	/*
	   Controls export of scan configurations and results to Cloud Security Command Center.
	   Default value is `ENABLED`.
	   Possible values are: `ENABLED`, `DISABLED`.
	*/
	ExportToSecurityCommandCenter string `json:"exportToSecurityCommandCenter,omitempty" yaml:"exportToSecurityCommandCenter,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The schedule of the ScanConfig
	   Structure is documented below.
	*/
	Schedule types.Compute_SecurityScanConfigSchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	/*
	   The starting URLs from which the scanner finds site pages.


	   - - -
	*/
	StartingUrls []string `json:"startingUrls,omitempty" yaml:"startingUrls,omitempty"`

	/*
	   Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
	   Each value may be one of: `APP_ENGINE`, `COMPUTE`.
	*/
	TargetPlatforms []string `json:"targetPlatforms,omitempty" yaml:"targetPlatforms,omitempty"`

	// The user provider display name of the ScanConfig.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The blacklist URL patterns as described in
	   https://cloud.google.com/security-scanner/docs/excluded-urls
	*/
	BlacklistPatterns []string `json:"blacklistPatterns,omitempty" yaml:"blacklistPatterns,omitempty"`

	/*
	   The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
	   Defaults to 15.
	*/
	MaxQps int `json:"maxQps,omitempty" yaml:"maxQps,omitempty"`

	/*
	   Type of the user agents used for scanning
	   Default value is `CHROME_LINUX`.
	   Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
	*/
	UserAgent string `json:"userAgent,omitempty" yaml:"userAgent,omitempty"`

	/*
	   The authentication configuration.
	   If specified, service will use the authentication configuration during scanning.
	   Structure is documented below.
	*/
	Authentication types.Compute_SecurityScanConfigAuthentication `json:"authentication,omitempty" yaml:"authentication,omitempty"`
}
