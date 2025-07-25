// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprints

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlueprintWithCursorBlueprintPropertiesStatus string

const (
	BlueprintWithCursorBlueprintPropertiesStatusCreated = BlueprintWithCursorBlueprintPropertiesStatus("created")
	BlueprintWithCursorBlueprintPropertiesStatusPaused  = BlueprintWithCursorBlueprintPropertiesStatus("paused")
	BlueprintWithCursorBlueprintPropertiesStatusInSync  = BlueprintWithCursorBlueprintPropertiesStatus("in_sync")
	BlueprintWithCursorBlueprintPropertiesStatusSyncing = BlueprintWithCursorBlueprintPropertiesStatus("syncing")
	BlueprintWithCursorBlueprintPropertiesStatusError   = BlueprintWithCursorBlueprintPropertiesStatus("error")
)

type BlueprintWithCursorBlueprintPropertiesStatusOutput struct{ *pulumi.OutputState }

func (BlueprintWithCursorBlueprintPropertiesStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlueprintWithCursorBlueprintPropertiesStatus)(nil)).Elem()
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToBlueprintWithCursorBlueprintPropertiesStatusOutput() BlueprintWithCursorBlueprintPropertiesStatusOutput {
	return o
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToBlueprintWithCursorBlueprintPropertiesStatusOutputWithContext(ctx context.Context) BlueprintWithCursorBlueprintPropertiesStatusOutput {
	return o
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToBlueprintWithCursorBlueprintPropertiesStatusPtrOutput() BlueprintWithCursorBlueprintPropertiesStatusPtrOutput {
	return o.ToBlueprintWithCursorBlueprintPropertiesStatusPtrOutputWithContext(context.Background())
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToBlueprintWithCursorBlueprintPropertiesStatusPtrOutputWithContext(ctx context.Context) BlueprintWithCursorBlueprintPropertiesStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlueprintWithCursorBlueprintPropertiesStatus) *BlueprintWithCursorBlueprintPropertiesStatus {
		return &v
	}).(BlueprintWithCursorBlueprintPropertiesStatusPtrOutput)
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlueprintWithCursorBlueprintPropertiesStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlueprintWithCursorBlueprintPropertiesStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlueprintWithCursorBlueprintPropertiesStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlueprintWithCursorBlueprintPropertiesStatusPtrOutput struct{ *pulumi.OutputState }

func (BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlueprintWithCursorBlueprintPropertiesStatus)(nil)).Elem()
}

func (o BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) ToBlueprintWithCursorBlueprintPropertiesStatusPtrOutput() BlueprintWithCursorBlueprintPropertiesStatusPtrOutput {
	return o
}

func (o BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) ToBlueprintWithCursorBlueprintPropertiesStatusPtrOutputWithContext(ctx context.Context) BlueprintWithCursorBlueprintPropertiesStatusPtrOutput {
	return o
}

func (o BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) Elem() BlueprintWithCursorBlueprintPropertiesStatusOutput {
	return o.ApplyT(func(v *BlueprintWithCursorBlueprintPropertiesStatus) BlueprintWithCursorBlueprintPropertiesStatus {
		if v != nil {
			return *v
		}
		var ret BlueprintWithCursorBlueprintPropertiesStatus
		return ret
	}).(BlueprintWithCursorBlueprintPropertiesStatusOutput)
}

func (o BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlueprintWithCursorBlueprintPropertiesStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlueprintWithCursorBlueprintPropertiesStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// type of the resource (ex. web_service or postgres)
type GetBlueprintPropertiesResourcesItemPropertiesType string

const (
	GetBlueprintPropertiesResourcesItemPropertiesTypeStaticSite       = GetBlueprintPropertiesResourcesItemPropertiesType("static_site")
	GetBlueprintPropertiesResourcesItemPropertiesTypeWebService       = GetBlueprintPropertiesResourcesItemPropertiesType("web_service")
	GetBlueprintPropertiesResourcesItemPropertiesTypePrivateService   = GetBlueprintPropertiesResourcesItemPropertiesType("private_service")
	GetBlueprintPropertiesResourcesItemPropertiesTypeBackgroundWorker = GetBlueprintPropertiesResourcesItemPropertiesType("background_worker")
	GetBlueprintPropertiesResourcesItemPropertiesTypeCronJob          = GetBlueprintPropertiesResourcesItemPropertiesType("cron_job")
	GetBlueprintPropertiesResourcesItemPropertiesTypeRedis            = GetBlueprintPropertiesResourcesItemPropertiesType("redis")
	GetBlueprintPropertiesResourcesItemPropertiesTypeKeyValue         = GetBlueprintPropertiesResourcesItemPropertiesType("key_value")
	GetBlueprintPropertiesResourcesItemPropertiesTypePostgres         = GetBlueprintPropertiesResourcesItemPropertiesType("postgres")
	GetBlueprintPropertiesResourcesItemPropertiesTypeEnvironmentGroup = GetBlueprintPropertiesResourcesItemPropertiesType("environment_group")
)

type GetBlueprintPropertiesResourcesItemPropertiesTypeOutput struct{ *pulumi.OutputState }

func (GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBlueprintPropertiesResourcesItemPropertiesType)(nil)).Elem()
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypeOutput() GetBlueprintPropertiesResourcesItemPropertiesTypeOutput {
	return o
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypeOutputWithContext(ctx context.Context) GetBlueprintPropertiesResourcesItemPropertiesTypeOutput {
	return o
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput() GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput {
	return o.ToGetBlueprintPropertiesResourcesItemPropertiesTypePtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypePtrOutputWithContext(ctx context.Context) GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GetBlueprintPropertiesResourcesItemPropertiesType) *GetBlueprintPropertiesResourcesItemPropertiesType {
		return &v
	}).(GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput)
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GetBlueprintPropertiesResourcesItemPropertiesType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GetBlueprintPropertiesResourcesItemPropertiesType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput struct{ *pulumi.OutputState }

func (GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetBlueprintPropertiesResourcesItemPropertiesType)(nil)).Elem()
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput() GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput {
	return o
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) ToGetBlueprintPropertiesResourcesItemPropertiesTypePtrOutputWithContext(ctx context.Context) GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput {
	return o
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) Elem() GetBlueprintPropertiesResourcesItemPropertiesTypeOutput {
	return o.ApplyT(func(v *GetBlueprintPropertiesResourcesItemPropertiesType) GetBlueprintPropertiesResourcesItemPropertiesType {
		if v != nil {
			return *v
		}
		var ret GetBlueprintPropertiesResourcesItemPropertiesType
		return ret
	}).(GetBlueprintPropertiesResourcesItemPropertiesTypeOutput)
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GetBlueprintPropertiesResourcesItemPropertiesType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GetBlueprintPropertiesStatus string

const (
	GetBlueprintPropertiesStatusCreated = GetBlueprintPropertiesStatus("created")
	GetBlueprintPropertiesStatusPaused  = GetBlueprintPropertiesStatus("paused")
	GetBlueprintPropertiesStatusInSync  = GetBlueprintPropertiesStatus("in_sync")
	GetBlueprintPropertiesStatusSyncing = GetBlueprintPropertiesStatus("syncing")
	GetBlueprintPropertiesStatusError   = GetBlueprintPropertiesStatus("error")
)

type GetBlueprintPropertiesStatusOutput struct{ *pulumi.OutputState }

func (GetBlueprintPropertiesStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBlueprintPropertiesStatus)(nil)).Elem()
}

func (o GetBlueprintPropertiesStatusOutput) ToGetBlueprintPropertiesStatusOutput() GetBlueprintPropertiesStatusOutput {
	return o
}

func (o GetBlueprintPropertiesStatusOutput) ToGetBlueprintPropertiesStatusOutputWithContext(ctx context.Context) GetBlueprintPropertiesStatusOutput {
	return o
}

func (o GetBlueprintPropertiesStatusOutput) ToGetBlueprintPropertiesStatusPtrOutput() GetBlueprintPropertiesStatusPtrOutput {
	return o.ToGetBlueprintPropertiesStatusPtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesStatusOutput) ToGetBlueprintPropertiesStatusPtrOutputWithContext(ctx context.Context) GetBlueprintPropertiesStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GetBlueprintPropertiesStatus) *GetBlueprintPropertiesStatus {
		return &v
	}).(GetBlueprintPropertiesStatusPtrOutput)
}

func (o GetBlueprintPropertiesStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GetBlueprintPropertiesStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GetBlueprintPropertiesStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GetBlueprintPropertiesStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GetBlueprintPropertiesStatusPtrOutput struct{ *pulumi.OutputState }

func (GetBlueprintPropertiesStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetBlueprintPropertiesStatus)(nil)).Elem()
}

func (o GetBlueprintPropertiesStatusPtrOutput) ToGetBlueprintPropertiesStatusPtrOutput() GetBlueprintPropertiesStatusPtrOutput {
	return o
}

func (o GetBlueprintPropertiesStatusPtrOutput) ToGetBlueprintPropertiesStatusPtrOutputWithContext(ctx context.Context) GetBlueprintPropertiesStatusPtrOutput {
	return o
}

func (o GetBlueprintPropertiesStatusPtrOutput) Elem() GetBlueprintPropertiesStatusOutput {
	return o.ApplyT(func(v *GetBlueprintPropertiesStatus) GetBlueprintPropertiesStatus {
		if v != nil {
			return *v
		}
		var ret GetBlueprintPropertiesStatus
		return ret
	}).(GetBlueprintPropertiesStatusOutput)
}

func (o GetBlueprintPropertiesStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GetBlueprintPropertiesStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GetBlueprintPropertiesStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncWithCursorSyncPropertiesState string

const (
	SyncWithCursorSyncPropertiesStateCreated = SyncWithCursorSyncPropertiesState("created")
	SyncWithCursorSyncPropertiesStatePending = SyncWithCursorSyncPropertiesState("pending")
	SyncWithCursorSyncPropertiesStateRunning = SyncWithCursorSyncPropertiesState("running")
	SyncWithCursorSyncPropertiesStateError   = SyncWithCursorSyncPropertiesState("error")
	SyncWithCursorSyncPropertiesStateSuccess = SyncWithCursorSyncPropertiesState("success")
)

type SyncWithCursorSyncPropertiesStateOutput struct{ *pulumi.OutputState }

func (SyncWithCursorSyncPropertiesStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncWithCursorSyncPropertiesState)(nil)).Elem()
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToSyncWithCursorSyncPropertiesStateOutput() SyncWithCursorSyncPropertiesStateOutput {
	return o
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToSyncWithCursorSyncPropertiesStateOutputWithContext(ctx context.Context) SyncWithCursorSyncPropertiesStateOutput {
	return o
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToSyncWithCursorSyncPropertiesStatePtrOutput() SyncWithCursorSyncPropertiesStatePtrOutput {
	return o.ToSyncWithCursorSyncPropertiesStatePtrOutputWithContext(context.Background())
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToSyncWithCursorSyncPropertiesStatePtrOutputWithContext(ctx context.Context) SyncWithCursorSyncPropertiesStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncWithCursorSyncPropertiesState) *SyncWithCursorSyncPropertiesState {
		return &v
	}).(SyncWithCursorSyncPropertiesStatePtrOutput)
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncWithCursorSyncPropertiesState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncWithCursorSyncPropertiesStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncWithCursorSyncPropertiesState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncWithCursorSyncPropertiesStatePtrOutput struct{ *pulumi.OutputState }

func (SyncWithCursorSyncPropertiesStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncWithCursorSyncPropertiesState)(nil)).Elem()
}

func (o SyncWithCursorSyncPropertiesStatePtrOutput) ToSyncWithCursorSyncPropertiesStatePtrOutput() SyncWithCursorSyncPropertiesStatePtrOutput {
	return o
}

func (o SyncWithCursorSyncPropertiesStatePtrOutput) ToSyncWithCursorSyncPropertiesStatePtrOutputWithContext(ctx context.Context) SyncWithCursorSyncPropertiesStatePtrOutput {
	return o
}

func (o SyncWithCursorSyncPropertiesStatePtrOutput) Elem() SyncWithCursorSyncPropertiesStateOutput {
	return o.ApplyT(func(v *SyncWithCursorSyncPropertiesState) SyncWithCursorSyncPropertiesState {
		if v != nil {
			return *v
		}
		var ret SyncWithCursorSyncPropertiesState
		return ret
	}).(SyncWithCursorSyncPropertiesStateOutput)
}

func (o SyncWithCursorSyncPropertiesStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncWithCursorSyncPropertiesStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncWithCursorSyncPropertiesState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BlueprintWithCursorBlueprintPropertiesStatusOutput{})
	pulumi.RegisterOutputType(BlueprintWithCursorBlueprintPropertiesStatusPtrOutput{})
	pulumi.RegisterOutputType(GetBlueprintPropertiesResourcesItemPropertiesTypeOutput{})
	pulumi.RegisterOutputType(GetBlueprintPropertiesResourcesItemPropertiesTypePtrOutput{})
	pulumi.RegisterOutputType(GetBlueprintPropertiesStatusOutput{})
	pulumi.RegisterOutputType(GetBlueprintPropertiesStatusPtrOutput{})
	pulumi.RegisterOutputType(SyncWithCursorSyncPropertiesStateOutput{})
	pulumi.RegisterOutputType(SyncWithCursorSyncPropertiesStatePtrOutput{})
}
