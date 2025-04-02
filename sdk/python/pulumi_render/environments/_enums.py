# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'EnvironmentProtectedStatus',
    'ProtectedStatus',
]


class EnvironmentProtectedStatus(builtins.str, Enum):
    """
    Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
    """
    UNPROTECTED = "unprotected"
    PROTECTED = "protected"


class ProtectedStatus(builtins.str, Enum):
    """
    Indicates whether an environment is `unprotected` or `protected`. Only admin users can perform destructive actions in `protected` environments.
    """
    UNPROTECTED = "unprotected"
    PROTECTED = "protected"
