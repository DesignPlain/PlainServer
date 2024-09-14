package types

type Dns_ResponsePolicyRuleLocalData struct {
	/*
	   All resource record sets for this selector, one per resource record type. The name must match the dns_name.
	   Structure is documented below.
	*/
	LocalDatas []Dns_ResponsePolicyRuleLocalDataLocalData `json:"localDatas,omitempty" yaml:"localDatas,omitempty"`
}
