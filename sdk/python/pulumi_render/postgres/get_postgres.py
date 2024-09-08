# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'PostgresDetail',
    'AwaitablePostgresDetail',
    'get_postgres',
    'get_postgres_output',
]

@pulumi.output_type
class PostgresDetail:
    def __init__(__self__, created_at=None, dashboard_url=None, database_name=None, database_user=None, disk_size_gb=None, environment_id=None, expires_at=None, high_availability_enabled=None, id=None, ip_allow_list=None, maintenance=None, name=None, owner=None, plan=None, primary_postgres_id=None, read_replicas=None, region=None, role=None, status=None, suspended=None, suspenders=None, updated_at=None, version=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dashboard_url and not isinstance(dashboard_url, str):
            raise TypeError("Expected argument 'dashboard_url' to be a str")
        pulumi.set(__self__, "dashboard_url", dashboard_url)
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if database_user and not isinstance(database_user, str):
            raise TypeError("Expected argument 'database_user' to be a str")
        pulumi.set(__self__, "database_user", database_user)
        if disk_size_gb and not isinstance(disk_size_gb, int):
            raise TypeError("Expected argument 'disk_size_gb' to be a int")
        pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if environment_id and not isinstance(environment_id, str):
            raise TypeError("Expected argument 'environment_id' to be a str")
        pulumi.set(__self__, "environment_id", environment_id)
        if expires_at and not isinstance(expires_at, str):
            raise TypeError("Expected argument 'expires_at' to be a str")
        pulumi.set(__self__, "expires_at", expires_at)
        if high_availability_enabled and not isinstance(high_availability_enabled, bool):
            raise TypeError("Expected argument 'high_availability_enabled' to be a bool")
        pulumi.set(__self__, "high_availability_enabled", high_availability_enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_allow_list and not isinstance(ip_allow_list, list):
            raise TypeError("Expected argument 'ip_allow_list' to be a list")
        pulumi.set(__self__, "ip_allow_list", ip_allow_list)
        if maintenance and not isinstance(maintenance, dict):
            raise TypeError("Expected argument 'maintenance' to be a dict")
        pulumi.set(__self__, "maintenance", maintenance)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, dict):
            raise TypeError("Expected argument 'owner' to be a dict")
        pulumi.set(__self__, "owner", owner)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if primary_postgres_id and not isinstance(primary_postgres_id, str):
            raise TypeError("Expected argument 'primary_postgres_id' to be a str")
        pulumi.set(__self__, "primary_postgres_id", primary_postgres_id)
        if read_replicas and not isinstance(read_replicas, list):
            raise TypeError("Expected argument 'read_replicas' to be a list")
        pulumi.set(__self__, "read_replicas", read_replicas)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if suspended and not isinstance(suspended, str):
            raise TypeError("Expected argument 'suspended' to be a str")
        pulumi.set(__self__, "suspended", suspended)
        if suspenders and not isinstance(suspenders, list):
            raise TypeError("Expected argument 'suspenders' to be a list")
        pulumi.set(__self__, "suspenders", suspenders)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dashboardUrl")
    def dashboard_url(self) -> str:
        """
        The URL to view the PostgreSQL instance in the Render Dashboard
        """
        return pulumi.get(self, "dashboard_url")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="databaseUser")
    def database_user(self) -> str:
        return pulumi.get(self, "database_user")

    @property
    @pulumi.getter(name="diskSizeGB")
    def disk_size_gb(self) -> Optional[int]:
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[str]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[str]:
        """
        The time at which the database will be expire. Applies to free tier databases only.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="highAvailabilityEnabled")
    def high_availability_enabled(self) -> bool:
        return pulumi.get(self, "high_availability_enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAllowList")
    def ip_allow_list(self) -> Sequence['outputs.CidrBlockAndDescription']:
        return pulumi.get(self, "ip_allow_list")

    @property
    @pulumi.getter
    def maintenance(self) -> Optional['outputs.RedisDetailpropertiesmaintenance']:
        return pulumi.get(self, "maintenance")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> 'outputs.Owner':
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def plan(self) -> 'PostgresDetailPlan':
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="primaryPostgresID")
    def primary_postgres_id(self) -> Optional[str]:
        return pulumi.get(self, "primary_postgres_id")

    @property
    @pulumi.getter(name="readReplicas")
    def read_replicas(self) -> Sequence['outputs.ReadReplica']:
        return pulumi.get(self, "read_replicas")

    @property
    @pulumi.getter
    def region(self) -> 'PostgresDetailRegion':
        """
        Defaults to "oregon"
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def role(self) -> 'PostgresDetailRole':
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def status(self) -> 'PostgresDetailStatus':
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def suspended(self) -> 'PostgresDetailSuspended':
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def suspenders(self) -> Sequence['PostgresDetailSuspendersItem']:
        return pulumi.get(self, "suspenders")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> 'PostgresDetailVersion':
        """
        The PostgreSQL version
        """
        return pulumi.get(self, "version")


class AwaitablePostgresDetail(PostgresDetail):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return PostgresDetail(
            created_at=self.created_at,
            dashboard_url=self.dashboard_url,
            database_name=self.database_name,
            database_user=self.database_user,
            disk_size_gb=self.disk_size_gb,
            environment_id=self.environment_id,
            expires_at=self.expires_at,
            high_availability_enabled=self.high_availability_enabled,
            id=self.id,
            ip_allow_list=self.ip_allow_list,
            maintenance=self.maintenance,
            name=self.name,
            owner=self.owner,
            plan=self.plan,
            primary_postgres_id=self.primary_postgres_id,
            read_replicas=self.read_replicas,
            region=self.region,
            role=self.role,
            status=self.status,
            suspended=self.suspended,
            suspenders=self.suspenders,
            updated_at=self.updated_at,
            version=self.version)


def get_postgres(postgres_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitablePostgresDetail:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['postgresId'] = postgres_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('render:postgres:getPostgres', __args__, opts=opts, typ=PostgresDetail).value

    return AwaitablePostgresDetail(
        created_at=pulumi.get(__ret__, 'created_at'),
        dashboard_url=pulumi.get(__ret__, 'dashboard_url'),
        database_name=pulumi.get(__ret__, 'database_name'),
        database_user=pulumi.get(__ret__, 'database_user'),
        disk_size_gb=pulumi.get(__ret__, 'disk_size_gb'),
        environment_id=pulumi.get(__ret__, 'environment_id'),
        expires_at=pulumi.get(__ret__, 'expires_at'),
        high_availability_enabled=pulumi.get(__ret__, 'high_availability_enabled'),
        id=pulumi.get(__ret__, 'id'),
        ip_allow_list=pulumi.get(__ret__, 'ip_allow_list'),
        maintenance=pulumi.get(__ret__, 'maintenance'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        plan=pulumi.get(__ret__, 'plan'),
        primary_postgres_id=pulumi.get(__ret__, 'primary_postgres_id'),
        read_replicas=pulumi.get(__ret__, 'read_replicas'),
        region=pulumi.get(__ret__, 'region'),
        role=pulumi.get(__ret__, 'role'),
        status=pulumi.get(__ret__, 'status'),
        suspended=pulumi.get(__ret__, 'suspended'),
        suspenders=pulumi.get(__ret__, 'suspenders'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_postgres)
def get_postgres_output(postgres_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[PostgresDetail]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
