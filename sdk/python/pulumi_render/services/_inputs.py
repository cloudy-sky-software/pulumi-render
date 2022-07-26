# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'ServerPropertiesArgs',
    'StaticSiteServiceHeaderArgs',
    'StaticSiteStaticSiteRouteArgs',
    'StaticSiteStaticSiteServiceDetailsParentServerPropertiesArgs',
    'StaticSiteStaticSiteServiceDetailsArgs',
    'WebServiceDiskArgs',
    'WebServiceDockerDetailsArgs',
    'WebServiceNativeEnvironmentDetailsArgs',
    'WebServiceWebServiceServiceDetailsArgs',
]

@pulumi.input_type
class ServerPropertiesArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class StaticSiteServiceHeaderArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 path: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A service header object
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class StaticSiteStaticSiteRouteArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 source: pulumi.Input[str],
                 type: pulumi.Input['StaticSiteStaticSiteRouteType']):
        """
        A route object for a static site
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['StaticSiteStaticSiteRouteType']:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['StaticSiteStaticSiteRouteType']):
        pulumi.set(self, "type", value)


@pulumi.input_type
class StaticSiteStaticSiteServiceDetailsParentServerPropertiesArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class StaticSiteStaticSiteServiceDetailsArgs:
    def __init__(__self__, *,
                 build_command: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteServiceHeaderArgs']]]] = None,
                 parent_server: Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsParentServerPropertiesArgs']] = None,
                 publish_path: Optional[pulumi.Input[str]] = None,
                 pull_request_previews_enabled: Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled']] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteStaticSiteRouteArgs']]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] url: The HTTPS service URL. A subdomain of onrender.com, by default.
        """
        if build_command is not None:
            pulumi.set(__self__, "build_command", build_command)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if parent_server is not None:
            pulumi.set(__self__, "parent_server", parent_server)
        if publish_path is None:
            publish_path = 'public'
        if publish_path is not None:
            pulumi.set(__self__, "publish_path", publish_path)
        if pull_request_previews_enabled is None:
            pull_request_previews_enabled = 'no'
        if pull_request_previews_enabled is not None:
            pulumi.set(__self__, "pull_request_previews_enabled", pull_request_previews_enabled)
        if routes is not None:
            pulumi.set(__self__, "routes", routes)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="buildCommand")
    def build_command(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "build_command")

    @build_command.setter
    def build_command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "build_command", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteServiceHeaderArgs']]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteServiceHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="parentServer")
    def parent_server(self) -> Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsParentServerPropertiesArgs']]:
        return pulumi.get(self, "parent_server")

    @parent_server.setter
    def parent_server(self, value: Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsParentServerPropertiesArgs']]):
        pulumi.set(self, "parent_server", value)

    @property
    @pulumi.getter(name="publishPath")
    def publish_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "publish_path")

    @publish_path.setter
    def publish_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publish_path", value)

    @property
    @pulumi.getter(name="pullRequestPreviewsEnabled")
    def pull_request_previews_enabled(self) -> Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled']]:
        return pulumi.get(self, "pull_request_previews_enabled")

    @pull_request_previews_enabled.setter
    def pull_request_previews_enabled(self, value: Optional[pulumi.Input['StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled']]):
        pulumi.set(self, "pull_request_previews_enabled", value)

    @property
    @pulumi.getter
    def routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteStaticSiteRouteArgs']]]]:
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StaticSiteStaticSiteRouteArgs']]]]):
        pulumi.set(self, "routes", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTPS service URL. A subdomain of onrender.com, by default.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class WebServiceDiskArgs:
    def __init__(__self__, *,
                 mount_path: pulumi.Input[str],
                 name: pulumi.Input[str],
                 size_gb: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "mount_path", mount_path)
        pulumi.set(__self__, "name", name)
        if size_gb is None:
            size_gb = 1
        if size_gb is not None:
            pulumi.set(__self__, "size_gb", size_gb)

    @property
    @pulumi.getter(name="mountPath")
    def mount_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "mount_path")

    @mount_path.setter
    def mount_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount_path", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sizeGB")
    def size_gb(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "size_gb")

    @size_gb.setter
    def size_gb(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "size_gb", value)


@pulumi.input_type
class WebServiceDockerDetailsArgs:
    def __init__(__self__, *,
                 docker_command: Optional[pulumi.Input[str]] = None,
                 docker_context: Optional[pulumi.Input[str]] = None,
                 dockerfile_path: Optional[pulumi.Input[str]] = None):
        if docker_command is not None:
            pulumi.set(__self__, "docker_command", docker_command)
        if docker_context is not None:
            pulumi.set(__self__, "docker_context", docker_context)
        if dockerfile_path is None:
            dockerfile_path = './Dockerfile'
        if dockerfile_path is not None:
            pulumi.set(__self__, "dockerfile_path", dockerfile_path)

    @property
    @pulumi.getter(name="dockerCommand")
    def docker_command(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "docker_command")

    @docker_command.setter
    def docker_command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "docker_command", value)

    @property
    @pulumi.getter(name="dockerContext")
    def docker_context(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "docker_context")

    @docker_context.setter
    def docker_context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "docker_context", value)

    @property
    @pulumi.getter(name="dockerfilePath")
    def dockerfile_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dockerfile_path")

    @dockerfile_path.setter
    def dockerfile_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dockerfile_path", value)


@pulumi.input_type
class WebServiceNativeEnvironmentDetailsArgs:
    def __init__(__self__, *,
                 build_command: pulumi.Input[str],
                 start_command: pulumi.Input[str]):
        pulumi.set(__self__, "build_command", build_command)
        pulumi.set(__self__, "start_command", start_command)

    @property
    @pulumi.getter(name="buildCommand")
    def build_command(self) -> pulumi.Input[str]:
        return pulumi.get(self, "build_command")

    @build_command.setter
    def build_command(self, value: pulumi.Input[str]):
        pulumi.set(self, "build_command", value)

    @property
    @pulumi.getter(name="startCommand")
    def start_command(self) -> pulumi.Input[str]:
        return pulumi.get(self, "start_command")

    @start_command.setter
    def start_command(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_command", value)


@pulumi.input_type
class WebServiceWebServiceServiceDetailsArgs:
    def __init__(__self__, *,
                 env: pulumi.Input['WebServiceWebServiceServiceDetailsEnv'],
                 disk: Optional[pulumi.Input['WebServiceDiskArgs']] = None,
                 env_specific_details: Optional[pulumi.Input[Union['WebServiceDockerDetailsArgs', 'WebServiceNativeEnvironmentDetailsArgs']]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 num_instances: Optional[pulumi.Input[float]] = None,
                 plan: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPlan']] = None,
                 pull_request_previews_enabled: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled']] = None,
                 region: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsRegion']] = None):
        pulumi.set(__self__, "env", env)
        if disk is not None:
            pulumi.set(__self__, "disk", disk)
        if env_specific_details is not None:
            pulumi.set(__self__, "env_specific_details", env_specific_details)
        if health_check_path is not None:
            pulumi.set(__self__, "health_check_path", health_check_path)
        if num_instances is None:
            num_instances = 1
        if num_instances is not None:
            pulumi.set(__self__, "num_instances", num_instances)
        if plan is None:
            plan = 'starter'
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if pull_request_previews_enabled is None:
            pull_request_previews_enabled = 'no'
        if pull_request_previews_enabled is not None:
            pulumi.set(__self__, "pull_request_previews_enabled", pull_request_previews_enabled)
        if region is None:
            region = 'oregon'
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def env(self) -> pulumi.Input['WebServiceWebServiceServiceDetailsEnv']:
        return pulumi.get(self, "env")

    @env.setter
    def env(self, value: pulumi.Input['WebServiceWebServiceServiceDetailsEnv']):
        pulumi.set(self, "env", value)

    @property
    @pulumi.getter
    def disk(self) -> Optional[pulumi.Input['WebServiceDiskArgs']]:
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: Optional[pulumi.Input['WebServiceDiskArgs']]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter(name="envSpecificDetails")
    def env_specific_details(self) -> Optional[pulumi.Input[Union['WebServiceDockerDetailsArgs', 'WebServiceNativeEnvironmentDetailsArgs']]]:
        return pulumi.get(self, "env_specific_details")

    @env_specific_details.setter
    def env_specific_details(self, value: Optional[pulumi.Input[Union['WebServiceDockerDetailsArgs', 'WebServiceNativeEnvironmentDetailsArgs']]]):
        pulumi.set(self, "env_specific_details", value)

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_path")

    @health_check_path.setter
    def health_check_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_path", value)

    @property
    @pulumi.getter(name="numInstances")
    def num_instances(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "num_instances")

    @num_instances.setter
    def num_instances(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "num_instances", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPlan']]:
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPlan']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="pullRequestPreviewsEnabled")
    def pull_request_previews_enabled(self) -> Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled']]:
        return pulumi.get(self, "pull_request_previews_enabled")

    @pull_request_previews_enabled.setter
    def pull_request_previews_enabled(self, value: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled']]):
        pulumi.set(self, "pull_request_previews_enabled", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input['WebServiceWebServiceServiceDetailsRegion']]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input['WebServiceWebServiceServiceDetailsRegion']]):
        pulumi.set(self, "region", value)

