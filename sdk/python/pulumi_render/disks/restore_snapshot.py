# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['RestoreSnapshotArgs', 'RestoreSnapshot']

@pulumi.input_type
class RestoreSnapshotArgs:
    def __init__(__self__, *,
                 snapshot_key: pulumi.Input[builtins.str],
                 disk_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RestoreSnapshot resource.
        :param pulumi.Input[builtins.str] disk_id: The ID of the disk
        :param pulumi.Input[builtins.str] instance_id: When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
        """
        pulumi.set(__self__, "snapshot_key", snapshot_key)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="snapshotKey")
    def snapshot_key(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "snapshot_key")

    @snapshot_key.setter
    def snapshot_key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "snapshot_key", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the disk
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)


class RestoreSnapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 snapshot_key: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a RestoreSnapshot resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] disk_id: The ID of the disk
        :param pulumi.Input[builtins.str] instance_id: When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RestoreSnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RestoreSnapshot resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RestoreSnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RestoreSnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 snapshot_key: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RestoreSnapshotArgs.__new__(RestoreSnapshotArgs)

            __props__.__dict__["disk_id"] = disk_id
            __props__.__dict__["instance_id"] = instance_id
            if snapshot_key is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_key'")
            __props__.__dict__["snapshot_key"] = snapshot_key
            __props__.__dict__["created_at"] = None
            __props__.__dict__["mount_path"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["service_id"] = None
            __props__.__dict__["size_gb"] = None
            __props__.__dict__["updated_at"] = None
        super(RestoreSnapshot, __self__).__init__(
            'render:disks:RestoreSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RestoreSnapshot':
        """
        Get an existing RestoreSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RestoreSnapshotArgs.__new__(RestoreSnapshotArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["instance_id"] = None
        __props__.__dict__["mount_path"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["service_id"] = None
        __props__.__dict__["size_gb"] = None
        __props__.__dict__["snapshot_key"] = None
        __props__.__dict__["updated_at"] = None
        return RestoreSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        When a service with a disk is scaled, the instanceId is used to identify the instance that the disk is attached to. Each instance's disks get their own snapshots, and can be restored separately.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="mountPath")
    def mount_path(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "mount_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="sizeGB")
    def size_gb(self) -> pulumi.Output[builtins.int]:
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter(name="snapshotKey")
    def snapshot_key(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "snapshot_key")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "updated_at")

