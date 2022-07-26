# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['CustomDomainArgs', 'CustomDomain']

@pulumi.input_type
class CustomDomainArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 server: pulumi.Input['ServerPropertiesArgs']):
        """
        The set of arguments for constructing a CustomDomain resource.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "server", server)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def server(self) -> pulumi.Input['ServerPropertiesArgs']:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: pulumi.Input['ServerPropertiesArgs']):
        pulumi.set(self, "server", value)


class CustomDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[pulumi.InputType['ServerPropertiesArgs']]] = None,
                 __props__=None):
        """
        Create a CustomDomain resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CustomDomain resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CustomDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[pulumi.InputType['ServerPropertiesArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomDomainArgs.__new__(CustomDomainArgs)

            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if server is None and not opts.urn:
                raise TypeError("Missing required property 'server'")
            __props__.__dict__["server"] = server
            __props__.__dict__["created_at"] = None
            __props__.__dict__["domain_type"] = None
            __props__.__dict__["public_suffix"] = None
            __props__.__dict__["redirect_for_name"] = None
            __props__.__dict__["verification_status"] = None
        super(CustomDomain, __self__).__init__(
            'render:services:CustomDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CustomDomain':
        """
        Get an existing CustomDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CustomDomainArgs.__new__(CustomDomainArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["domain_type"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["public_suffix"] = None
        __props__.__dict__["redirect_for_name"] = None
        __props__.__dict__["server"] = None
        __props__.__dict__["verification_status"] = None
        return CustomDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="domainType")
    def domain_type(self) -> pulumi.Output['CustomDomainDomainType']:
        return pulumi.get(self, "domain_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicSuffix")
    def public_suffix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "public_suffix")

    @property
    @pulumi.getter(name="redirectForName")
    def redirect_for_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "redirect_for_name")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output['outputs.ServerProperties']:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="verificationStatus")
    def verification_status(self) -> pulumi.Output['CustomDomainVerificationStatus']:
        return pulumi.get(self, "verification_status")
