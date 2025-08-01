# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import pulumi
from enum import Enum

__all__ = [
    'MaintenancePropertiesState',
    'MaxmemoryPolicy',
    'OwnerType',
    'Plan',
    'RedisDetailMaintenancePropertiesState',
    'RedisDetailPlan',
    'RedisDetailRegion',
    'RedisDetailStatus',
    'RedisPlan',
    'RedisRegion',
    'RedisStatus',
    'Region',
    'Status',
]


@pulumi.type_token("render:redis:MaintenancePropertiesState")
class MaintenancePropertiesState(_builtins.str, Enum):
    SCHEDULED = "scheduled"
    IN_PROGRESS = "in_progress"
    USER_FIX_REQUIRED = "user_fix_required"
    CANCELLED = "cancelled"
    SUCCEEDED = "succeeded"
    FAILED = "failed"


@pulumi.type_token("render:redis:MaxmemoryPolicy")
class MaxmemoryPolicy(_builtins.str, Enum):
    """
    The eviction policy for the Key Value instance
    """
    NOEVICTION = "noeviction"
    ALLKEYS_LFU = "allkeys_lfu"
    ALLKEYS_LRU = "allkeys_lru"
    ALLKEYS_RANDOM = "allkeys_random"
    VOLATILE_LFU = "volatile_lfu"
    VOLATILE_LRU = "volatile_lru"
    VOLATILE_RANDOM = "volatile_random"
    VOLATILE_TTL = "volatile_ttl"


@pulumi.type_token("render:redis:OwnerType")
class OwnerType(_builtins.str, Enum):
    USER = "user"
    TEAM = "team"


@pulumi.type_token("render:redis:Plan")
class Plan(_builtins.str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


@pulumi.type_token("render:redis:RedisDetailMaintenancePropertiesState")
class RedisDetailMaintenancePropertiesState(_builtins.str, Enum):
    SCHEDULED = "scheduled"
    IN_PROGRESS = "in_progress"
    USER_FIX_REQUIRED = "user_fix_required"
    CANCELLED = "cancelled"
    SUCCEEDED = "succeeded"
    FAILED = "failed"


@pulumi.type_token("render:redis:RedisDetailPlan")
class RedisDetailPlan(_builtins.str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


@pulumi.type_token("render:redis:RedisDetailRegion")
class RedisDetailRegion(_builtins.str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


@pulumi.type_token("render:redis:RedisDetailStatus")
class RedisDetailStatus(_builtins.str, Enum):
    CREATING = "creating"
    AVAILABLE = "available"
    UNAVAILABLE = "unavailable"
    CONFIG_RESTART = "config_restart"
    SUSPENDED = "suspended"
    MAINTENANCE_SCHEDULED = "maintenance_scheduled"
    MAINTENANCE_IN_PROGRESS = "maintenance_in_progress"
    RECOVERY_FAILED = "recovery_failed"
    RECOVERY_IN_PROGRESS = "recovery_in_progress"
    UNKNOWN = "unknown"
    UPDATING_INSTANCE = "updating_instance"


@pulumi.type_token("render:redis:RedisPlan")
class RedisPlan(_builtins.str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


@pulumi.type_token("render:redis:RedisRegion")
class RedisRegion(_builtins.str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


@pulumi.type_token("render:redis:RedisStatus")
class RedisStatus(_builtins.str, Enum):
    CREATING = "creating"
    AVAILABLE = "available"
    UNAVAILABLE = "unavailable"
    CONFIG_RESTART = "config_restart"
    SUSPENDED = "suspended"
    MAINTENANCE_SCHEDULED = "maintenance_scheduled"
    MAINTENANCE_IN_PROGRESS = "maintenance_in_progress"
    RECOVERY_FAILED = "recovery_failed"
    RECOVERY_IN_PROGRESS = "recovery_in_progress"
    UNKNOWN = "unknown"
    UPDATING_INSTANCE = "updating_instance"


@pulumi.type_token("render:redis:Region")
class Region(_builtins.str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


@pulumi.type_token("render:redis:Status")
class Status(_builtins.str, Enum):
    CREATING = "creating"
    AVAILABLE = "available"
    UNAVAILABLE = "unavailable"
    CONFIG_RESTART = "config_restart"
    SUSPENDED = "suspended"
    MAINTENANCE_SCHEDULED = "maintenance_scheduled"
    MAINTENANCE_IN_PROGRESS = "maintenance_in_progress"
    RECOVERY_FAILED = "recovery_failed"
    RECOVERY_IN_PROGRESS = "recovery_in_progress"
    UNKNOWN = "unknown"
    UPDATING_INSTANCE = "updating_instance"
