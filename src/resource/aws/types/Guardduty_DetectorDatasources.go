package types

type Guardduty_DetectorDatasources struct {
	/*
	   Configures [Malware Protection](https://docs.aws.amazon.com/guardduty/latest/ug/malware-protection.html).
	   See Malware Protection, Scan EC2 instance with findings and EBS volumes below for more details.

	   The `datasources` block is deprecated since March 2023. Use the `features` block instead and [map each `datasources` block to the corresponding `features` block](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-feature-object-api-changes-march2023.html#guardduty-feature-enablement-datasource-relation).
	*/
	MalwareProtection Guardduty_DetectorDatasourcesMalwareProtection `json:"malwareProtection,omitempty" yaml:"malwareProtection,omitempty"`

	/*
	   Configures [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
	   See S3 Logs below for more details.
	*/
	S3Logs Guardduty_DetectorDatasourcesS3Logs `json:"s3Logs,omitempty" yaml:"s3Logs,omitempty"`

	/*
	   Configures [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
	   See Kubernetes and Kubernetes Audit Logs below for more details.
	*/
	Kubernetes Guardduty_DetectorDatasourcesKubernetes `json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`
}
