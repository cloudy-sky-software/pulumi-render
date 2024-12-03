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
    'User',
    'AwaitableUser',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class User:
    def __init__(__self__, email=None, name=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableUser(User):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return User(
            email=self.email,
            name=self.name)


def get_user(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableUser:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:users:getUser', __args__, opts=opts, typ=User).value

    return AwaitableUser(
        email=pulumi.get(__ret__, 'email'),
        name=pulumi.get(__ret__, 'name'))
def get_user_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[User]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('render:users:getUser', __args__, opts=opts, typ=User)
    return __ret__.apply(lambda __response__: User(
        email=pulumi.get(__response__, 'email'),
        name=pulumi.get(__response__, 'name')))
