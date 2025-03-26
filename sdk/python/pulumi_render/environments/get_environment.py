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
from ._enums import *

__all__ = [
    'Environment',
    'AwaitableEnvironment',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class Environment:
    def __init__(__self__, databases_ids=None, env_group_ids=None, id=None, name=None, network_isolation_enabled=None, project_id=None, protected_status=None, redis_ids=None, service_ids=None):
        if databases_ids and not isinstance(databases_ids, list):
            raise TypeError("Expected argument 'databases_ids' to be a list")
        pulumi.set(__self__, "databases_ids", databases_ids)
        if env_group_ids and not isinstance(env_group_ids, list):
            raise TypeError("Expected argument 'env_group_ids' to be a list")
        pulumi.set(__self__, "env_group_ids", env_group_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_isolation_enabled and not isinstance(network_isolation_enabled, bool):
            raise TypeError("Expected argument 'network_isolation_enabled' to be a bool")
        pulumi.set(__self__, "network_isolation_enabled", network_isolation_enabled)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if protected_status and not isinstance(protected_status, str):
            raise TypeError("Expected argument 'protected_status' to be a str")
        pulumi.set(__self__, "protected_status", protected_status)
        if redis_ids and not isinstance(redis_ids, list):
            raise TypeError("Expected argument 'redis_ids' to be a list")
        pulumi.set(__self__, "redis_ids", redis_ids)
        if service_ids and not isinstance(service_ids, list):
            raise TypeError("Expected argument 'service_ids' to be a list")
        pulumi.set(__self__, "service_ids", service_ids)

    @property
    @pulumi.getter(name="databasesIds")
    def databases_ids(self) -> Sequence[str]:
        return pulumi.get(self, "databases_ids")

    @property
    @pulumi.getter(name="envGroupIds")
    def env_group_ids(self) -> Sequence[str]:
        return pulumi.get(self, "env_group_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkIsolationEnabled")
    def network_isolation_enabled(self) -> bool:
        """
        Indicates whether network connections across environments are allowed.
        """
        return pulumi.get(self, "network_isolation_enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="protectedStatus")
    def protected_status(self) -> 'EnvironmentProtectedStatus':
        """
        Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        return pulumi.get(self, "protected_status")

    @property
    @pulumi.getter(name="redisIds")
    def redis_ids(self) -> Sequence[str]:
        return pulumi.get(self, "redis_ids")

    @property
    @pulumi.getter(name="serviceIds")
    def service_ids(self) -> Sequence[str]:
        return pulumi.get(self, "service_ids")


class AwaitableEnvironment(Environment):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Environment(
            databases_ids=self.databases_ids,
            env_group_ids=self.env_group_ids,
            id=self.id,
            name=self.name,
            network_isolation_enabled=self.network_isolation_enabled,
            project_id=self.project_id,
            protected_status=self.protected_status,
            redis_ids=self.redis_ids,
            service_ids=self.service_ids)


def get_environment(environment_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableEnvironment:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:environments:getEnvironment', __args__, opts=opts, typ=Environment).value

    return AwaitableEnvironment(
        databases_ids=pulumi.get(__ret__, 'databases_ids'),
        env_group_ids=pulumi.get(__ret__, 'env_group_ids'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        network_isolation_enabled=pulumi.get(__ret__, 'network_isolation_enabled'),
        project_id=pulumi.get(__ret__, 'project_id'),
        protected_status=pulumi.get(__ret__, 'protected_status'),
        redis_ids=pulumi.get(__ret__, 'redis_ids'),
        service_ids=pulumi.get(__ret__, 'service_ids'))
def get_environment_output(environment_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[Environment]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:environments:getEnvironment', __args__, opts=opts, typ=Environment)
    return __ret__.apply(lambda __response__: Environment(
        databases_ids=pulumi.get(__response__, 'databases_ids'),
        env_group_ids=pulumi.get(__response__, 'env_group_ids'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        network_isolation_enabled=pulumi.get(__response__, 'network_isolation_enabled'),
        project_id=pulumi.get(__response__, 'project_id'),
        protected_status=pulumi.get(__response__, 'protected_status'),
        redis_ids=pulumi.get(__response__, 'redis_ids'),
        service_ids=pulumi.get(__response__, 'service_ids')))
