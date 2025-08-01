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
from ._inputs import *

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 environments: pulumi.Input[Sequence[pulumi.Input['ProjectCreateEnvironmentInputArgs']]],
                 owner_id: pulumi.Input[_builtins.str],
                 name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectCreateEnvironmentInputArgs']]] environments: The environments to create when creating the project
        :param pulumi.Input[_builtins.str] owner_id: The ID of the owner that the project belongs to
        :param pulumi.Input[_builtins.str] name: The name of the project
        """
        pulumi.set(__self__, "environments", environments)
        pulumi.set(__self__, "owner_id", owner_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @_builtins.property
    @pulumi.getter
    def environments(self) -> pulumi.Input[Sequence[pulumi.Input['ProjectCreateEnvironmentInputArgs']]]:
        """
        The environments to create when creating the project
        """
        return pulumi.get(self, "environments")

    @environments.setter
    def environments(self, value: pulumi.Input[Sequence[pulumi.Input['ProjectCreateEnvironmentInputArgs']]]):
        pulumi.set(self, "environments", value)

    @_builtins.property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the owner that the project belongs to
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "owner_id", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the project
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("render:projects:Project")
class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectCreateEnvironmentInputArgs', 'ProjectCreateEnvironmentInputArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 owner_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Create a Project resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectCreateEnvironmentInputArgs', 'ProjectCreateEnvironmentInputArgsDict']]]] environments: The environments to create when creating the project
        :param pulumi.Input[_builtins.str] name: The name of the project
        :param pulumi.Input[_builtins.str] owner_id: The ID of the owner that the project belongs to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Project resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environments: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectCreateEnvironmentInputArgs', 'ProjectCreateEnvironmentInputArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 owner_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            if environments is None and not opts.urn:
                raise TypeError("Missing required property 'environments'")
            __props__.__dict__["environments"] = environments
            __props__.__dict__["name"] = name
            if owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'owner_id'")
            __props__.__dict__["owner_id"] = owner_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["environment_ids"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["updated_at"] = None
        super(Project, __self__).__init__(
            'render:projects:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProjectArgs.__new__(ProjectArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["environment_ids"] = None
        __props__.__dict__["environments"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["updated_at"] = None
        return Project(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "created_at")

    @_builtins.property
    @pulumi.getter(name="environmentIds")
    def environment_ids(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        The environments associated with the project
        """
        return pulumi.get(self, "environment_ids")

    @_builtins.property
    @pulumi.getter
    def environments(self) -> pulumi.Output[Sequence['outputs.ProjectCreateEnvironmentInput']]:
        """
        The environments to create when creating the project
        """
        return pulumi.get(self, "environments")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the project
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def owner(self) -> pulumi.Output['outputs.Owner']:
        return pulumi.get(self, "owner")

    @_builtins.property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the owner that the project belongs to
        """
        return pulumi.get(self, "owner_id")

    @_builtins.property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "updated_at")

