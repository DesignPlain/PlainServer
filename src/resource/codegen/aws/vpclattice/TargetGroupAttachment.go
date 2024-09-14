package vpclattice

import types "libds/aws/types"

type TargetGroupAttachment struct {
	// The target.
	Target types.Vpclattice_TargetGroupAttachmentTarget `json:"target,omitempty" yaml:"target,omitempty"`

	// The ID or Amazon Resource Name (ARN) of the target group.
	TargetGroupIdentifier string `json:"targetGroupIdentifier,omitempty" yaml:"targetGroupIdentifier,omitempty"`
}
