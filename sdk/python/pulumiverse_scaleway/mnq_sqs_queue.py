# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqSqsQueueArgs', 'MnqSqsQueue']

@pulumi.input_type
class MnqSqsQueueArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_queue: Optional[pulumi.Input[bool]] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 receive_wait_time_seconds: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs_endpoint: Optional[pulumi.Input[str]] = None,
                 visibility_timeout_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a MnqSqsQueue resource.
        :param pulumi.Input[str] access_key: The access key of the SQS queue.
        :param pulumi.Input[str] secret_key: The secret key of the SQS queue.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication. Defaults to `false`.
        :param pulumi.Input[bool] fifo_queue: Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        :param pulumi.Input[str] name: The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SQS is enabled.
        :param pulumi.Input[int] receive_wait_time_seconds: The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        :param pulumi.Input[str] region: `region`). The region in which SQS is enabled.
        :param pulumi.Input[str] sqs_endpoint: The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        :param pulumi.Input[int] visibility_timeout_seconds: The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "secret_key", secret_key)
        if content_based_deduplication is not None:
            pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if fifo_queue is not None:
            pulumi.set(__self__, "fifo_queue", fifo_queue)
        if message_max_age is not None:
            pulumi.set(__self__, "message_max_age", message_max_age)
        if message_max_size is not None:
            pulumi.set(__self__, "message_max_size", message_max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if receive_wait_time_seconds is not None:
            pulumi.set(__self__, "receive_wait_time_seconds", receive_wait_time_seconds)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs_endpoint is not None:
            pulumi.set(__self__, "sqs_endpoint", sqs_endpoint)
        if visibility_timeout_seconds is not None:
            pulumi.set(__self__, "visibility_timeout_seconds", visibility_timeout_seconds)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The access key of the SQS queue.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The secret key of the SQS queue.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable content-based deduplication. Defaults to `false`.
        """
        return pulumi.get(self, "content_based_deduplication")

    @content_based_deduplication.setter
    def content_based_deduplication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "content_based_deduplication", value)

    @property
    @pulumi.getter(name="fifoQueue")
    def fifo_queue(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        """
        return pulumi.get(self, "fifo_queue")

    @fifo_queue.setter
    def fifo_queue(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fifo_queue", value)

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        """
        return pulumi.get(self, "message_max_age")

    @message_max_age.setter
    def message_max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_age", value)

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        """
        return pulumi.get(self, "message_max_size")

    @message_max_size.setter
    def message_max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
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
        `project_id`) The ID of the Project in which SQS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="receiveWaitTimeSeconds")
    def receive_wait_time_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        """
        return pulumi.get(self, "receive_wait_time_seconds")

    @receive_wait_time_seconds.setter
    def receive_wait_time_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "receive_wait_time_seconds", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which SQS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sqsEndpoint")
    def sqs_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sqs_endpoint")

    @sqs_endpoint.setter
    def sqs_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sqs_endpoint", value)

    @property
    @pulumi.getter(name="visibilityTimeoutSeconds")
    def visibility_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        return pulumi.get(self, "visibility_timeout_seconds")

    @visibility_timeout_seconds.setter
    def visibility_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "visibility_timeout_seconds", value)


@pulumi.input_type
class _MnqSqsQueueState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_queue: Optional[pulumi.Input[bool]] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 receive_wait_time_seconds: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sqs_endpoint: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 visibility_timeout_seconds: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering MnqSqsQueue resources.
        :param pulumi.Input[str] access_key: The access key of the SQS queue.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication. Defaults to `false`.
        :param pulumi.Input[bool] fifo_queue: Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        :param pulumi.Input[str] name: The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SQS is enabled.
        :param pulumi.Input[int] receive_wait_time_seconds: The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        :param pulumi.Input[str] region: `region`). The region in which SQS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SQS queue.
        :param pulumi.Input[str] sqs_endpoint: The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] url: The URL of the queue.
        :param pulumi.Input[int] visibility_timeout_seconds: The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if content_based_deduplication is not None:
            pulumi.set(__self__, "content_based_deduplication", content_based_deduplication)
        if fifo_queue is not None:
            pulumi.set(__self__, "fifo_queue", fifo_queue)
        if message_max_age is not None:
            pulumi.set(__self__, "message_max_age", message_max_age)
        if message_max_size is not None:
            pulumi.set(__self__, "message_max_size", message_max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if receive_wait_time_seconds is not None:
            pulumi.set(__self__, "receive_wait_time_seconds", receive_wait_time_seconds)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sqs_endpoint is not None:
            pulumi.set(__self__, "sqs_endpoint", sqs_endpoint)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if visibility_timeout_seconds is not None:
            pulumi.set(__self__, "visibility_timeout_seconds", visibility_timeout_seconds)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key of the SQS queue.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to enable content-based deduplication. Defaults to `false`.
        """
        return pulumi.get(self, "content_based_deduplication")

    @content_based_deduplication.setter
    def content_based_deduplication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "content_based_deduplication", value)

    @property
    @pulumi.getter(name="fifoQueue")
    def fifo_queue(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        """
        return pulumi.get(self, "fifo_queue")

    @fifo_queue.setter
    def fifo_queue(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fifo_queue", value)

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        """
        return pulumi.get(self, "message_max_age")

    @message_max_age.setter
    def message_max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_age", value)

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        """
        return pulumi.get(self, "message_max_size")

    @message_max_size.setter
    def message_max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
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
        `project_id`) The ID of the Project in which SQS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="receiveWaitTimeSeconds")
    def receive_wait_time_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        """
        return pulumi.get(self, "receive_wait_time_seconds")

    @receive_wait_time_seconds.setter
    def receive_wait_time_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "receive_wait_time_seconds", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which SQS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key of the SQS queue.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="sqsEndpoint")
    def sqs_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sqs_endpoint")

    @sqs_endpoint.setter
    def sqs_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sqs_endpoint", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the queue.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="visibilityTimeoutSeconds")
    def visibility_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        return pulumi.get(self, "visibility_timeout_seconds")

    @visibility_timeout_seconds.setter
    def visibility_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "visibility_timeout_seconds", value)


class MnqSqsQueue(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_queue: Optional[pulumi.Input[bool]] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 receive_wait_time_seconds: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sqs_endpoint: Optional[pulumi.Input[str]] = None,
                 visibility_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and Queuing SQS queues.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-queues/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_sqs = scaleway.MnqSqs("mainMnqSqs")
        main_mnq_sqs_credentials = scaleway.MnqSqsCredentials("mainMnqSqsCredentials",
            project_id=main_mnq_sqs.project_id,
            permissions=scaleway.MnqSqsCredentialsPermissionsArgs(
                can_manage=True,
                can_receive=False,
                can_publish=False,
            ))
        main_mnq_sqs_queue = scaleway.MnqSqsQueue("mainMnqSqsQueue",
            project_id=main_mnq_sqs.project_id,
            sqs_endpoint=main_mnq_sqs.endpoint,
            access_key=main_mnq_sqs_credentials.access_key,
            secret_key=main_mnq_sqs_credentials.secret_key)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SQS queue.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication. Defaults to `false`.
        :param pulumi.Input[bool] fifo_queue: Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        :param pulumi.Input[str] name: The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SQS is enabled.
        :param pulumi.Input[int] receive_wait_time_seconds: The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        :param pulumi.Input[str] region: `region`). The region in which SQS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SQS queue.
        :param pulumi.Input[str] sqs_endpoint: The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        :param pulumi.Input[int] visibility_timeout_seconds: The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqSqsQueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and Queuing SQS queues.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/how-to/create-manage-queues/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main_mnq_sqs = scaleway.MnqSqs("mainMnqSqs")
        main_mnq_sqs_credentials = scaleway.MnqSqsCredentials("mainMnqSqsCredentials",
            project_id=main_mnq_sqs.project_id,
            permissions=scaleway.MnqSqsCredentialsPermissionsArgs(
                can_manage=True,
                can_receive=False,
                can_publish=False,
            ))
        main_mnq_sqs_queue = scaleway.MnqSqsQueue("mainMnqSqsQueue",
            project_id=main_mnq_sqs.project_id,
            sqs_endpoint=main_mnq_sqs.endpoint,
            access_key=main_mnq_sqs_credentials.access_key,
            secret_key=main_mnq_sqs_credentials.secret_key)
        ```

        :param str resource_name: The name of the resource.
        :param MnqSqsQueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqSqsQueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 content_based_deduplication: Optional[pulumi.Input[bool]] = None,
                 fifo_queue: Optional[pulumi.Input[bool]] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 receive_wait_time_seconds: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sqs_endpoint: Optional[pulumi.Input[str]] = None,
                 visibility_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqSqsQueueArgs.__new__(MnqSqsQueueArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = None if access_key is None else pulumi.Output.secret(access_key)
            __props__.__dict__["content_based_deduplication"] = content_based_deduplication
            __props__.__dict__["fifo_queue"] = fifo_queue
            __props__.__dict__["message_max_age"] = message_max_age
            __props__.__dict__["message_max_size"] = message_max_size
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["receive_wait_time_seconds"] = receive_wait_time_seconds
            __props__.__dict__["region"] = region
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["sqs_endpoint"] = sqs_endpoint
            __props__.__dict__["visibility_timeout_seconds"] = visibility_timeout_seconds
            __props__.__dict__["url"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MnqSqsQueue, __self__).__init__(
            'scaleway:index/mnqSqsQueue:MnqSqsQueue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            content_based_deduplication: Optional[pulumi.Input[bool]] = None,
            fifo_queue: Optional[pulumi.Input[bool]] = None,
            message_max_age: Optional[pulumi.Input[int]] = None,
            message_max_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            receive_wait_time_seconds: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            sqs_endpoint: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            visibility_timeout_seconds: Optional[pulumi.Input[int]] = None) -> 'MnqSqsQueue':
        """
        Get an existing MnqSqsQueue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SQS queue.
        :param pulumi.Input[bool] content_based_deduplication: Specifies whether to enable content-based deduplication. Defaults to `false`.
        :param pulumi.Input[bool] fifo_queue: Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        :param pulumi.Input[str] name: The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SQS is enabled.
        :param pulumi.Input[int] receive_wait_time_seconds: The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        :param pulumi.Input[str] region: `region`). The region in which SQS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SQS queue.
        :param pulumi.Input[str] sqs_endpoint: The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] url: The URL of the queue.
        :param pulumi.Input[int] visibility_timeout_seconds: The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqSqsQueueState.__new__(_MnqSqsQueueState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["content_based_deduplication"] = content_based_deduplication
        __props__.__dict__["fifo_queue"] = fifo_queue
        __props__.__dict__["message_max_age"] = message_max_age
        __props__.__dict__["message_max_size"] = message_max_size
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["receive_wait_time_seconds"] = receive_wait_time_seconds
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["sqs_endpoint"] = sqs_endpoint
        __props__.__dict__["url"] = url
        __props__.__dict__["visibility_timeout_seconds"] = visibility_timeout_seconds
        return MnqSqsQueue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The access key of the SQS queue.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="contentBasedDeduplication")
    def content_based_deduplication(self) -> pulumi.Output[bool]:
        """
        Specifies whether to enable content-based deduplication. Defaults to `false`.
        """
        return pulumi.get(self, "content_based_deduplication")

    @property
    @pulumi.getter(name="fifoQueue")
    def fifo_queue(self) -> pulumi.Output[bool]:
        """
        Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
        """
        return pulumi.get(self, "fifo_queue")

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> pulumi.Output[Optional[int]]:
        """
        The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.
        """
        return pulumi.get(self, "message_max_age")

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.
        """
        return pulumi.get(self, "message_max_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the SQS queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.
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
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the Project in which SQS is enabled.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="receiveWaitTimeSeconds")
    def receive_wait_time_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
        """
        return pulumi.get(self, "receive_wait_time_seconds")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which SQS is enabled.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret key of the SQS queue.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="sqsEndpoint")
    def sqs_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `https://sqs.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sqs_endpoint")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the queue.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="visibilityTimeoutSeconds")
    def visibility_timeout_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
        """
        return pulumi.get(self, "visibility_timeout_seconds")

