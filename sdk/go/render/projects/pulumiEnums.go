// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OwnerType string

const (
	OwnerTypeUser = OwnerType("user")
	OwnerTypeTeam = OwnerType("team")
)

type OwnerTypeOutput struct{ *pulumi.OutputState }

func (OwnerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OwnerType)(nil)).Elem()
}

func (o OwnerTypeOutput) ToOwnerTypeOutput() OwnerTypeOutput {
	return o
}

func (o OwnerTypeOutput) ToOwnerTypeOutputWithContext(ctx context.Context) OwnerTypeOutput {
	return o
}

func (o OwnerTypeOutput) ToOwnerTypePtrOutput() OwnerTypePtrOutput {
	return o.ToOwnerTypePtrOutputWithContext(context.Background())
}

func (o OwnerTypeOutput) ToOwnerTypePtrOutputWithContext(ctx context.Context) OwnerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OwnerType) *OwnerType {
		return &v
	}).(OwnerTypePtrOutput)
}

func (o OwnerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OwnerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwnerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OwnerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwnerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwnerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OwnerTypePtrOutput struct{ *pulumi.OutputState }

func (OwnerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OwnerType)(nil)).Elem()
}

func (o OwnerTypePtrOutput) ToOwnerTypePtrOutput() OwnerTypePtrOutput {
	return o
}

func (o OwnerTypePtrOutput) ToOwnerTypePtrOutputWithContext(ctx context.Context) OwnerTypePtrOutput {
	return o
}

func (o OwnerTypePtrOutput) Elem() OwnerTypeOutput {
	return o.ApplyT(func(v *OwnerType) OwnerType {
		if v != nil {
			return *v
		}
		var ret OwnerType
		return ret
	}).(OwnerTypeOutput)
}

func (o OwnerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwnerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OwnerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
type ProjectCreateEnvironmentInputProtectedStatus string

const (
	ProjectCreateEnvironmentInputProtectedStatusUnprotected = ProjectCreateEnvironmentInputProtectedStatus("unprotected")
	ProjectCreateEnvironmentInputProtectedStatusProtected   = ProjectCreateEnvironmentInputProtectedStatus("protected")
)

func (ProjectCreateEnvironmentInputProtectedStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectCreateEnvironmentInputProtectedStatus)(nil)).Elem()
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToProjectCreateEnvironmentInputProtectedStatusOutput() ProjectCreateEnvironmentInputProtectedStatusOutput {
	return pulumi.ToOutput(e).(ProjectCreateEnvironmentInputProtectedStatusOutput)
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToProjectCreateEnvironmentInputProtectedStatusOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProjectCreateEnvironmentInputProtectedStatusOutput)
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToProjectCreateEnvironmentInputProtectedStatusPtrOutput() ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return e.ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(context.Background())
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return ProjectCreateEnvironmentInputProtectedStatus(e).ToProjectCreateEnvironmentInputProtectedStatusOutputWithContext(ctx).ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(ctx)
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectCreateEnvironmentInputProtectedStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProjectCreateEnvironmentInputProtectedStatusOutput struct{ *pulumi.OutputState }

func (ProjectCreateEnvironmentInputProtectedStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectCreateEnvironmentInputProtectedStatus)(nil)).Elem()
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToProjectCreateEnvironmentInputProtectedStatusOutput() ProjectCreateEnvironmentInputProtectedStatusOutput {
	return o
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToProjectCreateEnvironmentInputProtectedStatusOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusOutput {
	return o
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToProjectCreateEnvironmentInputProtectedStatusPtrOutput() ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return o.ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(context.Background())
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectCreateEnvironmentInputProtectedStatus) *ProjectCreateEnvironmentInputProtectedStatus {
		return &v
	}).(ProjectCreateEnvironmentInputProtectedStatusPtrOutput)
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectCreateEnvironmentInputProtectedStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectCreateEnvironmentInputProtectedStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProjectCreateEnvironmentInputProtectedStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProjectCreateEnvironmentInputProtectedStatusPtrOutput struct{ *pulumi.OutputState }

func (ProjectCreateEnvironmentInputProtectedStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectCreateEnvironmentInputProtectedStatus)(nil)).Elem()
}

func (o ProjectCreateEnvironmentInputProtectedStatusPtrOutput) ToProjectCreateEnvironmentInputProtectedStatusPtrOutput() ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return o
}

func (o ProjectCreateEnvironmentInputProtectedStatusPtrOutput) ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return o
}

func (o ProjectCreateEnvironmentInputProtectedStatusPtrOutput) Elem() ProjectCreateEnvironmentInputProtectedStatusOutput {
	return o.ApplyT(func(v *ProjectCreateEnvironmentInputProtectedStatus) ProjectCreateEnvironmentInputProtectedStatus {
		if v != nil {
			return *v
		}
		var ret ProjectCreateEnvironmentInputProtectedStatus
		return ret
	}).(ProjectCreateEnvironmentInputProtectedStatusOutput)
}

func (o ProjectCreateEnvironmentInputProtectedStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProjectCreateEnvironmentInputProtectedStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProjectCreateEnvironmentInputProtectedStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ProjectCreateEnvironmentInputProtectedStatusInput is an input type that accepts values of the ProjectCreateEnvironmentInputProtectedStatus enum
// A concrete instance of `ProjectCreateEnvironmentInputProtectedStatusInput` can be one of the following:
//
//	ProjectCreateEnvironmentInputProtectedStatusUnprotected
//	ProjectCreateEnvironmentInputProtectedStatusProtected
type ProjectCreateEnvironmentInputProtectedStatusInput interface {
	pulumi.Input

	ToProjectCreateEnvironmentInputProtectedStatusOutput() ProjectCreateEnvironmentInputProtectedStatusOutput
	ToProjectCreateEnvironmentInputProtectedStatusOutputWithContext(context.Context) ProjectCreateEnvironmentInputProtectedStatusOutput
}

var projectCreateEnvironmentInputProtectedStatusPtrType = reflect.TypeOf((**ProjectCreateEnvironmentInputProtectedStatus)(nil)).Elem()

type ProjectCreateEnvironmentInputProtectedStatusPtrInput interface {
	pulumi.Input

	ToProjectCreateEnvironmentInputProtectedStatusPtrOutput() ProjectCreateEnvironmentInputProtectedStatusPtrOutput
	ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(context.Context) ProjectCreateEnvironmentInputProtectedStatusPtrOutput
}

type projectCreateEnvironmentInputProtectedStatusPtr string

func ProjectCreateEnvironmentInputProtectedStatusPtr(v string) ProjectCreateEnvironmentInputProtectedStatusPtrInput {
	return (*projectCreateEnvironmentInputProtectedStatusPtr)(&v)
}

func (*projectCreateEnvironmentInputProtectedStatusPtr) ElementType() reflect.Type {
	return projectCreateEnvironmentInputProtectedStatusPtrType
}

func (in *projectCreateEnvironmentInputProtectedStatusPtr) ToProjectCreateEnvironmentInputProtectedStatusPtrOutput() ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return pulumi.ToOutput(in).(ProjectCreateEnvironmentInputProtectedStatusPtrOutput)
}

func (in *projectCreateEnvironmentInputProtectedStatusPtr) ToProjectCreateEnvironmentInputProtectedStatusPtrOutputWithContext(ctx context.Context) ProjectCreateEnvironmentInputProtectedStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProjectCreateEnvironmentInputProtectedStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectCreateEnvironmentInputProtectedStatusInput)(nil)).Elem(), ProjectCreateEnvironmentInputProtectedStatus("unprotected"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectCreateEnvironmentInputProtectedStatusPtrInput)(nil)).Elem(), ProjectCreateEnvironmentInputProtectedStatus("unprotected"))
	pulumi.RegisterOutputType(OwnerTypeOutput{})
	pulumi.RegisterOutputType(OwnerTypePtrOutput{})
	pulumi.RegisterOutputType(ProjectCreateEnvironmentInputProtectedStatusOutput{})
	pulumi.RegisterOutputType(ProjectCreateEnvironmentInputProtectedStatusPtrOutput{})
}
