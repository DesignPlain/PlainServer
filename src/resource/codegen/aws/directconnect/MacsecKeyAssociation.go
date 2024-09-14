package directconnect

type MacsecKeyAssociation struct {
	// The MAC Security (MACsec) CKN to associate with the dedicated connection. The valid values are 64 hexadecimal characters (0-9, A-E). Required if using `cak`.
	Ckn string `json:"ckn,omitempty" yaml:"ckn,omitempty"`

	// The ID of the dedicated Direct Connect connection. The connection must be a dedicated connection in the `AVAILABLE` state.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	/*
	   The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key to associate with the dedicated connection.

	   > --Note:-- `ckn` and `cak` are mutually exclusive with `secret_arn` - these arguments cannot be used together. If you use `ckn` and `cak`, you should not use `secret_arn`. If you use the `secret_arn` argument to reference an existing MAC Security (MACSec) secret key, you should not use `ckn` or `cak`.
	*/
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// The MAC Security (MACsec) CAK to associate with the dedicated connection. The valid values are 64 hexadecimal characters (0-9, A-E). Required if using `ckn`.
	Cak string `json:"cak,omitempty" yaml:"cak,omitempty"`
}
