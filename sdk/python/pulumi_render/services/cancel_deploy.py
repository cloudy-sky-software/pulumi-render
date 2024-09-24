# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import outputs
from ._enums import *

__all__ = ['CancelDeployArgs', 'CancelDeploy']

@pulumi.input_type
class CancelDeployArgs:
    def __init__(__self__, *,
                 deploy_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CancelDeploy resource.
        :param pulumi.Input[str] deploy_id: The ID of the deploy
        :param pulumi.Input[str] service_id: The ID of the service
        """
        if deploy_id is not None:
            pulumi.set(__self__, "deploy_id", deploy_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="deployId")
    def deploy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the deploy
        """
        return pulumi.get(self, "deploy_id")

    @deploy_id.setter
    def deploy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_id", value)

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


class CancelDeploy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CancelDeploy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deploy_id: The ID of the deploy
        :param pulumi.Input[str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CancelDeployArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CancelDeploy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CancelDeployArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CancelDeployArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CancelDeployArgs.__new__(CancelDeployArgs)

            __props__.__dict__["deploy_id"] = deploy_id
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["commit"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["finished_at"] = None
            __props__.__dict__["image"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["trigger"] = None
            __props__.__dict__["updated_at"] = None
        super(CancelDeploy, __self__).__init__(
            'render:services:CancelDeploy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CancelDeploy':
        """
        Get an existing CancelDeploy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CancelDeployArgs.__new__(CancelDeployArgs)

        __props__.__dict__["commit"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["finished_at"] = None
        __props__.__dict__["image"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["trigger"] = None
        __props__.__dict__["updated_at"] = None
        return CancelDeploy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def commit(self) -> pulumi.Output[Optional['outputs.CommitProperties']]:
        return pulumi.get(self, "commit")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="finishedAt")
    def finished_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "finished_at")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output[Optional['outputs.ImageProperties']]:
        """
        Image information used when creating the deploy. Not present for Git-backed deploys
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['Status']]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def trigger(self) -> pulumi.Output[Optional['Trigger']]:
        return pulumi.get(self, "trigger")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "updated_at")

