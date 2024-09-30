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
    'GetOwnerLogStreamProperties',
    'AwaitableGetOwnerLogStreamProperties',
    'get_owner_log_stream',
    'get_owner_log_stream_output',
]

@pulumi.output_type
class GetOwnerLogStreamProperties:
    """
    Owner log stream settings
    """
    def __init__(__self__, endpoint=None, owner_id=None, preview=None):
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if preview and not isinstance(preview, str):
            raise TypeError("Expected argument 'preview' to be a str")
        pulumi.set(__self__, "preview", preview)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[str]:
        """
        The endpoint to stream logs to.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[str]:
        """
        The ID of the owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def preview(self) -> Optional['GetOwnerLogStreamPropertiesPreview']:
        """
        Whether to send logs or drop them.
        """
        return pulumi.get(self, "preview")


class AwaitableGetOwnerLogStreamProperties(GetOwnerLogStreamProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOwnerLogStreamProperties(
            endpoint=self.endpoint,
            owner_id=self.owner_id,
            preview=self.preview)


def get_owner_log_stream(owner_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOwnerLogStreamProperties:
    """
    Use this data source to access information about an existing resource.

    :param str owner_id: The ID of the owner (team or personal user) whose log streams should be returned
    """
    __args__ = dict()
    __args__['ownerId'] = owner_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:logs:getOwnerLogStream', __args__, opts=opts, typ=GetOwnerLogStreamProperties).value

    return AwaitableGetOwnerLogStreamProperties(
        endpoint=pulumi.get(__ret__, 'endpoint'),
        owner_id=pulumi.get(__ret__, 'owner_id'),
        preview=pulumi.get(__ret__, 'preview'))
def get_owner_log_stream_output(owner_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOwnerLogStreamProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str owner_id: The ID of the owner (team or personal user) whose log streams should be returned
    """
    __args__ = dict()
    __args__['ownerId'] = owner_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:logs:getOwnerLogStream', __args__, opts=opts, typ=GetOwnerLogStreamProperties)
    return __ret__.apply(lambda __response__: GetOwnerLogStreamProperties(
        endpoint=pulumi.get(__response__, 'endpoint'),
        owner_id=pulumi.get(__response__, 'owner_id'),
        preview=pulumi.get(__response__, 'preview')))
