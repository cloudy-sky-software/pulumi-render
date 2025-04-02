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
from . import outputs
from ._enums import *

__all__ = ['EnvGroupEnvVarArgs', 'EnvGroupEnvVar']

@pulumi.input_type
class EnvGroupEnvVarArgs:
    def __init__(__self__, *,
                 env_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 env_var_key: Optional[pulumi.Input[builtins.str]] = None,
                 generate_value: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a EnvGroupEnvVar resource.
        :param pulumi.Input[builtins.str] env_group_id: Filter for resources that belong to an environment group
        :param pulumi.Input[builtins.str] env_var_key: The name of the environment variable
        """
        if env_group_id is not None:
            pulumi.set(__self__, "env_group_id", env_group_id)
        if env_var_key is not None:
            pulumi.set(__self__, "env_var_key", env_var_key)
        if generate_value is not None:
            pulumi.set(__self__, "generate_value", generate_value)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="envGroupId")
    def env_group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Filter for resources that belong to an environment group
        """
        return pulumi.get(self, "env_group_id")

    @env_group_id.setter
    def env_group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "env_group_id", value)

    @property
    @pulumi.getter(name="envVarKey")
    def env_var_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the environment variable
        """
        return pulumi.get(self, "env_var_key")

    @env_var_key.setter
    def env_var_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "env_var_key", value)

    @property
    @pulumi.getter(name="generateValue")
    def generate_value(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "generate_value")

    @generate_value.setter
    def generate_value(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "generate_value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "value", value)


class EnvGroupEnvVar(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 env_var_key: Optional[pulumi.Input[builtins.str]] = None,
                 generate_value: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a EnvGroupEnvVar resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] env_group_id: Filter for resources that belong to an environment group
        :param pulumi.Input[builtins.str] env_var_key: The name of the environment variable
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EnvGroupEnvVarArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EnvGroupEnvVar resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EnvGroupEnvVarArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvGroupEnvVarArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 env_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 env_var_key: Optional[pulumi.Input[builtins.str]] = None,
                 generate_value: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvGroupEnvVarArgs.__new__(EnvGroupEnvVarArgs)

            __props__.__dict__["env_group_id"] = env_group_id
            __props__.__dict__["env_var_key"] = env_var_key
            __props__.__dict__["generate_value"] = generate_value
            __props__.__dict__["value"] = value
            __props__.__dict__["created_at"] = None
            __props__.__dict__["env_vars"] = None
            __props__.__dict__["environment_id"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["owner_id"] = None
            __props__.__dict__["secret_files"] = None
            __props__.__dict__["service_links"] = None
            __props__.__dict__["updated_at"] = None
        super(EnvGroupEnvVar, __self__).__init__(
            'render:env-groups:EnvGroupEnvVar',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EnvGroupEnvVar':
        """
        Get an existing EnvGroupEnvVar resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvGroupEnvVarArgs.__new__(EnvGroupEnvVarArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["env_vars"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["generate_value"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["secret_files"] = None
        __props__.__dict__["service_links"] = None
        __props__.__dict__["updated_at"] = None
        __props__.__dict__["value"] = None
        return EnvGroupEnvVar(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="envVars")
    def env_vars(self) -> pulumi.Output[Optional[Sequence['outputs.EnvVar']]]:
        return pulumi.get(self, "env_vars")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="generateValue")
    def generate_value(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "generate_value")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> pulumi.Output[Optional[Sequence['outputs.SecretFile']]]:
        return pulumi.get(self, "secret_files")

    @property
    @pulumi.getter(name="serviceLinks")
    def service_links(self) -> pulumi.Output[Optional[Sequence['outputs.EnvGroupLink']]]:
        """
        List of serviceIds linked to the envGroup
        """
        return pulumi.get(self, "service_links")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "value")

