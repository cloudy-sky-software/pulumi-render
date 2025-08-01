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
from . import outputs
from ._enums import *

__all__ = [
    'WebhookWithCursor',
    'WebhookWithCursorWebhookProperties',
]

@pulumi.output_type
class WebhookWithCursor(dict):
    def __init__(__self__, *,
                 cursor: _builtins.str,
                 webhook: 'outputs.WebhookWithCursorWebhookProperties'):
        pulumi.set(__self__, "cursor", cursor)
        pulumi.set(__self__, "webhook", webhook)

    @_builtins.property
    @pulumi.getter
    def cursor(self) -> _builtins.str:
        return pulumi.get(self, "cursor")

    @_builtins.property
    @pulumi.getter
    def webhook(self) -> 'outputs.WebhookWithCursorWebhookProperties':
        return pulumi.get(self, "webhook")


@pulumi.output_type
class WebhookWithCursorWebhookProperties(dict):
    def __init__(__self__, *,
                 enabled: _builtins.bool,
                 event_filter: Sequence['WebhookWithCursorWebhookPropertiesEventFilterItem'],
                 id: _builtins.str,
                 name: _builtins.str,
                 secret: _builtins.str,
                 url: _builtins.str):
        """
        :param Sequence['WebhookWithCursorWebhookPropertiesEventFilterItem'] event_filter: The event types that will trigger the webhook. An empty list means all event types will trigger the webhook.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "event_filter", event_filter)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "secret", secret)
        pulumi.set(__self__, "url", url)

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> _builtins.bool:
        return pulumi.get(self, "enabled")

    @_builtins.property
    @pulumi.getter(name="eventFilter")
    def event_filter(self) -> Sequence['WebhookWithCursorWebhookPropertiesEventFilterItem']:
        """
        The event types that will trigger the webhook. An empty list means all event types will trigger the webhook.
        """
        return pulumi.get(self, "event_filter")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def secret(self) -> _builtins.str:
        return pulumi.get(self, "secret")

    @_builtins.property
    @pulumi.getter
    def url(self) -> _builtins.str:
        return pulumi.get(self, "url")


