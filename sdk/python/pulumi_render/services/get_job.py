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

__all__ = [
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.Job':
        return pulumi.get(self, "items")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            items=self.items)


def get_job(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_job)
def get_job_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
