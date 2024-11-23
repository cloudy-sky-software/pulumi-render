# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'RedisDetail',
    'AwaitableRedisDetail',
    'get_redis',
    'get_redis_output',
]

@pulumi.output_type
class RedisDetail:
    """
    A Redis instance
    """
    def __init__(__self__, created_at=None, environment_id=None, id=None, ip_allow_list=None, maintenance=None, name=None, options=None, owner=None, plan=None, region=None, status=None, updated_at=None, version=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if environment_id and not isinstance(environment_id, str):
            raise TypeError("Expected argument 'environment_id' to be a str")
        pulumi.set(__self__, "environment_id", environment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_allow_list and not isinstance(ip_allow_list, list):
            raise TypeError("Expected argument 'ip_allow_list' to be a list")
        pulumi.set(__self__, "ip_allow_list", ip_allow_list)
        if maintenance and not isinstance(maintenance, dict):
            raise TypeError("Expected argument 'maintenance' to be a dict")
        pulumi.set(__self__, "maintenance", maintenance)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if owner and not isinstance(owner, dict):
            raise TypeError("Expected argument 'owner' to be a dict")
        pulumi.set(__self__, "owner", owner)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The creation time of the Redis instance
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[str]:
        """
        The ID of the environment the Redis instance is associated with
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Redis instance
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAllowList")
    def ip_allow_list(self) -> Sequence['outputs.CidrBlockAndDescription']:
        """
        The IP allow list for the Redis instance
        """
        return pulumi.get(self, "ip_allow_list")

    @property
    @pulumi.getter
    def maintenance(self) -> Optional['outputs.RedisDetailMaintenanceProperties']:
        return pulumi.get(self, "maintenance")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Redis instance
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> 'outputs.RedisOptions':
        """
        Options for a Redis instance
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def owner(self) -> 'outputs.Owner':
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def plan(self) -> 'RedisDetailPlan':
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def region(self) -> 'RedisDetailRegion':
        """
        Defaults to "oregon"
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> 'RedisDetailStatus':
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The last updated time of the Redis instance
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of Redis
        """
        return pulumi.get(self, "version")


class AwaitableRedisDetail(RedisDetail):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RedisDetail(
            created_at=self.created_at,
            environment_id=self.environment_id,
            id=self.id,
            ip_allow_list=self.ip_allow_list,
            maintenance=self.maintenance,
            name=self.name,
            options=self.options,
            owner=self.owner,
            plan=self.plan,
            region=self.region,
            status=self.status,
            updated_at=self.updated_at,
            version=self.version)


def get_redis(redis_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRedisDetail:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['redisId'] = redis_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:redis:getRedis', __args__, opts=opts, typ=RedisDetail).value

    return AwaitableRedisDetail(
        created_at=pulumi.get(__ret__, 'created_at'),
        environment_id=pulumi.get(__ret__, 'environment_id'),
        id=pulumi.get(__ret__, 'id'),
        ip_allow_list=pulumi.get(__ret__, 'ip_allow_list'),
        maintenance=pulumi.get(__ret__, 'maintenance'),
        name=pulumi.get(__ret__, 'name'),
        options=pulumi.get(__ret__, 'options'),
        owner=pulumi.get(__ret__, 'owner'),
        plan=pulumi.get(__ret__, 'plan'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        version=pulumi.get(__ret__, 'version'))
def get_redis_output(redis_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[RedisDetail]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['redisId'] = redis_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:redis:getRedis', __args__, opts=opts, typ=RedisDetail)
    return __ret__.apply(lambda __response__: RedisDetail(
        created_at=pulumi.get(__response__, 'created_at'),
        environment_id=pulumi.get(__response__, 'environment_id'),
        id=pulumi.get(__response__, 'id'),
        ip_allow_list=pulumi.get(__response__, 'ip_allow_list'),
        maintenance=pulumi.get(__response__, 'maintenance'),
        name=pulumi.get(__response__, 'name'),
        options=pulumi.get(__response__, 'options'),
        owner=pulumi.get(__response__, 'owner'),
        plan=pulumi.get(__response__, 'plan'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        version=pulumi.get(__response__, 'version')))
