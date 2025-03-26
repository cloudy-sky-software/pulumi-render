# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *
from ._inputs import *

__all__ = ['PostgresArgs', 'Postgres']

@pulumi.input_type
class PostgresArgs:
    def __init__(__self__, *,
                 owner_id: pulumi.Input[str],
                 plan: pulumi.Input['Plan'],
                 version: pulumi.Input['Version'],
                 database_name: Optional[pulumi.Input[str]] = None,
                 database_user: Optional[pulumi.Input[str]] = None,
                 datadog_api_key: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 enable_high_availability: Optional[pulumi.Input[bool]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 ip_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input['CidrBlockAndDescriptionArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_replicas: Optional[pulumi.Input[Sequence[pulumi.Input['ReadReplicaInputArgs']]]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Postgres resource.
        :param pulumi.Input[str] owner_id: The ID of the owner (team or personal user) whose resources should be returned
        :param pulumi.Input['Version'] version: The PostgreSQL version
        :param pulumi.Input[int] disk_size_gb: The number of gigabytes of disk space to allocate for the database
        :param pulumi.Input[str] name: The name of the database as it will appear in the Render Dashboard
        """
        pulumi.set(__self__, "owner_id", owner_id)
        pulumi.set(__self__, "plan", plan)
        pulumi.set(__self__, "version", version)
        if database_name is None:
            database_name = 'randomly generated'
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if database_user is None:
            database_user = 'randomly generated'
        if database_user is not None:
            pulumi.set(__self__, "database_user", database_user)
        if datadog_api_key is not None:
            pulumi.set(__self__, "datadog_api_key", datadog_api_key)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if enable_high_availability is None:
            enable_high_availability = False
        if enable_high_availability is not None:
            pulumi.set(__self__, "enable_high_availability", enable_high_availability)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if ip_allow_list is not None:
            pulumi.set(__self__, "ip_allow_list", ip_allow_list)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if read_replicas is not None:
            pulumi.set(__self__, "read_replicas", read_replicas)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Input[str]:
        """
        The ID of the owner (team or personal user) whose resources should be returned
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Input['Plan']:
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: pulumi.Input['Plan']):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input['Version']:
        """
        The PostgreSQL version
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input['Version']):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="databaseUser")
    def database_user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_user")

    @database_user.setter
    def database_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_user", value)

    @property
    @pulumi.getter(name="datadogAPIKey")
    def datadog_api_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "datadog_api_key")

    @datadog_api_key.setter
    def datadog_api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datadog_api_key", value)

    @property
    @pulumi.getter(name="diskSizeGB")
    def disk_size_gb(self) -> Optional[pulumi.Input[int]]:
        """
        The number of gigabytes of disk space to allocate for the database
        """
        return pulumi.get(self, "disk_size_gb")

    @disk_size_gb.setter
    def disk_size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_size_gb", value)

    @property
    @pulumi.getter(name="enableHighAvailability")
    def enable_high_availability(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_high_availability")

    @enable_high_availability.setter
    def enable_high_availability(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_high_availability", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="ipAllowList")
    def ip_allow_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CidrBlockAndDescriptionArgs']]]]:
        return pulumi.get(self, "ip_allow_list")

    @ip_allow_list.setter
    def ip_allow_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CidrBlockAndDescriptionArgs']]]]):
        pulumi.set(self, "ip_allow_list", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database as it will appear in the Render Dashboard
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="readReplicas")
    def read_replicas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReadReplicaInputArgs']]]]:
        return pulumi.get(self, "read_replicas")

    @read_replicas.setter
    def read_replicas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReadReplicaInputArgs']]]]):
        pulumi.set(self, "read_replicas", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Postgres(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 database_user: Optional[pulumi.Input[str]] = None,
                 datadog_api_key: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 enable_high_availability: Optional[pulumi.Input[bool]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 ip_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CidrBlockAndDescriptionArgs', 'CidrBlockAndDescriptionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['Plan']] = None,
                 read_replicas: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ReadReplicaInputArgs', 'ReadReplicaInputArgsDict']]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input['Version']] = None,
                 __props__=None):
        """
        Input for creating a database

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] disk_size_gb: The number of gigabytes of disk space to allocate for the database
        :param pulumi.Input[str] name: The name of the database as it will appear in the Render Dashboard
        :param pulumi.Input[str] owner_id: The ID of the owner (team or personal user) whose resources should be returned
        :param pulumi.Input['Version'] version: The PostgreSQL version
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PostgresArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Input for creating a database

        :param str resource_name: The name of the resource.
        :param PostgresArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PostgresArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 database_user: Optional[pulumi.Input[str]] = None,
                 datadog_api_key: Optional[pulumi.Input[str]] = None,
                 disk_size_gb: Optional[pulumi.Input[int]] = None,
                 enable_high_availability: Optional[pulumi.Input[bool]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 ip_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CidrBlockAndDescriptionArgs', 'CidrBlockAndDescriptionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['Plan']] = None,
                 read_replicas: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ReadReplicaInputArgs', 'ReadReplicaInputArgsDict']]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input['Version']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PostgresArgs.__new__(PostgresArgs)

            if database_name is None:
                database_name = 'randomly generated'
            __props__.__dict__["database_name"] = database_name
            if database_user is None:
                database_user = 'randomly generated'
            __props__.__dict__["database_user"] = database_user
            __props__.__dict__["datadog_api_key"] = datadog_api_key
            __props__.__dict__["disk_size_gb"] = disk_size_gb
            if enable_high_availability is None:
                enable_high_availability = False
            __props__.__dict__["enable_high_availability"] = enable_high_availability
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["ip_allow_list"] = ip_allow_list
            __props__.__dict__["name"] = name
            if owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'owner_id'")
            __props__.__dict__["owner_id"] = owner_id
            if plan is None and not opts.urn:
                raise TypeError("Missing required property 'plan'")
            __props__.__dict__["plan"] = plan
            __props__.__dict__["read_replicas"] = read_replicas
            __props__.__dict__["region"] = region
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["created_at"] = None
            __props__.__dict__["dashboard_url"] = None
            __props__.__dict__["expires_at"] = None
            __props__.__dict__["high_availability_enabled"] = None
            __props__.__dict__["maintenance"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["primary_postgres_id"] = None
            __props__.__dict__["role"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["suspended"] = None
            __props__.__dict__["suspenders"] = None
            __props__.__dict__["updated_at"] = None
        super(Postgres, __self__).__init__(
            'render:postgres:Postgres',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Postgres':
        """
        Get an existing Postgres resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PostgresArgs.__new__(PostgresArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["dashboard_url"] = None
        __props__.__dict__["database_name"] = None
        __props__.__dict__["database_user"] = None
        __props__.__dict__["datadog_api_key"] = None
        __props__.__dict__["disk_size_gb"] = None
        __props__.__dict__["enable_high_availability"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["expires_at"] = None
        __props__.__dict__["high_availability_enabled"] = None
        __props__.__dict__["ip_allow_list"] = None
        __props__.__dict__["maintenance"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["plan"] = None
        __props__.__dict__["primary_postgres_id"] = None
        __props__.__dict__["read_replicas"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["suspended"] = None
        __props__.__dict__["suspenders"] = None
        __props__.__dict__["updated_at"] = None
        __props__.__dict__["version"] = None
        return Postgres(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dashboardUrl")
    def dashboard_url(self) -> pulumi.Output[str]:
        """
        The URL to view the Postgres instance in the Render Dashboard
        """
        return pulumi.get(self, "dashboard_url")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="databaseUser")
    def database_user(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_user")

    @property
    @pulumi.getter(name="datadogAPIKey")
    def datadog_api_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "datadog_api_key")

    @property
    @pulumi.getter(name="diskSizeGB")
    def disk_size_gb(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter(name="enableHighAvailability")
    def enable_high_availability(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_high_availability")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        The time at which the database will be expire. Applies to free tier databases only.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="highAvailabilityEnabled")
    def high_availability_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "high_availability_enabled")

    @property
    @pulumi.getter(name="ipAllowList")
    def ip_allow_list(self) -> pulumi.Output[Sequence['outputs.CidrBlockAndDescription']]:
        return pulumi.get(self, "ip_allow_list")

    @property
    @pulumi.getter
    def maintenance(self) -> pulumi.Output[Optional['outputs.RedisDetailpropertiesmaintenance']]:
        return pulumi.get(self, "maintenance")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output['outputs.Owner']:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The ID of the owner (team or personal user) whose resources should be returned
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output['Plan']:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="primaryPostgresID")
    def primary_postgres_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "primary_postgres_id")

    @property
    @pulumi.getter(name="readReplicas")
    def read_replicas(self) -> pulumi.Output[Sequence['outputs.ReadReplica']]:
        return pulumi.get(self, "read_replicas")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output['Region']:
        """
        Defaults to "oregon"
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output['Role']:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['Status']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def suspended(self) -> pulumi.Output['Suspended']:
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def suspenders(self) -> pulumi.Output[Sequence['SuspendersItem']]:
        return pulumi.get(self, "suspenders")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output['Version']:
        """
        The PostgreSQL version
        """
        return pulumi.get(self, "version")

