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
from ._enums import *

__all__ = [
    'ListJobResult',
    'AwaitableListJobResult',
    'list_job',
    'list_job_output',
]

@pulumi.output_type
class ListJobResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @_builtins.property
    @pulumi.getter
    def items(self) -> Sequence['outputs.JobWithCursor']:
        return pulumi.get(self, "items")


class AwaitableListJobResult(ListJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListJobResult(
            items=self.items)


def list_job(service_id: Optional[_builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListJobResult:
    """
    Use this data source to access information about an existing resource.

    :param _builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:listJob', __args__, opts=opts, typ=ListJobResult).value

    return AwaitableListJobResult(
        items=pulumi.get(__ret__, 'items'))
def list_job_output(service_id: Optional[pulumi.Input[_builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListJobResult]:
    """
    Use this data source to access information about an existing resource.

    :param _builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:services:listJob', __args__, opts=opts, typ=ListJobResult)
    return __ret__.apply(lambda __response__: ListJobResult(
        items=pulumi.get(__response__, 'items')))
