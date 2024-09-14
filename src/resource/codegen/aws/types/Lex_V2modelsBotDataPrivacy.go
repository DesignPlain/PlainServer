package types

type Lex_V2modelsBotDataPrivacy struct {
	// (Required) -  For each Amazon Lex bot created with the Amazon Lex Model Building Service, you must specify whether your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to the Children's Online Privacy Protection Act (COPPA) by specifying true or false in the childDirected field.
	ChildDirected bool `json:"childDirected,omitempty" yaml:"childDirected,omitempty"`
}
