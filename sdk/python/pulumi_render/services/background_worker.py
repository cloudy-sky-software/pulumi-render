# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['BackgroundWorkerArgs', 'BackgroundWorker']

@pulumi.input_type
class BackgroundWorkerArgs:
    def __init__(__self__, *,
                 service_details: Optional[pulumi.Input['BackgroundWorkerDetailsCreateArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BackgroundWorker resource.
        """
        if service_details is not None:
            pulumi.set(__self__, "service_details", service_details)
        if type is None:
            type = 'background_worker'
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> Optional[pulumi.Input['BackgroundWorkerDetailsCreateArgs']]:
        return pulumi.get(self, "service_details")

    @service_details.setter
    def service_details(self, value: Optional[pulumi.Input['BackgroundWorkerDetailsCreateArgs']]):
        pulumi.set(self, "service_details", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class BackgroundWorker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_details: Optional[pulumi.Input[pulumi.InputType['BackgroundWorkerDetailsCreateArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BackgroundWorker resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BackgroundWorkerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BackgroundWorker resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BackgroundWorkerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackgroundWorkerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_details: Optional[pulumi.Input[pulumi.InputType['BackgroundWorkerDetailsCreateArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackgroundWorkerArgs.__new__(BackgroundWorkerArgs)

            __props__.__dict__["service_details"] = service_details
            if type is None:
                type = 'background_worker'
            __props__.__dict__["type"] = type
            __props__.__dict__["auto_deploy"] = None
            __props__.__dict__["branch"] = None
            __props__.__dict__["build_filter"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["image_path"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["notify_on_fail"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["repo"] = None
            __props__.__dict__["root_dir"] = None
            __props__.__dict__["slug"] = None
            __props__.__dict__["suspended"] = None
            __props__.__dict__["suspenders"] = None
            __props__.__dict__["updated_at"] = None
        super(BackgroundWorker, __self__).__init__(
            'render:services:BackgroundWorker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackgroundWorker':
        """
        Get an existing BackgroundWorker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BackgroundWorkerArgs.__new__(BackgroundWorkerArgs)

        __props__.__dict__["auto_deploy"] = None
        __props__.__dict__["branch"] = None
        __props__.__dict__["build_filter"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["image_path"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notify_on_fail"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["repo"] = None
        __props__.__dict__["root_dir"] = None
        __props__.__dict__["service_details"] = None
        __props__.__dict__["slug"] = None
        __props__.__dict__["suspended"] = None
        __props__.__dict__["suspenders"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return BackgroundWorker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> pulumi.Output['BackgroundWorkerAutoDeploy']:
        return pulumi.get(self, "auto_deploy")

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="buildFilter")
    def build_filter(self) -> pulumi.Output[Optional['outputs.BuildFilter']]:
        return pulumi.get(self, "build_filter")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="imagePath")
    def image_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyOnFail")
    def notify_on_fail(self) -> pulumi.Output['BackgroundWorkerNotifyOnFail']:
        return pulumi.get(self, "notify_on_fail")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter(name="rootDir")
    def root_dir(self) -> pulumi.Output[str]:
        return pulumi.get(self, "root_dir")

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> pulumi.Output['outputs.BackgroundWorkerDetailsCreate']:
        return pulumi.get(self, "service_details")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def suspended(self) -> pulumi.Output['BackgroundWorkerSuspended']:
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def suspenders(self) -> pulumi.Output[Sequence['BackgroundWorkerSuspendersItem']]:
        return pulumi.get(self, "suspenders")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

