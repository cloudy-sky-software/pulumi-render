// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package envgroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceLinkType string

const (
	ServiceLinkTypeStatic = ServiceLinkType("static")
	ServiceLinkTypeWeb    = ServiceLinkType("web")
	ServiceLinkTypePserv  = ServiceLinkType("pserv")
	ServiceLinkTypeWorker = ServiceLinkType("worker")
	ServiceLinkTypeCron   = ServiceLinkType("cron")
)

type ServiceLinkTypeOutput struct{ *pulumi.OutputState }

func (ServiceLinkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLinkType)(nil)).Elem()
}

func (o ServiceLinkTypeOutput) ToServiceLinkTypeOutput() ServiceLinkTypeOutput {
	return o
}

func (o ServiceLinkTypeOutput) ToServiceLinkTypeOutputWithContext(ctx context.Context) ServiceLinkTypeOutput {
	return o
}

func (o ServiceLinkTypeOutput) ToServiceLinkTypePtrOutput() ServiceLinkTypePtrOutput {
	return o.ToServiceLinkTypePtrOutputWithContext(context.Background())
}

func (o ServiceLinkTypeOutput) ToServiceLinkTypePtrOutputWithContext(ctx context.Context) ServiceLinkTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceLinkType) *ServiceLinkType {
		return &v
	}).(ServiceLinkTypePtrOutput)
}

func (o ServiceLinkTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceLinkTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLinkType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceLinkTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLinkTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLinkType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceLinkTypePtrOutput struct{ *pulumi.OutputState }

func (ServiceLinkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLinkType)(nil)).Elem()
}

func (o ServiceLinkTypePtrOutput) ToServiceLinkTypePtrOutput() ServiceLinkTypePtrOutput {
	return o
}

func (o ServiceLinkTypePtrOutput) ToServiceLinkTypePtrOutputWithContext(ctx context.Context) ServiceLinkTypePtrOutput {
	return o
}

func (o ServiceLinkTypePtrOutput) Elem() ServiceLinkTypeOutput {
	return o.ApplyT(func(v *ServiceLinkType) ServiceLinkType {
		if v != nil {
			return *v
		}
		var ret ServiceLinkType
		return ret
	}).(ServiceLinkTypeOutput)
}

func (o ServiceLinkTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLinkTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceLinkType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceLinkTypeOutput{})
	pulumi.RegisterOutputType(ServiceLinkTypePtrOutput{})
}
