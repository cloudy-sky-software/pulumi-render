# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'Owner',
    'OwnerWithCursor',
]

@pulumi.output_type
class Owner(dict):
    def __init__(__self__, *,
                 email: str,
                 id: str,
                 name: str,
                 type: 'OwnerType',
                 two_factor_auth_enabled: Optional[bool] = None):
        """
        :param bool two_factor_auth_enabled: Whether two-factor authentication is enabled for the owner. Only present for user owners.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if two_factor_auth_enabled is not None:
            pulumi.set(__self__, "two_factor_auth_enabled", two_factor_auth_enabled)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> 'OwnerType':
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="twoFactorAuthEnabled")
    def two_factor_auth_enabled(self) -> Optional[bool]:
        """
        Whether two-factor authentication is enabled for the owner. Only present for user owners.
        """
        return pulumi.get(self, "two_factor_auth_enabled")


@pulumi.output_type
class OwnerWithCursor(dict):
    def __init__(__self__, *,
                 cursor: Optional[str] = None,
                 owner: Optional['outputs.Owner'] = None):
        if cursor is not None:
            pulumi.set(__self__, "cursor", cursor)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)

    @property
    @pulumi.getter
    def cursor(self) -> Optional[str]:
        return pulumi.get(self, "cursor")

    @property
    @pulumi.getter
    def owner(self) -> Optional['outputs.Owner']:
        return pulumi.get(self, "owner")


