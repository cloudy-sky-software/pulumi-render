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
    'ListEnvironmentsResult',
    'AwaitableListEnvironmentsResult',
    'list_environments',
    'list_environments_output',
]

@pulumi.output_type
class ListEnvironmentsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.EnvironmentWithCursor']:
        return pulumi.get(self, "items")


class AwaitableListEnvironmentsResult(ListEnvironmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListEnvironmentsResult(
            items=self.items)


def list_environments(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListEnvironmentsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:environments:listEnvironments', __args__, opts=opts, typ=ListEnvironmentsResult).value

    return AwaitableListEnvironmentsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_environments)
def list_environments_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListEnvironmentsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
