# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import services

import types

__config__ = pulumi.Config('render')


class _ExportableConfig(types.ModuleType):
    @property
    def api_key(self) -> Optional[str]:
        """
        The Render API key
        """
        return __config__.get('apiKey')

    @property
    def clear_cache_on_service_update_deployments(self) -> Optional[str]:
        """
        When a service is updated, a deployment is automatically triggered. This variable controls whether or not the service cache should be cleared upon deployment.
        """
        return __config__.get('clearCacheOnServiceUpdateDeployments')

