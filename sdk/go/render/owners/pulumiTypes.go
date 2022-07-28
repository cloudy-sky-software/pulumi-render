// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package owners

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ListOwnersResponse struct {
	Cursor *string `pulumi:"cursor"`
	// The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
	Owner *Owner `pulumi:"owner"`
}

// The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
type Owner struct {
	// The email of the owner.
	Email *string `pulumi:"email"`
	// The owner ID.
	Id *string `pulumi:"id"`
	// The name of the owner.
	Name *string `pulumi:"name"`
	// The type of the authorized user.
	Type *OwnerType `pulumi:"type"`
}

// The owner object represents an authorized user or team. The `type` property indicates if the owner is a user or team.
type OwnerOutput struct{ *pulumi.OutputState }

func (OwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Owner)(nil)).Elem()
}

func (o OwnerOutput) ToOwnerOutput() OwnerOutput {
	return o
}

func (o OwnerOutput) ToOwnerOutputWithContext(ctx context.Context) OwnerOutput {
	return o
}

// The email of the owner.
func (o OwnerOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Owner) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The owner ID.
func (o OwnerOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Owner) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the owner.
func (o OwnerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Owner) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The type of the authorized user.
func (o OwnerOutput) Type() OwnerTypePtrOutput {
	return o.ApplyT(func(v Owner) *OwnerType { return v.Type }).(OwnerTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OwnerOutput{})
}
