# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'GetOwnerNotificationSettingPropertiesNotificationsToSend',
    'GetServiceNotificationOverridePropertiesNotificationsToSend',
    'GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled',
    'NotificationOverrideWithCursorOverridePropertiesNotificationsToSend',
    'NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled',
]


@pulumi.type_token("render:notification-settings:GetOwnerNotificationSettingPropertiesNotificationsToSend")
class GetOwnerNotificationSettingPropertiesNotificationsToSend(builtins.str, Enum):
    NONE = "none"
    FAILURE = "failure"
    ALL = "all"


@pulumi.type_token("render:notification-settings:GetServiceNotificationOverridePropertiesNotificationsToSend")
class GetServiceNotificationOverridePropertiesNotificationsToSend(builtins.str, Enum):
    DEFAULT = "default"
    NONE = "none"
    FAILURE = "failure"
    ALL = "all"


@pulumi.type_token("render:notification-settings:GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled")
class GetServiceNotificationOverridePropertiesPreviewNotificationsEnabled(builtins.str, Enum):
    DEFAULT = "default"
    FALSE = "false"
    TRUE = "true"


@pulumi.type_token("render:notification-settings:NotificationOverrideWithCursorOverridePropertiesNotificationsToSend")
class NotificationOverrideWithCursorOverridePropertiesNotificationsToSend(builtins.str, Enum):
    DEFAULT = "default"
    NONE = "none"
    FAILURE = "failure"
    ALL = "all"


@pulumi.type_token("render:notification-settings:NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled")
class NotificationOverrideWithCursorOverridePropertiesPreviewNotificationsEnabled(builtins.str, Enum):
    DEFAULT = "default"
    FALSE = "false"
    TRUE = "true"
