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

__all__ = [
    'RedisConnectionInfo',
    'AwaitableRedisConnectionInfo',
    'get_redis_connection_info',
    'get_redis_connection_info_output',
]

@pulumi.output_type
class RedisConnectionInfo:
    """
    A Redis instance
    """
    def __init__(__self__, external_connection_string=None, internal_connection_string=None, redis_cli_command=None):
        if external_connection_string and not isinstance(external_connection_string, str):
            raise TypeError("Expected argument 'external_connection_string' to be a str")
        pulumi.set(__self__, "external_connection_string", external_connection_string)
        if internal_connection_string and not isinstance(internal_connection_string, str):
            raise TypeError("Expected argument 'internal_connection_string' to be a str")
        pulumi.set(__self__, "internal_connection_string", internal_connection_string)
        if redis_cli_command and not isinstance(redis_cli_command, str):
            raise TypeError("Expected argument 'redis_cli_command' to be a str")
        pulumi.set(__self__, "redis_cli_command", redis_cli_command)

    @property
    @pulumi.getter(name="externalConnectionString")
    def external_connection_string(self) -> str:
        """
        The connection string to use from outside Render
        """
        return pulumi.get(self, "external_connection_string")

    @property
    @pulumi.getter(name="internalConnectionString")
    def internal_connection_string(self) -> str:
        """
        The connection string to use from within Render
        """
        return pulumi.get(self, "internal_connection_string")

    @property
    @pulumi.getter(name="redisCLICommand")
    def redis_cli_command(self) -> str:
        """
        The Redis CLI command to connect to the Redis instance
        """
        return pulumi.get(self, "redis_cli_command")


class AwaitableRedisConnectionInfo(RedisConnectionInfo):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RedisConnectionInfo(
            external_connection_string=self.external_connection_string,
            internal_connection_string=self.internal_connection_string,
            redis_cli_command=self.redis_cli_command)


def get_redis_connection_info(redis_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRedisConnectionInfo:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['redisId'] = redis_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:redis:getRedisConnectionInfo', __args__, opts=opts, typ=RedisConnectionInfo).value

    return AwaitableRedisConnectionInfo(
        external_connection_string=pulumi.get(__ret__, 'external_connection_string'),
        internal_connection_string=pulumi.get(__ret__, 'internal_connection_string'),
        redis_cli_command=pulumi.get(__ret__, 'redis_cli_command'))
def get_redis_connection_info_output(redis_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[RedisConnectionInfo]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['redisId'] = redis_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:redis:getRedisConnectionInfo', __args__, opts=opts, typ=RedisConnectionInfo)
    return __ret__.apply(lambda __response__: RedisConnectionInfo(
        external_connection_string=pulumi.get(__response__, 'external_connection_string'),
        internal_connection_string=pulumi.get(__response__, 'internal_connection_string'),
        redis_cli_command=pulumi.get(__response__, 'redis_cli_command')))
