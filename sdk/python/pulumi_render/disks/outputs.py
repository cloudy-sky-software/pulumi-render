# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = [
    'DiskWithCursor',
    'DiskWithCursorDiskProperties',
]

@pulumi.output_type
class DiskWithCursor(dict):
    def __init__(__self__, *,
                 cursor: builtins.str,
                 disk: 'outputs.DiskWithCursorDiskProperties'):
        pulumi.set(__self__, "cursor", cursor)
        pulumi.set(__self__, "disk", disk)

    @property
    @pulumi.getter
    def cursor(self) -> builtins.str:
        return pulumi.get(self, "cursor")

    @property
    @pulumi.getter
    def disk(self) -> 'outputs.DiskWithCursorDiskProperties':
        return pulumi.get(self, "disk")


@pulumi.output_type
class DiskWithCursorDiskProperties(dict):
    def __init__(__self__, *,
                 created_at: builtins.str,
                 id: builtins.str,
                 mount_path: builtins.str,
                 name: builtins.str,
                 size_gb: builtins.int,
                 updated_at: builtins.str,
                 service_id: Optional[builtins.str] = None):
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "mount_path", mount_path)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "size_gb", size_gb)
        pulumi.set(__self__, "updated_at", updated_at)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mountPath")
    def mount_path(self) -> builtins.str:
        return pulumi.get(self, "mount_path")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sizeGB")
    def size_gb(self) -> builtins.int:
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "service_id")


