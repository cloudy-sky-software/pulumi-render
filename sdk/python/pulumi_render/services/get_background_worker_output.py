# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetBackgroundWorkerOutputResult',
    'AwaitableGetBackgroundWorkerOutputResult',
    'get_background_worker_output',
    'get_background_worker_output_output',
]

@pulumi.output_type
class GetBackgroundWorkerOutputResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetBackgroundWorkerOutput':
        return pulumi.get(self, "items")


class AwaitableGetBackgroundWorkerOutputResult(GetBackgroundWorkerOutputResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackgroundWorkerOutputResult(
            items=self.items)


def get_background_worker_output(service_id: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackgroundWorkerOutputResult:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:getBackgroundWorkerOutput', __args__, opts=opts, typ=GetBackgroundWorkerOutputResult).value

    return AwaitableGetBackgroundWorkerOutputResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_background_worker_output)
def get_background_worker_output_output(service_id: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackgroundWorkerOutputResult]:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: The ID of the service
    """
    ...