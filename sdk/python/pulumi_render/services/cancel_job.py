# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CancelJobArgs', 'CancelJob']

@pulumi.input_type
class CancelJobArgs:
    def __init__(__self__, *,
                 job_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CancelJob resource.
        :param pulumi.Input[str] job_id: The ID of the job
        :param pulumi.Input[str] service_id: The ID of the service
        """
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the job
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the service
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class CancelJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CancelJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_id: The ID of the job
        :param pulumi.Input[str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CancelJobArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CancelJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CancelJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CancelJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CancelJobArgs.__new__(CancelJobArgs)

            __props__.__dict__["job_id"] = job_id
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["finished_at"] = None
            __props__.__dict__["plan_id"] = None
            __props__.__dict__["start_command"] = None
            __props__.__dict__["started_at"] = None
            __props__.__dict__["status"] = None
        super(CancelJob, __self__).__init__(
            'render:services:CancelJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CancelJob':
        """
        Get an existing CancelJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CancelJobArgs.__new__(CancelJobArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["finished_at"] = None
        __props__.__dict__["plan_id"] = None
        __props__.__dict__["service_id"] = None
        __props__.__dict__["start_command"] = None
        __props__.__dict__["started_at"] = None
        __props__.__dict__["status"] = None
        return CancelJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="finishedAt")
    def finished_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "finished_at")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="startCommand")
    def start_command(self) -> pulumi.Output[str]:
        return pulumi.get(self, "start_command")

    @property
    @pulumi.getter(name="startedAt")
    def started_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "started_at")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "status")

