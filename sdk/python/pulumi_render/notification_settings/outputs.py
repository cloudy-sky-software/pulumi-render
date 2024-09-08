# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'NotificationOverrideWithCursor',
    'NotificationOverrideWithCursorOverrideProperties',
]

@pulumi.output_type
class NotificationOverrideWithCursor(dict):
    def __init__(__self__, *,
                 cursor: str,
                 override: 'outputs.NotificationOverrideWithCursorOverrideProperties'):
        pulumi.set(__self__, "cursor", cursor)
        pulumi.set(__self__, "override", override)

    @property
    @pulumi.getter
    def cursor(self) -> str:
        return pulumi.get(self, "cursor")

    @property
    @pulumi.getter
    def override(self) -> 'outputs.NotificationOverrideWithCursorOverrideProperties':
        return pulumi.get(self, "override")


@pulumi.output_type
class NotificationOverrideWithCursorOverrideProperties(dict):
    def __init__(__self__, *,
                 notifications_to_send: 'NotificationOverrideWithCursorOverridePropertiesNotificationsToSend',
                 preview_notifications_enabled: 'NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled',
                 service_id: str):
        pulumi.set(__self__, "notifications_to_send", notifications_to_send)
        pulumi.set(__self__, "preview_notifications_enabled", preview_notifications_enabled)
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="notificationsToSend")
    def notifications_to_send(self) -> 'NotificationOverrideWithCursorOverridePropertiesNotificationsToSend':
        return pulumi.get(self, "notifications_to_send")

    @property
    @pulumi.getter(name="previewNotificationsEnabled")
    def preview_notifications_enabled(self) -> 'NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled':
        return pulumi.get(self, "preview_notifications_enabled")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")

