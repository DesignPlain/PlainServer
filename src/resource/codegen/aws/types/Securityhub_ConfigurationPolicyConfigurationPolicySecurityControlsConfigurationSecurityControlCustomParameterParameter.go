package types

type Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameter struct {
	// The string list `value` for a EnumList-typed Security Hub Control Parameter.
	EnumList Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterEnumList `json:"enumList,omitempty" yaml:"enumList,omitempty"`

	// The int `value` for a Int-typed Security Hub Control Parameter.
	Int Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterInt `json:"int,omitempty" yaml:"int,omitempty"`

	// Identifies whether a control parameter uses a custom user-defined value or subscribes to the default Security Hub behavior. Valid values: `DEFAULT`, `CUSTOM`.
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`

	// The bool `value` for a Boolean-typed Security Hub Control Parameter.
	Bool Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterBool `json:"bool,omitempty" yaml:"bool,omitempty"`

	// The string `value` for a Enum-typed Security Hub Control Parameter.
	Enum Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterEnum `json:"enum,omitempty" yaml:"enum,omitempty"`

	// The name of the control parameter. For more information see the [Security Hub controls reference] documentation.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The string `value` for a String-typed Security Hub Control Parameter.
	String Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterString `json:"string,omitempty" yaml:"string,omitempty"`

	// The string list `value` for a StringList-typed Security Hub Control Parameter.
	StringList Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterStringList `json:"stringList,omitempty" yaml:"stringList,omitempty"`

	// The float `value` for a Double-typed Security Hub Control Parameter.
	Double Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterDouble `json:"double,omitempty" yaml:"double,omitempty"`

	// The int list `value` for a IntList-typed Security Hub Control Parameter.
	IntList Securityhub_ConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterIntList `json:"intList,omitempty" yaml:"intList,omitempty"`
}
