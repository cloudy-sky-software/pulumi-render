# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'Registry',
    'RegistryCredentialRegistry',
]


class Registry(str, Enum):
    GITHUB = "GITHUB"
    GITLAB = "GITLAB"
    DOCKER = "DOCKER"


class RegistryCredentialRegistry(str, Enum):
    """
    The registry to use this credential with
    """
    GITHUB = "GITHUB"
    GITLAB = "GITLAB"
    DOCKER = "DOCKER"
