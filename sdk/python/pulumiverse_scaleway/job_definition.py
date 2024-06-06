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

__all__ = ['JobDefinitionArgs', 'JobDefinition']

@pulumi.input_type
class JobDefinitionArgs:
    def __init__(__self__, *,
                 cpu_limit: pulumi.Input[int],
                 memory_limit: pulumi.Input[int],
                 command: Optional[pulumi.Input[str]] = None,
                 cron: Optional[pulumi.Input['JobDefinitionCronArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 image_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobDefinition resource.
        :param pulumi.Input[int] cpu_limit: The amount of vCPU computing resources to allocate to each container running the job.
        :param pulumi.Input[int] memory_limit: The memory computing resources in MB to allocate to each container running the job.
        :param pulumi.Input[str] command: The command that will be run in the container if specified.
        :param pulumi.Input['JobDefinitionCronArgs'] cron: The cron configuration
        :param pulumi.Input[str] description: The description of the job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: The environment variables of the container.
        :param pulumi.Input[str] image_uri: The uri of the container image that will be used for the job run.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Job is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Job.
        :param pulumi.Input[str] timeout: The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        pulumi.set(__self__, "cpu_limit", cpu_limit)
        pulumi.set(__self__, "memory_limit", memory_limit)
        if command is not None:
            pulumi.set(__self__, "command", command)
        if cron is not None:
            pulumi.set(__self__, "cron", cron)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if env is not None:
            pulumi.set(__self__, "env", env)
        if image_uri is not None:
            pulumi.set(__self__, "image_uri", image_uri)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> pulumi.Input[int]:
        """
        The amount of vCPU computing resources to allocate to each container running the job.
        """
        return pulumi.get(self, "cpu_limit")

    @cpu_limit.setter
    def cpu_limit(self, value: pulumi.Input[int]):
        pulumi.set(self, "cpu_limit", value)

    @property
    @pulumi.getter(name="memoryLimit")
    def memory_limit(self) -> pulumi.Input[int]:
        """
        The memory computing resources in MB to allocate to each container running the job.
        """
        return pulumi.get(self, "memory_limit")

    @memory_limit.setter
    def memory_limit(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory_limit", value)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[str]]:
        """
        The command that will be run in the container if specified.
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter
    def cron(self) -> Optional[pulumi.Input['JobDefinitionCronArgs']]:
        """
        The cron configuration
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: Optional[pulumi.Input['JobDefinitionCronArgs']]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the job
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def env(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The environment variables of the container.
        """
        return pulumi.get(self, "env")

    @env.setter
    def env(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "env", value)

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The uri of the container image that will be used for the job run.
        """
        return pulumi.get(self, "image_uri")

    @image_uri.setter
    def image_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_uri", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the Job is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the Job.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _JobDefinitionState:
    def __init__(__self__, *,
                 command: Optional[pulumi.Input[str]] = None,
                 cpu_limit: Optional[pulumi.Input[int]] = None,
                 cron: Optional[pulumi.Input['JobDefinitionCronArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 image_uri: Optional[pulumi.Input[str]] = None,
                 memory_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobDefinition resources.
        :param pulumi.Input[str] command: The command that will be run in the container if specified.
        :param pulumi.Input[int] cpu_limit: The amount of vCPU computing resources to allocate to each container running the job.
        :param pulumi.Input['JobDefinitionCronArgs'] cron: The cron configuration
        :param pulumi.Input[str] description: The description of the job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: The environment variables of the container.
        :param pulumi.Input[str] image_uri: The uri of the container image that will be used for the job run.
        :param pulumi.Input[int] memory_limit: The memory computing resources in MB to allocate to each container running the job.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Job is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Job.
        :param pulumi.Input[str] timeout: The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        if command is not None:
            pulumi.set(__self__, "command", command)
        if cpu_limit is not None:
            pulumi.set(__self__, "cpu_limit", cpu_limit)
        if cron is not None:
            pulumi.set(__self__, "cron", cron)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if env is not None:
            pulumi.set(__self__, "env", env)
        if image_uri is not None:
            pulumi.set(__self__, "image_uri", image_uri)
        if memory_limit is not None:
            pulumi.set(__self__, "memory_limit", memory_limit)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[str]]:
        """
        The command that will be run in the container if specified.
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of vCPU computing resources to allocate to each container running the job.
        """
        return pulumi.get(self, "cpu_limit")

    @cpu_limit.setter
    def cpu_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_limit", value)

    @property
    @pulumi.getter
    def cron(self) -> Optional[pulumi.Input['JobDefinitionCronArgs']]:
        """
        The cron configuration
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: Optional[pulumi.Input['JobDefinitionCronArgs']]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the job
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def env(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The environment variables of the container.
        """
        return pulumi.get(self, "env")

    @env.setter
    def env(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "env", value)

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The uri of the container image that will be used for the job run.
        """
        return pulumi.get(self, "image_uri")

    @image_uri.setter
    def image_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_uri", value)

    @property
    @pulumi.getter(name="memoryLimit")
    def memory_limit(self) -> Optional[pulumi.Input[int]]:
        """
        The memory computing resources in MB to allocate to each container running the job.
        """
        return pulumi.get(self, "memory_limit")

    @memory_limit.setter
    def memory_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_limit", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the Job is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the Job.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


class JobDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 cpu_limit: Optional[pulumi.Input[int]] = None,
                 cron: Optional[pulumi.Input[pulumi.InputType['JobDefinitionCronArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 image_uri: Optional[pulumi.Input[str]] = None,
                 memory_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages a Scaleway Serverless Job Definition. For more information, see [the documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/jobs/v1alpha1).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.JobDefinition("main",
            cpu_limit=140,
            memory_limit=256,
            image_uri="docker.io/alpine:latest",
            command="ls",
            timeout="10m",
            env={
                "foo": "bar",
            },
            cron=scaleway.JobDefinitionCronArgs(
                schedule="5 4 1 * *",
                timezone="Europe/Paris",
            ))
        ```

        ## Import

        Serverless Jobs can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/jobDefinition:JobDefinition job fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: The command that will be run in the container if specified.
        :param pulumi.Input[int] cpu_limit: The amount of vCPU computing resources to allocate to each container running the job.
        :param pulumi.Input[pulumi.InputType['JobDefinitionCronArgs']] cron: The cron configuration
        :param pulumi.Input[str] description: The description of the job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: The environment variables of the container.
        :param pulumi.Input[str] image_uri: The uri of the container image that will be used for the job run.
        :param pulumi.Input[int] memory_limit: The memory computing resources in MB to allocate to each container running the job.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Job is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Job.
        :param pulumi.Input[str] timeout: The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages a Scaleway Serverless Job Definition. For more information, see [the documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/jobs/v1alpha1).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        main = scaleway.JobDefinition("main",
            cpu_limit=140,
            memory_limit=256,
            image_uri="docker.io/alpine:latest",
            command="ls",
            timeout="10m",
            env={
                "foo": "bar",
            },
            cron=scaleway.JobDefinitionCronArgs(
                schedule="5 4 1 * *",
                timezone="Europe/Paris",
            ))
        ```

        ## Import

        Serverless Jobs can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/jobDefinition:JobDefinition job fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param JobDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 cpu_limit: Optional[pulumi.Input[int]] = None,
                 cron: Optional[pulumi.Input[pulumi.InputType['JobDefinitionCronArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 image_uri: Optional[pulumi.Input[str]] = None,
                 memory_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobDefinitionArgs.__new__(JobDefinitionArgs)

            __props__.__dict__["command"] = command
            if cpu_limit is None and not opts.urn:
                raise TypeError("Missing required property 'cpu_limit'")
            __props__.__dict__["cpu_limit"] = cpu_limit
            __props__.__dict__["cron"] = cron
            __props__.__dict__["description"] = description
            __props__.__dict__["env"] = env
            __props__.__dict__["image_uri"] = image_uri
            if memory_limit is None and not opts.urn:
                raise TypeError("Missing required property 'memory_limit'")
            __props__.__dict__["memory_limit"] = memory_limit
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["timeout"] = timeout
        super(JobDefinition, __self__).__init__(
            'scaleway:index/jobDefinition:JobDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command: Optional[pulumi.Input[str]] = None,
            cpu_limit: Optional[pulumi.Input[int]] = None,
            cron: Optional[pulumi.Input[pulumi.InputType['JobDefinitionCronArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            env: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            image_uri: Optional[pulumi.Input[str]] = None,
            memory_limit: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[str]] = None) -> 'JobDefinition':
        """
        Get an existing JobDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: The command that will be run in the container if specified.
        :param pulumi.Input[int] cpu_limit: The amount of vCPU computing resources to allocate to each container running the job.
        :param pulumi.Input[pulumi.InputType['JobDefinitionCronArgs']] cron: The cron configuration
        :param pulumi.Input[str] description: The description of the job
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env: The environment variables of the container.
        :param pulumi.Input[str] image_uri: The uri of the container image that will be used for the job run.
        :param pulumi.Input[int] memory_limit: The memory computing resources in MB to allocate to each container running the job.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the Job is associated with.
        :param pulumi.Input[str] region: `region`) The region of the Job.
        :param pulumi.Input[str] timeout: The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobDefinitionState.__new__(_JobDefinitionState)

        __props__.__dict__["command"] = command
        __props__.__dict__["cpu_limit"] = cpu_limit
        __props__.__dict__["cron"] = cron
        __props__.__dict__["description"] = description
        __props__.__dict__["env"] = env
        __props__.__dict__["image_uri"] = image_uri
        __props__.__dict__["memory_limit"] = memory_limit
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["timeout"] = timeout
        return JobDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output[Optional[str]]:
        """
        The command that will be run in the container if specified.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> pulumi.Output[int]:
        """
        The amount of vCPU computing resources to allocate to each container running the job.
        """
        return pulumi.get(self, "cpu_limit")

    @property
    @pulumi.getter
    def cron(self) -> pulumi.Output[Optional['outputs.JobDefinitionCron']]:
        """
        The cron configuration
        """
        return pulumi.get(self, "cron")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the job
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def env(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The environment variables of the container.
        """
        return pulumi.get(self, "env")

    @property
    @pulumi.getter(name="imageUri")
    def image_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The uri of the container image that will be used for the job run.
        """
        return pulumi.get(self, "image_uri")

    @property
    @pulumi.getter(name="memoryLimit")
    def memory_limit(self) -> pulumi.Output[int]:
        """
        The memory computing resources in MB to allocate to each container running the job.
        """
        return pulumi.get(self, "memory_limit")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the Job is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region of the Job.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[str]:
        """
        The job run timeout, in Go Time format (ex: `2h30m25s`)
        """
        return pulumi.get(self, "timeout")

