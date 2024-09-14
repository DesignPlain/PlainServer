package cleanrooms

import types "libds/aws/types"

type Collaboration struct {
	// The list of member abilities for the creator of the collaboration.  Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbilities).
	CreatorMemberAbilities []string `json:"creatorMemberAbilities,omitempty" yaml:"creatorMemberAbilities,omitempty"`

	/*
	   a collection of settings which determine how the [c3r client](https://docs.aws.amazon.com/clean-rooms/latest/userguide/crypto-computing.html) will encrypt data for use within this collaboration.
	   - `data_encryption_metadata.allow_clear_text` - (Required - Forces new resource) - Indicates whether encrypted tables can contain cleartext data. This is a boolea
	   field.
	   - `data_encryption_metadata.allow_duplicates` - (Required - Forces new resource ) - Indicates whether Fingerprint columns can contain duplicate entries. This is a
	   boolean field.
	   - `data_encryption_metadata.allow_joins_on_columns_with_different_names` - (Required - Forces new resource) - Indicates whether Fingerprint columns can be joined
	   n any other Fingerprint column with a different name. This is a boolean field.
	   - `data_encryption_metadata.preserve_nulls` - (Required - Forces new resource) - Indicates whether NULL values are to be copied as NULL to encrypted tables (true)
	   or cryptographically processed (false).
	*/
	DataEncryptionMetadata types.Cleanrooms_CollaborationDataEncryptionMetadata `json:"dataEncryptionMetadata,omitempty" yaml:"dataEncryptionMetadata,omitempty"`

	// A description for a collaboration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Additional members of the collaboration which will be invited to join the collaboration.
	   - `member.account_id` - (Required - Forces new resource) - The account id for the invited member.
	   - `member.display_name` - (Required - Forces new resource) - The display name for the invited member.
	   - `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbilities).
	*/
	Members []types.Cleanrooms_CollaborationMember `json:"members,omitempty" yaml:"members,omitempty"`

	// The name of the collaboration.  Collaboration names do not need to be unique.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Determines if members of the collaboration can enable query logs within their own.
	   emberships. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-queryLogStatus).
	*/
	QueryLogStatus string `json:"queryLogStatus,omitempty" yaml:"queryLogStatus,omitempty"`

	// Key value pairs which tag the collaboration.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name for the member record for the collaboration creator.
	CreatorDisplayName string `json:"creatorDisplayName,omitempty" yaml:"creatorDisplayName,omitempty"`
}
