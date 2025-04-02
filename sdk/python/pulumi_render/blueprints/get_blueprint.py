# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetBlueprintProperties',
    'AwaitableGetBlueprintProperties',
    'get_blueprint',
    'get_blueprint_output',
]

@pulumi.output_type
class GetBlueprintProperties:
    def __init__(__self__, auto_sync=None, branch=None, id=None, last_sync=None, name=None, repo=None, resources=None, status=None):
        if auto_sync and not isinstance(auto_sync, bool):
            raise TypeError("Expected argument 'auto_sync' to be a bool")
        pulumi.set(__self__, "auto_sync", auto_sync)
        if branch and not isinstance(branch, str):
            raise TypeError("Expected argument 'branch' to be a str")
        pulumi.set(__self__, "branch", branch)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_sync and not isinstance(last_sync, str):
            raise TypeError("Expected argument 'last_sync' to be a str")
        pulumi.set(__self__, "last_sync", last_sync)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if repo and not isinstance(repo, str):
            raise TypeError("Expected argument 'repo' to be a str")
        pulumi.set(__self__, "repo", repo)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="autoSync")
    def auto_sync(self) -> builtins.bool:
        """
        Automatically sync changes to render.yaml
        """
        return pulumi.get(self, "auto_sync")

    @property
    @pulumi.getter
    def branch(self) -> builtins.str:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastSync")
    def last_sync(self) -> Optional[builtins.str]:
        return pulumi.get(self, "last_sync")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def repo(self) -> builtins.str:
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter
    def resources(self) -> Sequence['outputs.GetBlueprintPropertiesResourcesItemProperties']:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def status(self) -> 'GetBlueprintPropertiesStatus':
        return pulumi.get(self, "status")


class AwaitableGetBlueprintProperties(GetBlueprintProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlueprintProperties(
            auto_sync=self.auto_sync,
            branch=self.branch,
            id=self.id,
            last_sync=self.last_sync,
            name=self.name,
            repo=self.repo,
            resources=self.resources,
            status=self.status)


def get_blueprint(blueprint_id: Optional[builtins.str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlueprintProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str blueprint_id: The ID of the blueprint
    """
    __args__ = dict()
    __args__['blueprintId'] = blueprint_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:blueprints:getBlueprint', __args__, opts=opts, typ=GetBlueprintProperties).value

    return AwaitableGetBlueprintProperties(
        auto_sync=pulumi.get(__ret__, 'auto_sync'),
        branch=pulumi.get(__ret__, 'branch'),
        id=pulumi.get(__ret__, 'id'),
        last_sync=pulumi.get(__ret__, 'last_sync'),
        name=pulumi.get(__ret__, 'name'),
        repo=pulumi.get(__ret__, 'repo'),
        resources=pulumi.get(__ret__, 'resources'),
        status=pulumi.get(__ret__, 'status'))
def get_blueprint_output(blueprint_id: Optional[pulumi.Input[builtins.str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBlueprintProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str blueprint_id: The ID of the blueprint
    """
    __args__ = dict()
    __args__['blueprintId'] = blueprint_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:blueprints:getBlueprint', __args__, opts=opts, typ=GetBlueprintProperties)
    return __ret__.apply(lambda __response__: GetBlueprintProperties(
        auto_sync=pulumi.get(__response__, 'auto_sync'),
        branch=pulumi.get(__response__, 'branch'),
        id=pulumi.get(__response__, 'id'),
        last_sync=pulumi.get(__response__, 'last_sync'),
        name=pulumi.get(__response__, 'name'),
        repo=pulumi.get(__response__, 'repo'),
        resources=pulumi.get(__response__, 'resources'),
        status=pulumi.get(__response__, 'status')))
