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
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[_builtins.str] api_key: The Render API key.
        """
        if api_key is None:
            api_key = _utilities.get_env('RENDER_APIKEY')
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)

    @_builtins.property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Render API key.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "api_key", value)


@pulumi.type_token("pulumi:providers:render")
class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        The provider type for the render package.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] api_key: The Render API key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the render package.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if api_key is None:
                api_key = _utilities.get_env('RENDER_APIKEY')
            __props__.__dict__["api_key"] = None if api_key is None else pulumi.Output.secret(api_key)
        super(Provider, __self__).__init__(
            'render',
            resource_name,
            __props__,
            opts)

