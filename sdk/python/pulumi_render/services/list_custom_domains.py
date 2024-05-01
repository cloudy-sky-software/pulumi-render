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
    'ListCustomDomainsResult',
    'AwaitableListCustomDomainsResult',
    'list_custom_domains',
    'list_custom_domains_output',
]

@pulumi.output_type
class ListCustomDomainsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.ListCustomDomainsResponse']:
        return pulumi.get(self, "items")


class AwaitableListCustomDomainsResult(ListCustomDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListCustomDomainsResult(
            items=self.items)


def list_custom_domains(service_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListCustomDomainsResult:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: (Required) The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:listCustomDomains', __args__, opts=opts, typ=ListCustomDomainsResult).value

    return AwaitableListCustomDomainsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_custom_domains)
def list_custom_domains_output(service_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListCustomDomainsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: (Required) The ID of the service
    """
    ...
