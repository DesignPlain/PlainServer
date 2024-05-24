package redshiftserverless

type UsageLimit struct {
	// The action that Amazon Redshift Serverless takes when the limit is reached. Valid values are `log`, `emit-metric`, and `deactivate`. The default is `log`.
	BreachAction string `json:"breachAction,omitempty" yaml:"breachAction,omitempty"`

	// The time period that the amount applies to. A weekly period begins on Sunday. Valid values are `daily`, `weekly`, and `monthly`. The default is `monthly`.
	Period string `json:"period,omitempty" yaml:"period,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Redshift Serverless resource to create the usage limit for.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The type of Amazon Redshift Serverless usage to create a usage limit for. Valid values are `serverless-compute` or `cross-region-datasharing`.
	UsageType string `json:"usageType,omitempty" yaml:"usageType,omitempty"`

	// The limit amount. If time-based, this amount is in Redshift Processing Units (RPU) consumed per hour. If data-based, this amount is in terabytes (TB) of data transferred between Regions in cross-account sharing. The value must be a positive number.
	Amount int `json:"amount,omitempty" yaml:"amount,omitempty"`
}
