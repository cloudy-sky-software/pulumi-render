# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'KeyValueDetailPlan',
    'KeyValueDetailRegion',
    'KeyValueDetailStatus',
    'KeyValuePlan',
    'KeyValueRegion',
    'KeyValueStatus',
    'MaxmemoryPolicy',
    'OwnerType',
    'Plan',
    'RedisDetailpropertiesmaintenanceState',
    'Region',
    'Status',
]


class KeyValueDetailPlan(str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


class KeyValueDetailRegion(str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


class KeyValueDetailStatus(str, Enum):
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


class KeyValuePlan(str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


class KeyValueRegion(str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


class KeyValueStatus(str, Enum):
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


class MaxmemoryPolicy(str, Enum):
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


class OwnerType(str, Enum):
    USER = "user"
    TEAM = "team"


class Plan(str, Enum):
    FREE = "free"
    STARTER = "starter"
    STANDARD = "standard"
    PRO = "pro"
    PRO_PLUS = "pro_plus"
    CUSTOM = "custom"


class RedisDetailpropertiesmaintenanceState(str, Enum):
    SCHEDULED = "scheduled"
    IN_PROGRESS = "in_progress"
    USER_FIX_REQUIRED = "user_fix_required"
    CANCELLED = "cancelled"
    SUCCEEDED = "succeeded"
    FAILED = "failed"


class Region(str, Enum):
    """
    Defaults to "oregon"
    """
    FRANKFURT = "frankfurt"
    OREGON = "oregon"
    OHIO = "ohio"
    SINGAPORE = "singapore"
    VIRGINIA = "virginia"


class Status(str, Enum):
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
