# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqSnsTopicArgs', 'MnqSnsTopic']

@pulumi.input_type
class MnqSnsTopicArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_topic: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqSnsTopic resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication.
        :param pulumi.Input[bool] fifo_topic: Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        :param pulumi.Input[str] name: The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "secret_key", secret_key)
        if content_based_deduplication is not None:
            pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if fifo_topic is not None:
            pulumi.set(__self__, "fifo_topic", fifo_topic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sns_endpoint is not None:
            pulumi.set(__self__, "sns_endpoint", sns_endpoint)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable content-based deduplication.
        """
        return pulumi.get(self, "content_based_deduplication")

    @content_based_deduplication.setter
    def content_based_deduplication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "content_based_deduplication", value)

    @property
    @pulumi.getter(name="fifoTopic")
    def fifo_topic(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        """
        return pulumi.get(self, "fifo_topic")

    @fifo_topic.setter
    def fifo_topic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fifo_topic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

    @sns_endpoint.setter
    def sns_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_endpoint", value)


@pulumi.input_type
class _MnqSnsTopicState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_topic: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqSnsTopic resources.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] arn: The ARN of the topic
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication.
        :param pulumi.Input[bool] fifo_topic: Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        :param pulumi.Input[str] name: The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] owner: Owner of the SNS topic, should have format 'project-${project_id}'
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if content_based_deduplication is not None:
            pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if fifo_topic is not None:
            pulumi.set(__self__, "fifo_topic", fifo_topic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sns_endpoint is not None:
            pulumi.set(__self__, "sns_endpoint", sns_endpoint)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the topic
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable content-based deduplication.
        """
        return pulumi.get(self, "content_based_deduplication")

    @content_based_deduplication.setter
    def content_based_deduplication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "content_based_deduplication", value)

    @property
    @pulumi.getter(name="fifoTopic")
    def fifo_topic(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        """
        return pulumi.get(self, "fifo_topic")

    @fifo_topic.setter
    def fifo_topic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fifo_topic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Owner of the SNS topic, should have format 'project-${project_id}'
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

    @sns_endpoint.setter
    def sns_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_endpoint", value)


class MnqSnsTopic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_topic: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage Scaleway Messaging and queuing SNS topics.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-topics/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions={
                "can_manage": True,
            })
        topic = scaleway.MnqSnsTopic("topic",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key)
        ```

        ## Import

        SNS topics can be imported using `{region}/{project-id}/{topic-name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsTopic:MnqSnsTopic main fr-par/11111111111111111111111111111111/my-topic
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication.
        :param pulumi.Input[bool] fifo_topic: Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        :param pulumi.Input[str] name: The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqSnsTopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage Scaleway Messaging and queuing SNS topics.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-topics/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions={
                "can_manage": True,
            })
        topic = scaleway.MnqSnsTopic("topic",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key)
        ```

        ## Import

        SNS topics can be imported using `{region}/{project-id}/{topic-name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsTopic:MnqSnsTopic main fr-par/11111111111111111111111111111111/my-topic
        ```

        :param str resource_name: The name of the resource.
        :param MnqSnsTopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqSnsTopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_topic: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqSnsTopicArgs.__new__(MnqSnsTopicArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = None if access_key is None else pulumi.Output.secret(access_key)
            __props__.__dict__["content_based_deduplication"] = content_based_deduplication
            __props__.__dict__["fifo_topic"] = fifo_topic
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["sns_endpoint"] = sns_endpoint
            __props__.__dict__["arn"] = None
            __props__.__dict__["owner"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MnqSnsTopic, __self__).__init__(
            'scaleway:index/mnqSnsTopic:MnqSnsTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            content_based_deduplication: Optional[pulumi.Input[bool]] = None,
            fifo_topic: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            sns_endpoint: Optional[pulumi.Input[str]] = None) -> 'MnqSnsTopic':
        """
        Get an existing MnqSnsTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] arn: The ARN of the topic
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication.
        :param pulumi.Input[bool] fifo_topic: Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        :param pulumi.Input[str] name: The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] owner: Owner of the SNS topic, should have format 'project-${project_id}'
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqSnsTopicState.__new__(_MnqSnsTopicState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["arn"] = arn
        __props__.__dict__["content_based_deduplication"] = content_based_deduplication
        __props__.__dict__["fifo_topic"] = fifo_topic
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["owner"] = owner
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["sns_endpoint"] = sns_endpoint
        return MnqSnsTopic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the topic
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> pulumi.Output[bool]:
        """
        Specifies whether to enable content-based deduplication.
        """
        return pulumi.get(self, "content_based_deduplication")

    @property
    @pulumi.getter(name="fifoTopic")
    def fifo_topic(self) -> pulumi.Output[bool]:
        """
        Whether the topic is a FIFO topic. If true, the topic name must end with .fifo.
        """
        return pulumi.get(self, "fifo_topic")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the SNS topic. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Owner of the SNS topic, should have format 'project-${project_id}'
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

