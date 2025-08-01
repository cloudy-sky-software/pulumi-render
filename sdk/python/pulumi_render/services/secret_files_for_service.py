# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
from ._inputs import *

__all__ = ['SecretFilesForServiceArgs', 'SecretFilesForService']

@pulumi.input_type
class SecretFilesForServiceArgs:
    def __init__(__self__, *,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]] = None,
                 service_id: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a SecretFilesForService resource.
        :param pulumi.Input[_builtins.str] service_id: The ID of the service
        """
        if secret_files is not None:
            pulumi.set(__self__, "secret_files", secret_files)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @_builtins.property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]]:
        return pulumi.get(self, "secret_files")

    @secret_files.setter
    def secret_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]]):
        pulumi.set(self, "secret_files", value)

    @_builtins.property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the service
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_id", value)


@pulumi.type_token("render:services:SecretFilesForService")
class SecretFilesForService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretFileInputArgs', 'SecretFileInputArgsDict']]]]] = None,
                 service_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Create a SecretFilesForService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] service_id: The ID of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretFilesForServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecretFilesForService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecretFilesForServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretFilesForServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretFileInputArgs', 'SecretFileInputArgsDict']]]]] = None,
                 service_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretFilesForServiceArgs.__new__(SecretFilesForServiceArgs)

            __props__.__dict__["secret_files"] = secret_files
            __props__.__dict__["service_id"] = service_id
        super(SecretFilesForService, __self__).__init__(
            'render:services:SecretFilesForService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecretFilesForService':
        """
        Get an existing SecretFilesForService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecretFilesForServiceArgs.__new__(SecretFilesForServiceArgs)

        __props__.__dict__["secret_files"] = None
        return SecretFilesForService(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> pulumi.Output[Optional[Sequence['outputs.SecretFileInput']]]:
        return pulumi.get(self, "secret_files")

