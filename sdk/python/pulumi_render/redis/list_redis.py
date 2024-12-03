# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = [
    'ListRedisResult',
    'AwaitableListRedisResult',
    'list_redis',
    'list_redis_output',
]

@pulumi.output_type
class ListRedisResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.RedisWithCursor']:
        return pulumi.get(self, "items")


class AwaitableListRedisResult(ListRedisResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListRedisResult(
            items=self.items)


def list_redis(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListRedisResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:redis:listRedis', __args__, opts=opts, typ=ListRedisResult).value

    return AwaitableListRedisResult(
        items=pulumi.get(__ret__, 'items'))
def list_redis_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListRedisResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:redis:listRedis', __args__, opts=opts, typ=ListRedisResult)
    return __ret__.apply(lambda __response__: ListRedisResult(
        items=pulumi.get(__response__, 'items')))
