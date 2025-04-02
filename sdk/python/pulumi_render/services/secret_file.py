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

__all__ = ['SecretFileArgs', 'SecretFile']

@pulumi.input_type
class SecretFileArgs:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[builtins.str]] = None,
                 secret_file_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SecretFile resource.
        :param pulumi.Input[builtins.str] secret_file_name: The file name of the secret file
        :param pulumi.Input[builtins.str] service_id: The ID of the service
        """
        if content is not None:
            pulumi.set(__self__, "content", content)
        if secret_file_name is not None:
            pulumi.set(__self__, "secret_file_name", secret_file_name)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="secretFileName")
    def secret_file_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The file name of the secret file
        """
        return pulumi.get(self, "secret_file_name")

    @secret_file_name.setter
    def secret_file_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_file_name", value)

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


class SecretFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[builtins.str]] = None,
                 secret_file_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create a SecretFile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] secret_file_name: The file name of the secret file
        :param pulumi.Input[builtins.str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretFileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecretFile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecretFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[builtins.str]] = None,
                 secret_file_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretFileArgs.__new__(SecretFileArgs)

            __props__.__dict__["content"] = content
            __props__.__dict__["secret_file_name"] = secret_file_name
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["name"] = None
        super(SecretFile, __self__).__init__(
            'render:services:SecretFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecretFile':
        """
        Get an existing SecretFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecretFileArgs.__new__(SecretFileArgs)

        __props__.__dict__["content"] = None
        __props__.__dict__["name"] = None
        return SecretFile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

