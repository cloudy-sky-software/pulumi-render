# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from ._enums import *

__all__ = [
    'GetResourceLogStreamProperties',
    'AwaitableGetResourceLogStreamProperties',
    'get_resource_log_stream',
    'get_resource_log_stream_output',
]

@pulumi.output_type
class GetResourceLogStreamProperties:
    """
    Resource log stream overrides
    """
    def __init__(__self__, endpoint=None, resource_id=None, setting=None):
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if setting and not isinstance(setting, str):
            raise TypeError("Expected argument 'setting' to be a str")
        pulumi.set(__self__, "setting", setting)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[str]:
        """
        The endpoint to stream logs to. Must be present if setting is send. Cannot be present if setting is drop.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        The ID of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def setting(self) -> Optional['GetResourceLogStreamPropertiesSetting']:
        """
        Whether to send logs or drop them.
        """
        return pulumi.get(self, "setting")


class AwaitableGetResourceLogStreamProperties(GetResourceLogStreamProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceLogStreamProperties(
            endpoint=self.endpoint,
            resource_id=self.resource_id,
            setting=self.setting)


def get_resource_log_stream(resource_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceLogStreamProperties:
    """
    Use this data source to access information about an existing resource.

    :param str resource_id: The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
    """
    __args__ = dict()
    __args__['resourceId'] = resource_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:logs:getResourceLogStream', __args__, opts=opts, typ=GetResourceLogStreamProperties).value

    return AwaitableGetResourceLogStreamProperties(
        endpoint=pulumi.get(__ret__, 'endpoint'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        setting=pulumi.get(__ret__, 'setting'))
def get_resource_log_stream_output(resource_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceLogStreamProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str resource_id: The ID of the resource (server, cron job, postgres, or redis) whose log streams should be returned
    """
    __args__ = dict()
    __args__['resourceId'] = resource_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:logs:getResourceLogStream', __args__, opts=opts, typ=GetResourceLogStreamProperties)
    return __ret__.apply(lambda __response__: GetResourceLogStreamProperties(
        endpoint=pulumi.get(__response__, 'endpoint'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        setting=pulumi.get(__response__, 'setting')))