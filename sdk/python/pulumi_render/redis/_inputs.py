# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
from ._enums import *

__all__ = [
    'CidrBlockAndDescriptionArgs',
    'CidrBlockAndDescriptionArgsDict',
]

MYPY = False

if not MYPY:
    class CidrBlockAndDescriptionArgsDict(TypedDict):
        cidr_block: pulumi.Input[_builtins.str]
        description: pulumi.Input[_builtins.str]
        """
        User-provided description of the CIDR block
        """
elif False:
    CidrBlockAndDescriptionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class CidrBlockAndDescriptionArgs:
    def __init__(__self__, *,
                 cidr_block: pulumi.Input[_builtins.str],
                 description: pulumi.Input[_builtins.str]):
        """
        :param pulumi.Input[_builtins.str] description: User-provided description of the CIDR block
        """
        pulumi.set(__self__, "cidr_block", cidr_block)
        pulumi.set(__self__, "description", description)

    @_builtins.property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Input[_builtins.str]:
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "cidr_block", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Input[_builtins.str]:
        """
        User-provided description of the CIDR block
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "description", value)


