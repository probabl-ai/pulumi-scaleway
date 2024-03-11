# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MnqCredentialArgs', 'MnqCredential']

@pulumi.input_type
class MnqCredentialArgs:
    def __init__(__self__, *,
                 namespace_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 nats_credentials: Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs_sns_credentials: Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']] = None):
        """
        The set of arguments for constructing a MnqCredential resource.
        :param pulumi.Input[str] namespace_id: The namespace containing the Credential.
        :param pulumi.Input[str] name: The credential name..
        :param pulumi.Input['MnqCredentialNatsCredentialsArgs'] nats_credentials: Credentials file used to connect to the NATS service.
        :param pulumi.Input[str] region: (Defaults to provider `region`). The region
               in which the namespace should be created.
        :param pulumi.Input['MnqCredentialSqsSnsCredentialsArgs'] sqs_sns_credentials: Credential used to connect to the SQS/SNS service.
        """
        pulumi.set(__self__, "namespace_id", namespace_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nats_credentials is not None:
            pulumi.set(__self__, "nats_credentials", nats_credentials)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs_sns_credentials is not None:
            pulumi.set(__self__, "sqs_sns_credentials", sqs_sns_credentials)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Input[str]:
        """
        The namespace containing the Credential.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The credential name..
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="natsCredentials")
    def nats_credentials(self) -> Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']]:
        """
        Credentials file used to connect to the NATS service.
        """
        return pulumi.get(self, "nats_credentials")

    @nats_credentials.setter
    def nats_credentials(self, value: Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']]):
        pulumi.set(self, "nats_credentials", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        (Defaults to provider `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sqsSnsCredentials")
    def sqs_sns_credentials(self) -> Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']]:
        """
        Credential used to connect to the SQS/SNS service.
        """
        return pulumi.get(self, "sqs_sns_credentials")

    @sqs_sns_credentials.setter
    def sqs_sns_credentials(self, value: Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']]):
        pulumi.set(self, "sqs_sns_credentials", value)


@pulumi.input_type
class _MnqCredentialState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats_credentials: Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs_sns_credentials: Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']] = None):
        """
        Input properties used for looking up and filtering MnqCredential resources.
        :param pulumi.Input[str] name: The credential name..
        :param pulumi.Input[str] namespace_id: The namespace containing the Credential.
        :param pulumi.Input['MnqCredentialNatsCredentialsArgs'] nats_credentials: Credentials file used to connect to the NATS service.
        :param pulumi.Input[str] protocol: The protocol associated to the Credential. Possible values are `nats` and `sqs_sns`.
        :param pulumi.Input[str] region: (Defaults to provider `region`). The region
               in which the namespace should be created.
        :param pulumi.Input['MnqCredentialSqsSnsCredentialsArgs'] sqs_sns_credentials: Credential used to connect to the SQS/SNS service.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if nats_credentials is not None:
            pulumi.set(__self__, "nats_credentials", nats_credentials)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs_sns_credentials is not None:
            pulumi.set(__self__, "sqs_sns_credentials", sqs_sns_credentials)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The credential name..
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace containing the Credential.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="natsCredentials")
    def nats_credentials(self) -> Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']]:
        """
        Credentials file used to connect to the NATS service.
        """
        return pulumi.get(self, "nats_credentials")

    @nats_credentials.setter
    def nats_credentials(self, value: Optional[pulumi.Input['MnqCredentialNatsCredentialsArgs']]):
        pulumi.set(self, "nats_credentials", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol associated to the Credential. Possible values are `nats` and `sqs_sns`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        (Defaults to provider `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sqsSnsCredentials")
    def sqs_sns_credentials(self) -> Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']]:
        """
        Credential used to connect to the SQS/SNS service.
        """
        return pulumi.get(self, "sqs_sns_credentials")

    @sqs_sns_credentials.setter
    def sqs_sns_credentials(self, value: Optional[pulumi.Input['MnqCredentialSqsSnsCredentialsArgs']]):
        pulumi.set(self, "sqs_sns_credentials", value)


class MnqCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialNatsCredentialsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs_sns_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialSqsSnsCredentialsArgs']]] = None,
                 __props__=None):
        """
        ## Import

        Credential can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqCredential:MnqCredential main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The credential name..
        :param pulumi.Input[str] namespace_id: The namespace containing the Credential.
        :param pulumi.Input[pulumi.InputType['MnqCredentialNatsCredentialsArgs']] nats_credentials: Credentials file used to connect to the NATS service.
        :param pulumi.Input[str] region: (Defaults to provider `region`). The region
               in which the namespace should be created.
        :param pulumi.Input[pulumi.InputType['MnqCredentialSqsSnsCredentialsArgs']] sqs_sns_credentials: Credential used to connect to the SQS/SNS service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Credential can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqCredential:MnqCredential main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialNatsCredentialsArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs_sns_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialSqsSnsCredentialsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqCredentialArgs.__new__(MnqCredentialArgs)

            __props__.__dict__["name"] = name
            if namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_id'")
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["nats_credentials"] = nats_credentials
            __props__.__dict__["region"] = region
            __props__.__dict__["sqs_sns_credentials"] = sqs_sns_credentials
            __props__.__dict__["protocol"] = None
        super(MnqCredential, __self__).__init__(
            'scaleway:index/mnqCredential:MnqCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            nats_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialNatsCredentialsArgs']]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            sqs_sns_credentials: Optional[pulumi.Input[pulumi.InputType['MnqCredentialSqsSnsCredentialsArgs']]] = None) -> 'MnqCredential':
        """
        Get an existing MnqCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The credential name..
        :param pulumi.Input[str] namespace_id: The namespace containing the Credential.
        :param pulumi.Input[pulumi.InputType['MnqCredentialNatsCredentialsArgs']] nats_credentials: Credentials file used to connect to the NATS service.
        :param pulumi.Input[str] protocol: The protocol associated to the Credential. Possible values are `nats` and `sqs_sns`.
        :param pulumi.Input[str] region: (Defaults to provider `region`). The region
               in which the namespace should be created.
        :param pulumi.Input[pulumi.InputType['MnqCredentialSqsSnsCredentialsArgs']] sqs_sns_credentials: Credential used to connect to the SQS/SNS service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqCredentialState.__new__(_MnqCredentialState)

        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["nats_credentials"] = nats_credentials
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["sqs_sns_credentials"] = sqs_sns_credentials
        return MnqCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The credential name..
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        The namespace containing the Credential.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="natsCredentials")
    def nats_credentials(self) -> pulumi.Output['outputs.MnqCredentialNatsCredentials']:
        """
        Credentials file used to connect to the NATS service.
        """
        return pulumi.get(self, "nats_credentials")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol associated to the Credential. Possible values are `nats` and `sqs_sns`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        (Defaults to provider `region`). The region
        in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sqsSnsCredentials")
    def sqs_sns_credentials(self) -> pulumi.Output[Optional['outputs.MnqCredentialSqsSnsCredentials']]:
        """
        Credential used to connect to the SQS/SNS service.
        """
        return pulumi.get(self, "sqs_sns_credentials")

