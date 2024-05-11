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
from ._inputs import *

__all__ = ['AutoscaleServiceArgs', 'AutoscaleService']

@pulumi.input_type
class AutoscaleServiceArgs:
    def __init__(__self__, *,
                 criteria: pulumi.Input['AutoscalingCriteriaArgs'],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max: pulumi.Input[int],
                 min: pulumi.Input[int],
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AutoscaleService resource.
        :param pulumi.Input[int] max: The maximum number of instances for the service
        :param pulumi.Input[int] min: The minimum number of instances for the service
        :param pulumi.Input[str] service_id: The ID of the service
        """
        pulumi.set(__self__, "criteria", criteria)
        if enabled is None:
            enabled = False
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "max", max)
        pulumi.set(__self__, "min", min)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Input['AutoscalingCriteriaArgs']:
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: pulumi.Input['AutoscalingCriteriaArgs']):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def max(self) -> pulumi.Input[int]:
        """
        The maximum number of instances for the service
        """
        return pulumi.get(self, "max")

    @max.setter
    def max(self, value: pulumi.Input[int]):
        pulumi.set(self, "max", value)

    @property
    @pulumi.getter
    def min(self) -> pulumi.Input[int]:
        """
        The minimum number of instances for the service
        """
        return pulumi.get(self, "min")

    @min.setter
    def min(self, value: pulumi.Input[int]):
        pulumi.set(self, "min", value)

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


class AutoscaleService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['AutoscalingCriteriaArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max: Optional[pulumi.Input[int]] = None,
                 min: Optional[pulumi.Input[int]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AutoscaleService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max: The maximum number of instances for the service
        :param pulumi.Input[int] min: The minimum number of instances for the service
        :param pulumi.Input[str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoscaleServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AutoscaleService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AutoscaleServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoscaleServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['AutoscalingCriteriaArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max: Optional[pulumi.Input[int]] = None,
                 min: Optional[pulumi.Input[int]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoscaleServiceArgs.__new__(AutoscaleServiceArgs)

            if criteria is None and not opts.urn:
                raise TypeError("Missing required property 'criteria'")
            __props__.__dict__["criteria"] = criteria
            if enabled is None:
                enabled = False
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if max is None and not opts.urn:
                raise TypeError("Missing required property 'max'")
            __props__.__dict__["max"] = max
            if min is None and not opts.urn:
                raise TypeError("Missing required property 'min'")
            __props__.__dict__["min"] = min
            __props__.__dict__["service_id"] = service_id
        super(AutoscaleService, __self__).__init__(
            'render:services:AutoscaleService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AutoscaleService':
        """
        Get an existing AutoscaleService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AutoscaleServiceArgs.__new__(AutoscaleServiceArgs)

        __props__.__dict__["criteria"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["max"] = None
        __props__.__dict__["min"] = None
        return AutoscaleService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output['outputs.AutoscalingCriteria']:
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def max(self) -> pulumi.Output[int]:
        """
        The maximum number of instances for the service
        """
        return pulumi.get(self, "max")

    @property
    @pulumi.getter
    def min(self) -> pulumi.Output[int]:
        """
        The minimum number of instances for the service
        """
        return pulumi.get(self, "min")

