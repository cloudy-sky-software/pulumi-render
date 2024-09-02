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
    'EnvVarInputArgs',
    'SecretFileInputArgs',
]

@pulumi.input_type
class EnvVarInputArgs:
    def __init__(__self__):
        pass


@pulumi.input_type
class SecretFileInputArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 name: pulumi.Input[str]):
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


