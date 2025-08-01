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

__all__ = [
    'GetEnvVarResult',
    'AwaitableGetEnvVarResult',
    'get_env_var',
    'get_env_var_output',
]

@pulumi.output_type
class GetEnvVarResult:
    def __init__(__self__, key=None, value=None):
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @_builtins.property
    @pulumi.getter
    def key(self) -> _builtins.str:
        return pulumi.get(self, "key")

    @_builtins.property
    @pulumi.getter
    def value(self) -> _builtins.str:
        return pulumi.get(self, "value")


class AwaitableGetEnvVarResult(GetEnvVarResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvVarResult(
            key=self.key,
            value=self.value)


def get_env_var(env_var_key: Optional[_builtins.str] = None,
                service_id: Optional[_builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvVarResult:
    """
    Use this data source to access information about an existing resource.

    :param _builtins.str env_var_key: The name of the environment variable
    :param _builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['envVarKey'] = env_var_key
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:getEnvVar', __args__, opts=opts, typ=GetEnvVarResult).value

    return AwaitableGetEnvVarResult(
        key=pulumi.get(__ret__, 'key'),
        value=pulumi.get(__ret__, 'value'))
def get_env_var_output(env_var_key: Optional[pulumi.Input[_builtins.str]] = None,
                       service_id: Optional[pulumi.Input[_builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEnvVarResult]:
    """
    Use this data source to access information about an existing resource.

    :param _builtins.str env_var_key: The name of the environment variable
    :param _builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['envVarKey'] = env_var_key
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:services:getEnvVar', __args__, opts=opts, typ=GetEnvVarResult)
    return __ret__.apply(lambda __response__: GetEnvVarResult(
        key=pulumi.get(__response__, 'key'),
        value=pulumi.get(__response__, 'value')))
