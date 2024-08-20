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

__all__ = ['FunctionTriggerArgs', 'FunctionTrigger']

@pulumi.input_type
class FunctionTriggerArgs:
    def __init__(__self__, *,
                 function_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['FunctionTriggerNatsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input['FunctionTriggerSqsArgs']] = None):
        """
        The set of arguments for constructing a FunctionTrigger resource.
        :param pulumi.Input[str] function_id: The ID of the function to create a trigger for
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input['FunctionTriggerNatsArgs'] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input['FunctionTriggerSqsArgs'] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        pulumi.set(__self__, "function_id", function_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Input[str]:
        """
        The ID of the function to create a trigger for
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['FunctionTriggerNatsArgs']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['FunctionTriggerNatsArgs']]):
        pulumi.set(self, "nats", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['FunctionTriggerSqsArgs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['FunctionTriggerSqsArgs']]):
        pulumi.set(self, "sqs", value)


@pulumi.input_type
class _FunctionTriggerState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['FunctionTriggerNatsArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input['FunctionTriggerSqsArgs']] = None):
        """
        Input properties used for looking up and filtering FunctionTrigger resources.
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] function_id: The ID of the function to create a trigger for
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input['FunctionTriggerNatsArgs'] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input['FunctionTriggerSqsArgs'] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the function to create a trigger for
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['FunctionTriggerNatsArgs']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['FunctionTriggerNatsArgs']]):
        pulumi.set(self, "nats", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['FunctionTriggerSqsArgs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['FunctionTriggerSqsArgs']]):
        pulumi.set(self, "sqs", value)


class FunctionTrigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[Union['FunctionTriggerNatsArgs', 'FunctionTriggerNatsArgsDict']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input[Union['FunctionTriggerSqsArgs', 'FunctionTriggerSqsArgsDict']]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Function Triggers.
        For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers).

        ## Example Usage

        ### SQS

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.FunctionTrigger("main",
            function_id=scaleway_function["main"]["id"],
            sqs={
                "project_id": scaleway_mnq_sqs["main"]["project_id"],
                "queue": "MyQueue",
                "region": scaleway_mnq_sqs["main"]["region"],
            })
        ```

        ### Nats

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.FunctionTrigger("main",
            function_id=scaleway_function["main"]["id"],
            nats={
                "account_id": scaleway_mnq_nats_account["main"]["id"],
                "subject": "MySubject",
                "region": scaleway_mnq_nats_account["main"]["region"],
            })
        ```

        ## Import

        Function Triggers can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/functionTrigger:FunctionTrigger main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] function_id: The ID of the function to create a trigger for
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input[Union['FunctionTriggerNatsArgs', 'FunctionTriggerNatsArgsDict']] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[Union['FunctionTriggerSqsArgs', 'FunctionTriggerSqsArgsDict']] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionTriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Function Triggers.
        For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers).

        ## Example Usage

        ### SQS

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.FunctionTrigger("main",
            function_id=scaleway_function["main"]["id"],
            sqs={
                "project_id": scaleway_mnq_sqs["main"]["project_id"],
                "queue": "MyQueue",
                "region": scaleway_mnq_sqs["main"]["region"],
            })
        ```

        ### Nats

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.FunctionTrigger("main",
            function_id=scaleway_function["main"]["id"],
            nats={
                "account_id": scaleway_mnq_nats_account["main"]["id"],
                "subject": "MySubject",
                "region": scaleway_mnq_nats_account["main"]["region"],
            })
        ```

        ## Import

        Function Triggers can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/functionTrigger:FunctionTrigger main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param FunctionTriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionTriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[Union['FunctionTriggerNatsArgs', 'FunctionTriggerNatsArgsDict']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input[Union['FunctionTriggerSqsArgs', 'FunctionTriggerSqsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionTriggerArgs.__new__(FunctionTriggerArgs)

            __props__.__dict__["description"] = description
            if function_id is None and not opts.urn:
                raise TypeError("Missing required property 'function_id'")
            __props__.__dict__["function_id"] = function_id
            __props__.__dict__["name"] = name
            __props__.__dict__["nats"] = nats
            __props__.__dict__["region"] = region
            __props__.__dict__["sqs"] = sqs
        super(FunctionTrigger, __self__).__init__(
            'scaleway:index/functionTrigger:FunctionTrigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nats: Optional[pulumi.Input[Union['FunctionTriggerNatsArgs', 'FunctionTriggerNatsArgsDict']]] = None,
            region: Optional[pulumi.Input[str]] = None,
            sqs: Optional[pulumi.Input[Union['FunctionTriggerSqsArgs', 'FunctionTriggerSqsArgsDict']]] = None) -> 'FunctionTrigger':
        """
        Get an existing FunctionTrigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the trigger.
        :param pulumi.Input[str] function_id: The ID of the function to create a trigger for
        :param pulumi.Input[str] name: The unique name of the trigger. Default to a generated name.
        :param pulumi.Input[Union['FunctionTriggerNatsArgs', 'FunctionTriggerNatsArgsDict']] nats: The configuration for the Scaleway's Nats used by the trigger
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
        :param pulumi.Input[Union['FunctionTriggerSqsArgs', 'FunctionTriggerSqsArgsDict']] sqs: The configuration of the Scaleway's SQS used by the trigger
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionTriggerState.__new__(_FunctionTriggerState)

        __props__.__dict__["description"] = description
        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["name"] = name
        __props__.__dict__["nats"] = nats
        __props__.__dict__["region"] = region
        __props__.__dict__["sqs"] = sqs
        return FunctionTrigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the trigger.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        """
        The ID of the function to create a trigger for
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the trigger. Default to a generated name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nats(self) -> pulumi.Output[Optional['outputs.FunctionTriggerNats']]:
        """
        The configuration for the Scaleway's Nats used by the trigger
        """
        return pulumi.get(self, "nats")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def sqs(self) -> pulumi.Output[Optional['outputs.FunctionTriggerSqs']]:
        """
        The configuration of the Scaleway's SQS used by the trigger
        """
        return pulumi.get(self, "sqs")

