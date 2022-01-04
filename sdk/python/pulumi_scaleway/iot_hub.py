# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IOTHubArgs', 'IOTHub']

@pulumi.input_type
class IOTHubArgs:
    def __init__(__self__, *,
                 product_plan: pulumi.Input[str],
                 device_auto_provisioning: Optional[pulumi.Input[bool]] = None,
                 disable_events: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_topic_prefix: Optional[pulumi.Input[str]] = None,
                 hub_ca: Optional[pulumi.Input[str]] = None,
                 hub_ca_challenge: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IOTHub resource.
        :param pulumi.Input[str] product_plan: Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        :param pulumi.Input[bool] device_auto_provisioning: Wether to enable the device auto provisioning or not
        :param pulumi.Input[bool] disable_events: Whether to enable the hub events or not
        :param pulumi.Input[bool] enabled: Wether the IoT Hub instance should be enabled or not.
        :param pulumi.Input[str] events_topic_prefix: Topic prefix for the hub events
        :param pulumi.Input[str] hub_ca: Custom user provided certificate authority
        :param pulumi.Input[str] hub_ca_challenge: Challenge certificate for the user provided hub CA
        :param pulumi.Input[str] name: The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IoT Hub Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        pulumi.set(__self__, "product_plan", product_plan)
        if device_auto_provisioning is not None:
            pulumi.set(__self__, "device_auto_provisioning", device_auto_provisioning)
        if disable_events is not None:
            pulumi.set(__self__, "disable_events", disable_events)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if events_topic_prefix is not None:
            pulumi.set(__self__, "events_topic_prefix", events_topic_prefix)
        if hub_ca is not None:
            pulumi.set(__self__, "hub_ca", hub_ca)
        if hub_ca_challenge is not None:
            pulumi.set(__self__, "hub_ca_challenge", hub_ca_challenge)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="productPlan")
    def product_plan(self) -> pulumi.Input[str]:
        """
        Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        """
        return pulumi.get(self, "product_plan")

    @product_plan.setter
    def product_plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "product_plan", value)

    @property
    @pulumi.getter(name="deviceAutoProvisioning")
    def device_auto_provisioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Wether to enable the device auto provisioning or not
        """
        return pulumi.get(self, "device_auto_provisioning")

    @device_auto_provisioning.setter
    def device_auto_provisioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "device_auto_provisioning", value)

    @property
    @pulumi.getter(name="disableEvents")
    def disable_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the hub events or not
        """
        return pulumi.get(self, "disable_events")

    @disable_events.setter
    def disable_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_events", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Wether the IoT Hub instance should be enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="eventsTopicPrefix")
    def events_topic_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Topic prefix for the hub events
        """
        return pulumi.get(self, "events_topic_prefix")

    @events_topic_prefix.setter
    def events_topic_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "events_topic_prefix", value)

    @property
    @pulumi.getter(name="hubCa")
    def hub_ca(self) -> Optional[pulumi.Input[str]]:
        """
        Custom user provided certificate authority
        """
        return pulumi.get(self, "hub_ca")

    @hub_ca.setter
    def hub_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_ca", value)

    @property
    @pulumi.getter(name="hubCaChallenge")
    def hub_ca_challenge(self) -> Optional[pulumi.Input[str]]:
        """
        Challenge certificate for the user provided hub CA
        """
        return pulumi.get(self, "hub_ca_challenge")

    @hub_ca_challenge.setter
    def hub_ca_challenge(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_ca_challenge", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the IoT Hub Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _IOTHubState:
    def __init__(__self__, *,
                 connected_device_count: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 device_auto_provisioning: Optional[pulumi.Input[bool]] = None,
                 device_count: Optional[pulumi.Input[int]] = None,
                 disable_events: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 events_topic_prefix: Optional[pulumi.Input[str]] = None,
                 hub_ca: Optional[pulumi.Input[str]] = None,
                 hub_ca_challenge: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 product_plan: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IOTHub resources.
        :param pulumi.Input[int] connected_device_count: The current number of connected devices in the Hub.
        :param pulumi.Input[str] created_at: The date and time the Hub was created.
        :param pulumi.Input[bool] device_auto_provisioning: Wether to enable the device auto provisioning or not
        :param pulumi.Input[int] device_count: The number of registered devices in the Hub.
        :param pulumi.Input[bool] disable_events: Whether to enable the hub events or not
        :param pulumi.Input[bool] enabled: Wether the IoT Hub instance should be enabled or not.
        :param pulumi.Input[str] endpoint: The MQTT network endpoint to connect MQTT devices to.
        :param pulumi.Input[str] events_topic_prefix: Topic prefix for the hub events
        :param pulumi.Input[str] hub_ca: Custom user provided certificate authority
        :param pulumi.Input[str] hub_ca_challenge: Challenge certificate for the user provided hub CA
        :param pulumi.Input[str] name: The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] product_plan: Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IoT Hub Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        :param pulumi.Input[str] status: The current status of the Hub.
        :param pulumi.Input[str] updated_at: The date and time the Hub resource was updated.
        """
        if connected_device_count is not None:
            pulumi.set(__self__, "connected_device_count", connected_device_count)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if device_auto_provisioning is not None:
            pulumi.set(__self__, "device_auto_provisioning", device_auto_provisioning)
        if device_count is not None:
            pulumi.set(__self__, "device_count", device_count)
        if disable_events is not None:
            pulumi.set(__self__, "disable_events", disable_events)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if events_topic_prefix is not None:
            pulumi.set(__self__, "events_topic_prefix", events_topic_prefix)
        if hub_ca is not None:
            pulumi.set(__self__, "hub_ca", hub_ca)
        if hub_ca_challenge is not None:
            pulumi.set(__self__, "hub_ca_challenge", hub_ca_challenge)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if product_plan is not None:
            pulumi.set(__self__, "product_plan", product_plan)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="connectedDeviceCount")
    def connected_device_count(self) -> Optional[pulumi.Input[int]]:
        """
        The current number of connected devices in the Hub.
        """
        return pulumi.get(self, "connected_device_count")

    @connected_device_count.setter
    def connected_device_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connected_device_count", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Hub was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="deviceAutoProvisioning")
    def device_auto_provisioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Wether to enable the device auto provisioning or not
        """
        return pulumi.get(self, "device_auto_provisioning")

    @device_auto_provisioning.setter
    def device_auto_provisioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "device_auto_provisioning", value)

    @property
    @pulumi.getter(name="deviceCount")
    def device_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of registered devices in the Hub.
        """
        return pulumi.get(self, "device_count")

    @device_count.setter
    def device_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_count", value)

    @property
    @pulumi.getter(name="disableEvents")
    def disable_events(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the hub events or not
        """
        return pulumi.get(self, "disable_events")

    @disable_events.setter
    def disable_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_events", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Wether the IoT Hub instance should be enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The MQTT network endpoint to connect MQTT devices to.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="eventsTopicPrefix")
    def events_topic_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Topic prefix for the hub events
        """
        return pulumi.get(self, "events_topic_prefix")

    @events_topic_prefix.setter
    def events_topic_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "events_topic_prefix", value)

    @property
    @pulumi.getter(name="hubCa")
    def hub_ca(self) -> Optional[pulumi.Input[str]]:
        """
        Custom user provided certificate authority
        """
        return pulumi.get(self, "hub_ca")

    @hub_ca.setter
    def hub_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_ca", value)

    @property
    @pulumi.getter(name="hubCaChallenge")
    def hub_ca_challenge(self) -> Optional[pulumi.Input[str]]:
        """
        Challenge certificate for the user provided hub CA
        """
        return pulumi.get(self, "hub_ca_challenge")

    @hub_ca_challenge.setter
    def hub_ca_challenge(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_ca_challenge", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization_id you want to attach the resource to
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="productPlan")
    def product_plan(self) -> Optional[pulumi.Input[str]]:
        """
        Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        """
        return pulumi.get(self, "product_plan")

    @product_plan.setter
    def product_plan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_plan", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the IoT Hub Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The current status of the Hub.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Hub resource was updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class IOTHub(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_auto_provisioning: Optional[pulumi.Input[bool]] = None,
                 disable_events: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_topic_prefix: Optional[pulumi.Input[str]] = None,
                 hub_ca: Optional[pulumi.Input[str]] = None,
                 hub_ca_challenge: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product_plan: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        IoT Hubs can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/iOTHub:IOTHub hub01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] device_auto_provisioning: Wether to enable the device auto provisioning or not
        :param pulumi.Input[bool] disable_events: Whether to enable the hub events or not
        :param pulumi.Input[bool] enabled: Wether the IoT Hub instance should be enabled or not.
        :param pulumi.Input[str] events_topic_prefix: Topic prefix for the hub events
        :param pulumi.Input[str] hub_ca: Custom user provided certificate authority
        :param pulumi.Input[str] hub_ca_challenge: Challenge certificate for the user provided hub CA
        :param pulumi.Input[str] name: The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        :param pulumi.Input[str] product_plan: Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IoT Hub Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IOTHubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        IoT Hubs can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/iOTHub:IOTHub hub01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IOTHubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IOTHubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_auto_provisioning: Optional[pulumi.Input[bool]] = None,
                 disable_events: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 events_topic_prefix: Optional[pulumi.Input[str]] = None,
                 hub_ca: Optional[pulumi.Input[str]] = None,
                 hub_ca_challenge: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 product_plan: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IOTHubArgs.__new__(IOTHubArgs)

            __props__.__dict__["device_auto_provisioning"] = device_auto_provisioning
            __props__.__dict__["disable_events"] = disable_events
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["events_topic_prefix"] = events_topic_prefix
            __props__.__dict__["hub_ca"] = hub_ca
            __props__.__dict__["hub_ca_challenge"] = hub_ca_challenge
            __props__.__dict__["name"] = name
            if product_plan is None and not opts.urn:
                raise TypeError("Missing required property 'product_plan'")
            __props__.__dict__["product_plan"] = product_plan
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["connected_device_count"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["device_count"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(IOTHub, __self__).__init__(
            'scaleway:index/iOTHub:IOTHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connected_device_count: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            device_auto_provisioning: Optional[pulumi.Input[bool]] = None,
            device_count: Optional[pulumi.Input[int]] = None,
            disable_events: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            events_topic_prefix: Optional[pulumi.Input[str]] = None,
            hub_ca: Optional[pulumi.Input[str]] = None,
            hub_ca_challenge: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            product_plan: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'IOTHub':
        """
        Get an existing IOTHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connected_device_count: The current number of connected devices in the Hub.
        :param pulumi.Input[str] created_at: The date and time the Hub was created.
        :param pulumi.Input[bool] device_auto_provisioning: Wether to enable the device auto provisioning or not
        :param pulumi.Input[int] device_count: The number of registered devices in the Hub.
        :param pulumi.Input[bool] disable_events: Whether to enable the hub events or not
        :param pulumi.Input[bool] enabled: Wether the IoT Hub instance should be enabled or not.
        :param pulumi.Input[str] endpoint: The MQTT network endpoint to connect MQTT devices to.
        :param pulumi.Input[str] events_topic_prefix: Topic prefix for the hub events
        :param pulumi.Input[str] hub_ca: Custom user provided certificate authority
        :param pulumi.Input[str] hub_ca_challenge: Challenge certificate for the user provided hub CA
        :param pulumi.Input[str] name: The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        :param pulumi.Input[str] organization_id: The organization_id you want to attach the resource to
        :param pulumi.Input[str] product_plan: Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IoT Hub Instance is associated with.
        :param pulumi.Input[str] region: `region`) The region in which the Database Instance should be created.
        :param pulumi.Input[str] status: The current status of the Hub.
        :param pulumi.Input[str] updated_at: The date and time the Hub resource was updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IOTHubState.__new__(_IOTHubState)

        __props__.__dict__["connected_device_count"] = connected_device_count
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["device_auto_provisioning"] = device_auto_provisioning
        __props__.__dict__["device_count"] = device_count
        __props__.__dict__["disable_events"] = disable_events
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["events_topic_prefix"] = events_topic_prefix
        __props__.__dict__["hub_ca"] = hub_ca
        __props__.__dict__["hub_ca_challenge"] = hub_ca_challenge
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["product_plan"] = product_plan
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["updated_at"] = updated_at
        return IOTHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectedDeviceCount")
    def connected_device_count(self) -> pulumi.Output[int]:
        """
        The current number of connected devices in the Hub.
        """
        return pulumi.get(self, "connected_device_count")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the Hub was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deviceAutoProvisioning")
    def device_auto_provisioning(self) -> pulumi.Output[Optional[bool]]:
        """
        Wether to enable the device auto provisioning or not
        """
        return pulumi.get(self, "device_auto_provisioning")

    @property
    @pulumi.getter(name="deviceCount")
    def device_count(self) -> pulumi.Output[int]:
        """
        The number of registered devices in the Hub.
        """
        return pulumi.get(self, "device_count")

    @property
    @pulumi.getter(name="disableEvents")
    def disable_events(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the hub events or not
        """
        return pulumi.get(self, "disable_events")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Wether the IoT Hub instance should be enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The MQTT network endpoint to connect MQTT devices to.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="eventsTopicPrefix")
    def events_topic_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Topic prefix for the hub events
        """
        return pulumi.get(self, "events_topic_prefix")

    @property
    @pulumi.getter(name="hubCa")
    def hub_ca(self) -> pulumi.Output[Optional[str]]:
        """
        Custom user provided certificate authority
        """
        return pulumi.get(self, "hub_ca")

    @property
    @pulumi.getter(name="hubCaChallenge")
    def hub_ca_challenge(self) -> pulumi.Output[Optional[str]]:
        """
        Challenge certificate for the user provided hub CA
        """
        return pulumi.get(self, "hub_ca_challenge")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IoT Hub instance you want to create (e.g. `my-hub`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization_id you want to attach the resource to
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="productPlan")
    def product_plan(self) -> pulumi.Output[str]:
        """
        Product plan to create the hub, see documentation for available product plans (e.g. `plan_shared`)
        """
        return pulumi.get(self, "product_plan")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the IoT Hub Instance is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the Database Instance should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current status of the Hub.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time the Hub resource was updated.
        """
        return pulumi.get(self, "updated_at")

