package securityhub

type StandardsSubscription struct {
	/*
	   The ARN of a standard - see below.

	   Currently available standards (remember to replace `${var.partition}` and `${var.region}` as appropriate):

	   | Name                                     | ARN                                                                                                          |
	   |------------------------------------------|--------------------------------------------------------------------------------------------------------------|
	   | AWS Foundational Security Best Practices | `arn:${var.partition}:securityhub:${var.region}::standards/aws-foundational-security-best-practices/v/1.0.0` |
	   | AWS Resource Tagging Standard            | `arn:${var.partition}:securityhub:${var.region}::standards/aws-resource-tagging-standard/v/1.0.0`            |
	   | CIS AWS Foundations Benchmark v1.2.0     | `arn:${var.partition}:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0`                           |
	   | CIS AWS Foundations Benchmark v1.4.0     | `arn:${var.partition}:securityhub:${var.region}::standards/cis-aws-foundations-benchmark/v/1.4.0`            |
	   | NIST SP 800-53 Rev. 5                    | `arn:${var.partition}:securityhub:${var.region}::standards/nist-800-53/v/5.0.0`                              |
	   | PCI DSS                                  | `arn:${var.partition}:securityhub:${var.region}::standards/pci-dss/v/3.2.1`                                  |
	*/
	StandardsArn string `json:"standardsArn,omitempty" yaml:"standardsArn,omitempty"`
}
