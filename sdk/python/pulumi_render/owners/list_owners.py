# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ListOwnersResult',
    'AwaitableListOwnersResult',
    'list_owners',
]

@pulumi.output_type
class ListOwnersResult:
    def __init__(__self__):
        pass

class AwaitableListOwnersResult(ListOwnersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListOwnersResult()


def list_owners(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListOwnersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:owners:listOwners', __args__, opts=opts, typ=ListOwnersResult).value

    return AwaitableListOwnersResult()
