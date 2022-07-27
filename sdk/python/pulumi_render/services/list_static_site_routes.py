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
    'ListStaticSiteRoutesResult',
    'AwaitableListStaticSiteRoutesResult',
    'list_static_site_routes',
    'list_static_site_routes_output',
]

@pulumi.output_type
class ListStaticSiteRoutesResult:
    def __init__(__self__):
        pass

class AwaitableListStaticSiteRoutesResult(ListStaticSiteRoutesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListStaticSiteRoutesResult()


def list_static_site_routes(service_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListStaticSiteRoutesResult:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: (Required) The ID of the service
    """
    __args__ = dict()
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:services:listStaticSiteRoutes', __args__, opts=opts, typ=ListStaticSiteRoutesResult).value

    return AwaitableListStaticSiteRoutesResult()


@_utilities.lift_output_func(list_static_site_routes)
def list_static_site_routes_output(service_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListStaticSiteRoutesResult]:
    """
    Use this data source to access information about an existing resource.

    :param str service_id: (Required) The ID of the service
    """
    ...