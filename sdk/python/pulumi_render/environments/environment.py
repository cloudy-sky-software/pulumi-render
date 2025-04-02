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

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_isolation_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 protected_status: Optional[pulumi.Input['ProtectedStatus']] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[builtins.bool] network_isolation_enabled: Indicates whether network connections across environments are allowed.
        :param pulumi.Input['ProtectedStatus'] protected_status: Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        pulumi.set(__self__, "project_id", project_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_isolation_enabled is not None:
            pulumi.set(__self__, "network_isolation_enabled", network_isolation_enabled)
        if protected_status is not None:
            pulumi.set(__self__, "protected_status", protected_status)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkIsolationEnabled")
    def network_isolation_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether network connections across environments are allowed.
        """
        return pulumi.get(self, "network_isolation_enabled")

    @network_isolation_enabled.setter
    def network_isolation_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "network_isolation_enabled", value)

    @property
    @pulumi.getter(name="protectedStatus")
    def protected_status(self) -> Optional[pulumi.Input['ProtectedStatus']]:
        """
        Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        return pulumi.get(self, "protected_status")

    @protected_status.setter
    def protected_status(self, value: Optional[pulumi.Input['ProtectedStatus']]):
        pulumi.set(self, "protected_status", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_isolation_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 protected_status: Optional[pulumi.Input['ProtectedStatus']] = None,
                 __props__=None):
        """
        Create a Environment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] network_isolation_enabled: Indicates whether network connections across environments are allowed.
        :param pulumi.Input['ProtectedStatus'] protected_status: Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Environment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_isolation_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 protected_status: Optional[pulumi.Input['ProtectedStatus']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["network_isolation_enabled"] = network_isolation_enabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["protected_status"] = protected_status
            __props__.__dict__["databases_ids"] = None
            __props__.__dict__["env_group_ids"] = None
            __props__.__dict__["redis_ids"] = None
            __props__.__dict__["service_ids"] = None
        super(Environment, __self__).__init__(
            'render:environments:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

        __props__.__dict__["databases_ids"] = None
        __props__.__dict__["env_group_ids"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_isolation_enabled"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["protected_status"] = None
        __props__.__dict__["redis_ids"] = None
        __props__.__dict__["service_ids"] = None
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databasesIds")
    def databases_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "databases_ids")

    @property
    @pulumi.getter(name="envGroupIds")
    def env_group_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "env_group_ids")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkIsolationEnabled")
    def network_isolation_enabled(self) -> pulumi.Output[builtins.bool]:
        """
        Indicates whether network connections across environments are allowed.
        """
        return pulumi.get(self, "network_isolation_enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="protectedStatus")
    def protected_status(self) -> pulumi.Output['ProtectedStatus']:
        """
        Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        return pulumi.get(self, "protected_status")

    @property
    @pulumi.getter(name="redisIds")
    def redis_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "redis_ids")

    @property
    @pulumi.getter(name="serviceIds")
    def service_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "service_ids")

