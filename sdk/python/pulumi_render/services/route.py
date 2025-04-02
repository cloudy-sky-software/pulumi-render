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
from ._enums import *

__all__ = ['RouteArgs', 'Route']

@pulumi.input_type
class RouteArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[builtins.str],
                 source: pulumi.Input[builtins.str],
                 type: pulumi.Input['Type'],
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Route resource.
        :param pulumi.Input[builtins.int] priority: Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
        :param pulumi.Input[builtins.str] service_id: The ID of the service
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "type", type)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['Type']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['Type']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the service
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_id", value)


class Route(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None,
                 source: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        """
        Create a Route resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] priority: Redirect and Rewrite Rules are applied in priority order starting at 0. Defaults to last in the priority list.
        :param pulumi.Input[builtins.str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Route resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None,
                 source: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input['Type']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteArgs.__new__(RouteArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            __props__.__dict__["priority"] = priority
            __props__.__dict__["service_id"] = service_id
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Route, __self__).__init__(
            'render:services:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RouteArgs.__new__(RouteArgs)

        __props__.__dict__["destination"] = None
        __props__.__dict__["priority"] = None
        __props__.__dict__["source"] = None
        __props__.__dict__["type"] = None
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[builtins.int]:
        """
        Redirect and Rewrite Rules are applied in priority order starting at 0
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['Type']:
        return pulumi.get(self, "type")

