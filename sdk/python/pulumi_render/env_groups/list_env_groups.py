# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'ListEnvGroupsResult',
    'AwaitableListEnvGroupsResult',
    'list_env_groups',
    'list_env_groups_output',
]

@pulumi.output_type
class ListEnvGroupsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.EnvGroupMeta']:
        return pulumi.get(self, "items")


class AwaitableListEnvGroupsResult(ListEnvGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListEnvGroupsResult(
            items=self.items)


def list_env_groups(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListEnvGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:env-groups:listEnvGroups', __args__, opts=opts, typ=ListEnvGroupsResult).value

    return AwaitableListEnvGroupsResult(
        items=pulumi.get(__ret__, 'items'))
def list_env_groups_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListEnvGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:env-groups:listEnvGroups', __args__, opts=opts, typ=ListEnvGroupsResult)
    return __ret__.apply(lambda __response__: ListEnvGroupsResult(
        items=pulumi.get(__response__, 'items')))
