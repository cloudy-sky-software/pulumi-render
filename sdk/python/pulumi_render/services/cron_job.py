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

__all__ = ['CronJobArgs', 'CronJob']

@pulumi.input_type
class CronJobArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 owner_id: pulumi.Input[str],
                 auto_deploy: Optional[pulumi.Input['ServiceCreateAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 build_filter: Optional[pulumi.Input['BuildFilterArgs']] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarInputArgs']]]] = None,
                 image: Optional[pulumi.Input['ImageArgs']] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 root_dir: Optional[pulumi.Input[str]] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]] = None,
                 service_details: Optional[pulumi.Input['CronJobDetailsCreateArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CronJob resource.
        :param pulumi.Input[str] branch: If left empty, this will fall back to the default branch of the repository
        :param pulumi.Input[str] repo: Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner_id", owner_id)
        if auto_deploy is None:
            auto_deploy = 'yes'
        if auto_deploy is not None:
            pulumi.set(__self__, "auto_deploy", auto_deploy)
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if build_filter is not None:
            pulumi.set(__self__, "build_filter", build_filter)
        if env_vars is not None:
            pulumi.set(__self__, "env_vars", env_vars)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)
        if root_dir is not None:
            pulumi.set(__self__, "root_dir", root_dir)
        if secret_files is not None:
            pulumi.set(__self__, "secret_files", secret_files)
        if service_details is not None:
            pulumi.set(__self__, "service_details", service_details)
        if type is None:
            type = 'cron_job'
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> Optional[pulumi.Input['ServiceCreateAutoDeploy']]:
        return pulumi.get(self, "auto_deploy")

    @auto_deploy.setter
    def auto_deploy(self, value: Optional[pulumi.Input['ServiceCreateAutoDeploy']]):
        pulumi.set(self, "auto_deploy", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        If left empty, this will fall back to the default branch of the repository
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="buildFilter")
    def build_filter(self) -> Optional[pulumi.Input['BuildFilterArgs']]:
        return pulumi.get(self, "build_filter")

    @build_filter.setter
    def build_filter(self, value: Optional[pulumi.Input['BuildFilterArgs']]):
        pulumi.set(self, "build_filter", value)

    @property
    @pulumi.getter(name="envVars")
    def env_vars(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarInputArgs']]]]:
        return pulumi.get(self, "env_vars")

    @env_vars.setter
    def env_vars(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvVarInputArgs']]]]):
        pulumi.set(self, "env_vars", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input['ImageArgs']]:
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input['ImageArgs']]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input[str]]:
        """
        Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo", value)

    @property
    @pulumi.getter(name="rootDir")
    def root_dir(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "root_dir")

    @root_dir.setter
    def root_dir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_dir", value)

    @property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]]:
        return pulumi.get(self, "secret_files")

    @secret_files.setter
    def secret_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretFileInputArgs']]]]):
        pulumi.set(self, "secret_files", value)

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> Optional[pulumi.Input['CronJobDetailsCreateArgs']]:
        return pulumi.get(self, "service_details")

    @service_details.setter
    def service_details(self, value: Optional[pulumi.Input['CronJobDetailsCreateArgs']]):
        pulumi.set(self, "service_details", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class CronJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_deploy: Optional[pulumi.Input['ServiceCreateAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 build_filter: Optional[pulumi.Input[Union['BuildFilterArgs', 'BuildFilterArgsDict']]] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvVarInputArgs', 'EnvVarInputArgsDict']]]]] = None,
                 image: Optional[pulumi.Input[Union['ImageArgs', 'ImageArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 root_dir: Optional[pulumi.Input[str]] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretFileInputArgs', 'SecretFileInputArgsDict']]]]] = None,
                 service_details: Optional[pulumi.Input[Union['CronJobDetailsCreateArgs', 'CronJobDetailsCreateArgsDict']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CronJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: If left empty, this will fall back to the default branch of the repository
        :param pulumi.Input[str] repo: Do not include the branch in the repo string. You can instead supply a 'branch' parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CronJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CronJob resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CronJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CronJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_deploy: Optional[pulumi.Input['ServiceCreateAutoDeploy']] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 build_filter: Optional[pulumi.Input[Union['BuildFilterArgs', 'BuildFilterArgsDict']]] = None,
                 env_vars: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EnvVarInputArgs', 'EnvVarInputArgsDict']]]]] = None,
                 image: Optional[pulumi.Input[Union['ImageArgs', 'ImageArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 root_dir: Optional[pulumi.Input[str]] = None,
                 secret_files: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretFileInputArgs', 'SecretFileInputArgsDict']]]]] = None,
                 service_details: Optional[pulumi.Input[Union['CronJobDetailsCreateArgs', 'CronJobDetailsCreateArgsDict']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CronJobArgs.__new__(CronJobArgs)

            if auto_deploy is None:
                auto_deploy = 'yes'
            __props__.__dict__["auto_deploy"] = auto_deploy
            __props__.__dict__["branch"] = branch
            __props__.__dict__["build_filter"] = build_filter
            __props__.__dict__["env_vars"] = env_vars
            __props__.__dict__["image"] = image
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'owner_id'")
            __props__.__dict__["owner_id"] = owner_id
            __props__.__dict__["repo"] = repo
            __props__.__dict__["root_dir"] = root_dir
            __props__.__dict__["secret_files"] = secret_files
            __props__.__dict__["service_details"] = service_details
            if type is None:
                type = 'cron_job'
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["dashboard_url"] = None
            __props__.__dict__["environment_id"] = None
            __props__.__dict__["image_path"] = None
            __props__.__dict__["notify_on_fail"] = None
            __props__.__dict__["registry_credential"] = None
            __props__.__dict__["slug"] = None
            __props__.__dict__["suspended"] = None
            __props__.__dict__["suspenders"] = None
            __props__.__dict__["updated_at"] = None
        super(CronJob, __self__).__init__(
            'render:services:CronJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CronJob':
        """
        Get an existing CronJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CronJobArgs.__new__(CronJobArgs)

        __props__.__dict__["auto_deploy"] = None
        __props__.__dict__["branch"] = None
        __props__.__dict__["build_filter"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["dashboard_url"] = None
        __props__.__dict__["env_vars"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["image"] = None
        __props__.__dict__["image_path"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notify_on_fail"] = None
        __props__.__dict__["owner_id"] = None
        __props__.__dict__["registry_credential"] = None
        __props__.__dict__["repo"] = None
        __props__.__dict__["root_dir"] = None
        __props__.__dict__["secret_files"] = None
        __props__.__dict__["service_details"] = None
        __props__.__dict__["slug"] = None
        __props__.__dict__["suspended"] = None
        __props__.__dict__["suspenders"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return CronJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoDeploy")
    def auto_deploy(self) -> pulumi.Output[Optional['ServiceAutoDeploy']]:
        return pulumi.get(self, "auto_deploy")

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="buildFilter")
    def build_filter(self) -> pulumi.Output[Optional['outputs.BuildFilter']]:
        return pulumi.get(self, "build_filter")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dashboardUrl")
    def dashboard_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL to view the service in the Render Dashboard
        """
        return pulumi.get(self, "dashboard_url")

    @property
    @pulumi.getter(name="envVars")
    def env_vars(self) -> pulumi.Output[Optional[Sequence['outputs.EnvVarInput']]]:
        return pulumi.get(self, "env_vars")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output[Optional['outputs.Image']]:
        return pulumi.get(self, "image")

    @property
    @pulumi.getter(name="imagePath")
    def image_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyOnFail")
    def notify_on_fail(self) -> pulumi.Output[Optional['ServiceNotifyOnFail']]:
        return pulumi.get(self, "notify_on_fail")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="registryCredential")
    def registry_credential(self) -> pulumi.Output[Optional['outputs.RegistryCredentialSummary']]:
        return pulumi.get(self, "registry_credential")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter(name="rootDir")
    def root_dir(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "root_dir")

    @property
    @pulumi.getter(name="secretFiles")
    def secret_files(self) -> pulumi.Output[Optional[Sequence['outputs.SecretFileInput']]]:
        return pulumi.get(self, "secret_files")

    @property
    @pulumi.getter(name="serviceDetails")
    def service_details(self) -> pulumi.Output[Optional['outputs.CronJobDetailsOutput']]:
        return pulumi.get(self, "service_details")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def suspended(self) -> pulumi.Output[Optional['ServiceSuspended']]:
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def suspenders(self) -> pulumi.Output[Optional[Sequence['ServiceSuspendersItem']]]:
        return pulumi.get(self, "suspenders")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "updated_at")

