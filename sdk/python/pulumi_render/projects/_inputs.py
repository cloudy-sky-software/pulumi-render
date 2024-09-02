# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ProjectCreateEnvironmentInputArgs',
]

@pulumi.input_type
class ProjectCreateEnvironmentInputArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 protected_status: Optional[pulumi.Input['ProjectCreateEnvironmentInputProtectedStatus']] = None):
        """
        :param pulumi.Input['ProjectCreateEnvironmentInputProtectedStatus'] protected_status: Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        pulumi.set(__self__, "name", name)
        if protected_status is not None:
            pulumi.set(__self__, "protected_status", protected_status)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="protectedStatus")
    def protected_status(self) -> Optional[pulumi.Input['ProjectCreateEnvironmentInputProtectedStatus']]:
        """
        Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
        """
        return pulumi.get(self, "protected_status")

    @protected_status.setter
    def protected_status(self, value: Optional[pulumi.Input['ProjectCreateEnvironmentInputProtectedStatus']]):
        pulumi.set(self, "protected_status", value)


