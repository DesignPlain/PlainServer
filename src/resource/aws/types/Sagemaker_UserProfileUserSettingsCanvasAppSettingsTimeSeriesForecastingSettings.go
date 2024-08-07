package types

type Sagemaker_UserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings struct {
	// The IAM role that Canvas passes to Amazon Forecast for time series forecasting. By default, Canvas uses the execution role specified in the UserProfile that launches the Canvas app. If an execution role is not specified in the UserProfile, Canvas uses the execution role specified in the Domain that owns the UserProfile. To allow time series forecasting, this IAM role should have the [AmazonSageMakerCanvasForecastAccess](https://docs.aws.amazon.com/sagemaker/latest/dg/security-iam-awsmanpol-canvas.html#security-iam-awsmanpol-AmazonSageMakerCanvasForecastAccess) policy attached and forecast.amazonaws.com added in the trust relationship as a service principal.
	AmazonForecastRoleArn string `json:"amazonForecastRoleArn,omitempty" yaml:"amazonForecastRoleArn,omitempty"`

	// Describes whether time series forecasting is enabled or disabled in the Canvas app. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
