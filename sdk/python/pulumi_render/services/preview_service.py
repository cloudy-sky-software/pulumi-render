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

__all__ = ['PreviewServiceArgs', 'PreviewService']

@pulumi.input_type
class PreviewServiceArgs:
    def __init__(__self__, *,
                 image_path: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['Plan']] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PreviewService resource.
        :param pulumi.Input[str] image_path: Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
        :param pulumi.Input[str] name: A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        :param pulumi.Input['Plan'] plan: The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        :param pulumi.Input[str] service_id: The ID of the service
        """
        pulumi.set(__self__, "image_path", image_path)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="imagePath")
    def image_path(self) -> pulumi.Input[str]:
        """
        Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
        """
        return pulumi.get(self, "image_path")

    @image_path.setter
    def image_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['Plan']]:
        """
        The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['Plan']]):
        pulumi.set(self, "plan", value)

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


class PreviewService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['Plan']] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PreviewService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_path: Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
        :param pulumi.Input[str] name: A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        :param pulumi.Input['Plan'] plan: The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        :param pulumi.Input[str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PreviewServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PreviewService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PreviewServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PreviewServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['Plan']] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PreviewServiceArgs.__new__(PreviewServiceArgs)

            if image_path is None and not opts.urn:
                raise TypeError("Missing required property 'image_path'")
            __props__.__dict__["image_path"] = image_path
            __props__.__dict__["name"] = name
            __props__.__dict__["plan"] = plan
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["deploy_id"] = None
            __props__.__dict__["service"] = None
        super(PreviewService, __self__).__init__(
            'render:services:PreviewService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PreviewService':
        """
        Get an existing PreviewService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PreviewServiceArgs.__new__(PreviewServiceArgs)

        __props__.__dict__["deploy_id"] = None
        __props__.__dict__["image_path"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["plan"] = None
        __props__.__dict__["service"] = None
        return PreviewService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deployId")
    def deploy_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "deploy_id")

    @property
    @pulumi.getter(name="imagePath")
    def image_path(self) -> pulumi.Output[str]:
        """
        Must be either a full URL or the relative path to an image. If a relative path, Render uses the base service's image URL as its root. For example, if the base service's image URL is `docker.io/library/nginx:latest`, then valid values are: `docker.io/library/nginx:<any tag or SHA>`, `library/nginx:<any tag or SHA>`, or `nginx:<any tag or SHA>`. Note that the path must match (only the tag or SHA can vary).
        """
        return pulumi.get(self, "image_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        A name for the service preview instance. If not specified, Render generates the name using the base service's name and the specified tag or SHA.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[Optional['Plan']]:
        """
        The instance type to use for the preview instance. Note that base services with any paid instance type can't create preview instances with the `free` instance type.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[Optional['outputs.Service']]:
        return pulumi.get(self, "service")
