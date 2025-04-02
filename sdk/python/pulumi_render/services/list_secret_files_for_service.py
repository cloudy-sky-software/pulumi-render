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

__all__ = [
    'ListSecretFilesForServiceResult',
    'AwaitableListSecretFilesForServiceResult',
    'list_secret_files_for_service',
    'list_secret_files_for_service_output',
]

@pulumi.output_type
class ListSecretFilesForServiceResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.SecretFileWithCursor']:
        return pulumi.get(self, "items")


class AwaitableListSecretFilesForServiceResult(ListSecretFilesForServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSecretFilesForServiceResult(
            items=self.items)


def list_secret_files_for_service(service_id: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSecretFilesForServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:listSecretFilesForService', __args__, opts=opts, typ=ListSecretFilesForServiceResult).value

    return AwaitableListSecretFilesForServiceResult(
        items=pulumi.get(__ret__, 'items'))
def list_secret_files_for_service_output(service_id: Optional[pulumi.Input[builtins.str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListSecretFilesForServiceResult]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str service_id: The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:services:listSecretFilesForService', __args__, opts=opts, typ=ListSecretFilesForServiceResult)
    return __ret__.apply(lambda __response__: ListSecretFilesForServiceResult(
        items=pulumi.get(__response__, 'items')))
