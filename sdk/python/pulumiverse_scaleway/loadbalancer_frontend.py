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

__all__ = ['LoadbalancerFrontendArgs', 'LoadbalancerFrontend']

@pulumi.input_type
class LoadbalancerFrontendArgs:
    def __init__(__self__, *,
                 backend_id: pulumi.Input[str],
                 inbound_port: pulumi.Input[int],
                 lb_id: pulumi.Input[str],
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoadbalancerFrontend resource.
        :param pulumi.Input[str] backend_id: The ID of the Load Balancer backend this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen to on the front side.
        :param pulumi.Input[str] lb_id: The ID of the Load Balancer this frontend is attached to.
        :param pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]] acls: A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g. `1s`)
        """
        pulumi.set(__self__, "backend_id", backend_id)
        pulumi.set(__self__, "inbound_port", inbound_port)
        pulumi.set(__self__, "lb_id", lb_id)
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if certificate_ids is not None:
            pulumi.set(__self__, "certificate_ids", certificate_ids)
        if enable_http3 is not None:
            pulumi.set(__self__, "enable_http3", enable_http3)
        if external_acls is not None:
            pulumi.set(__self__, "external_acls", external_acls)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timeout_client is not None:
            pulumi.set(__self__, "timeout_client", timeout_client)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Input[str]:
        """
        The ID of the Load Balancer backend this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> pulumi.Input[int]:
        """
        TCP port to listen to on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @inbound_port.setter
    def inbound_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "inbound_port", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Input[str]:
        """
        The ID of the Load Balancer this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]]:
        """
        A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> Optional[pulumi.Input[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @enable_http3.setter
    def enable_http3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_http3", value)

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        """
        return pulumi.get(self, "external_acls")

    @external_acls.setter
    def external_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_acls", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum inactivity time on the client side. (e.g. `1s`)
        """
        return pulumi.get(self, "timeout_client")

    @timeout_client.setter
    def timeout_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_client", value)


@pulumi.input_type
class _LoadbalancerFrontendState:
    def __init__(__self__, *,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoadbalancerFrontend resources.
        :param pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]] acls: A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The ID of the Load Balancer backend this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[str] certificate_id: (Deprecated) First certificate ID used by the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen to on the front side.
        :param pulumi.Input[str] lb_id: The ID of the Load Balancer this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g. `1s`)
        """
        if acls is not None:
            pulumi.set(__self__, "acls", acls)
        if backend_id is not None:
            pulumi.set(__self__, "backend_id", backend_id)
        if certificate_id is not None:
            warnings.warn("""Please use certificate_ids""", DeprecationWarning)
            pulumi.log.warn("""certificate_id is deprecated: Please use certificate_ids""")
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_ids is not None:
            pulumi.set(__self__, "certificate_ids", certificate_ids)
        if enable_http3 is not None:
            pulumi.set(__self__, "enable_http3", enable_http3)
        if external_acls is not None:
            pulumi.set(__self__, "external_acls", external_acls)
        if inbound_port is not None:
            pulumi.set(__self__, "inbound_port", inbound_port)
        if lb_id is not None:
            pulumi.set(__self__, "lb_id", lb_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timeout_client is not None:
            pulumi.set(__self__, "timeout_client", timeout_client)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]]:
        """
        A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LoadbalancerFrontendAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Load Balancer backend this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="certificateId")
    @_utilities.deprecated("""Please use certificate_ids""")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Deprecated) First certificate ID used by the frontend.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> Optional[pulumi.Input[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @enable_http3.setter
    def enable_http3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_http3", value)

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        """
        return pulumi.get(self, "external_acls")

    @external_acls.setter
    def external_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_acls", value)

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> Optional[pulumi.Input[int]]:
        """
        TCP port to listen to on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @inbound_port.setter
    def inbound_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inbound_port", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Load Balancer this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum inactivity time on the client side. (e.g. `1s`)
        """
        return pulumi.get(self, "timeout_client")

    @timeout_client.setter
    def timeout_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_client", value)


class LoadbalancerFrontend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LoadbalancerFrontendAclArgs']]]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Load Balancer frontends.

        For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        frontend01 = scaleway.LoadbalancerFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80)
        ```

        ## With ACLs

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        frontend01 = scaleway.LoadbalancerFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80,
            acls=[
                scaleway.LoadbalancerFrontendAclArgs(
                    name="blacklist wellknwon IPs",
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=[
                            "192.168.0.1",
                            "192.168.0.2",
                            "192.168.10.0/24",
                        ],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="deny",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=["51.51.51.51"],
                        http_filter="regex",
                        http_filter_values=["^foo*bar$"],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="path_begin",
                        http_filter_values=[
                            "foo",
                            "bar",
                        ],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="path_begin",
                        http_filter_values=["hi"],
                        invert=True,
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="http_header_match",
                        http_filter_values="foo",
                        http_filter_option="bar",
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="redirect",
                        redirects=[scaleway.LoadbalancerFrontendAclActionRedirectArgs(
                            type="location",
                            target="https://example.com",
                            code=307,
                        )],
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=["10.0.0.10"],
                        http_filter="path_begin",
                        http_filter_values=[
                            "foo",
                            "bar",
                        ],
                    ),
                ),
            ])
        ```

        ## Import

        Load Balancer frontends can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/loadbalancerFrontend:LoadbalancerFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LoadbalancerFrontendAclArgs']]]] acls: A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The ID of the Load Balancer backend this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen to on the front side.
        :param pulumi.Input[str] lb_id: The ID of the Load Balancer this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g. `1s`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadbalancerFrontendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Load Balancer frontends.

        For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        frontend01 = scaleway.LoadbalancerFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80)
        ```

        ## With ACLs

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        frontend01 = scaleway.LoadbalancerFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80,
            acls=[
                scaleway.LoadbalancerFrontendAclArgs(
                    name="blacklist wellknwon IPs",
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=[
                            "192.168.0.1",
                            "192.168.0.2",
                            "192.168.10.0/24",
                        ],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="deny",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=["51.51.51.51"],
                        http_filter="regex",
                        http_filter_values=["^foo*bar$"],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="path_begin",
                        http_filter_values=[
                            "foo",
                            "bar",
                        ],
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="path_begin",
                        http_filter_values=["hi"],
                        invert=True,
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="allow",
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        http_filter="http_header_match",
                        http_filter_values="foo",
                        http_filter_option="bar",
                    ),
                ),
                scaleway.LoadbalancerFrontendAclArgs(
                    action=scaleway.LoadbalancerFrontendAclActionArgs(
                        type="redirect",
                        redirects=[scaleway.LoadbalancerFrontendAclActionRedirectArgs(
                            type="location",
                            target="https://example.com",
                            code=307,
                        )],
                    ),
                    match=scaleway.LoadbalancerFrontendAclMatchArgs(
                        ip_subnets=["10.0.0.10"],
                        http_filter="path_begin",
                        http_filter_values=[
                            "foo",
                            "bar",
                        ],
                    ),
                ),
            ])
        ```

        ## Import

        Load Balancer frontends can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/loadbalancerFrontend:LoadbalancerFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param LoadbalancerFrontendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadbalancerFrontendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LoadbalancerFrontendAclArgs']]]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadbalancerFrontendArgs.__new__(LoadbalancerFrontendArgs)

            __props__.__dict__["acls"] = acls
            if backend_id is None and not opts.urn:
                raise TypeError("Missing required property 'backend_id'")
            __props__.__dict__["backend_id"] = backend_id
            __props__.__dict__["certificate_ids"] = certificate_ids
            __props__.__dict__["enable_http3"] = enable_http3
            __props__.__dict__["external_acls"] = external_acls
            if inbound_port is None and not opts.urn:
                raise TypeError("Missing required property 'inbound_port'")
            __props__.__dict__["inbound_port"] = inbound_port
            if lb_id is None and not opts.urn:
                raise TypeError("Missing required property 'lb_id'")
            __props__.__dict__["lb_id"] = lb_id
            __props__.__dict__["name"] = name
            __props__.__dict__["timeout_client"] = timeout_client
            __props__.__dict__["certificate_id"] = None
        super(LoadbalancerFrontend, __self__).__init__(
            'scaleway:index/loadbalancerFrontend:LoadbalancerFrontend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LoadbalancerFrontendAclArgs']]]]] = None,
            backend_id: Optional[pulumi.Input[str]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_http3: Optional[pulumi.Input[bool]] = None,
            external_acls: Optional[pulumi.Input[bool]] = None,
            inbound_port: Optional[pulumi.Input[int]] = None,
            lb_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            timeout_client: Optional[pulumi.Input[str]] = None) -> 'LoadbalancerFrontend':
        """
        Get an existing LoadbalancerFrontend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LoadbalancerFrontendAclArgs']]]] acls: A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The ID of the Load Balancer backend this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[str] certificate_id: (Deprecated) First certificate ID used by the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen to on the front side.
        :param pulumi.Input[str] lb_id: The ID of the Load Balancer this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g. `1s`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadbalancerFrontendState.__new__(_LoadbalancerFrontendState)

        __props__.__dict__["acls"] = acls
        __props__.__dict__["backend_id"] = backend_id
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["certificate_ids"] = certificate_ids
        __props__.__dict__["enable_http3"] = enable_http3
        __props__.__dict__["external_acls"] = external_acls
        __props__.__dict__["inbound_port"] = inbound_port
        __props__.__dict__["lb_id"] = lb_id
        __props__.__dict__["name"] = name
        __props__.__dict__["timeout_client"] = timeout_client
        return LoadbalancerFrontend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Optional[Sequence['outputs.LoadbalancerFrontendAcl']]]:
        """
        A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Output[str]:
        """
        The ID of the Load Balancer backend this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="certificateId")
    @_utilities.deprecated("""Please use certificate_ids""")
    def certificate_id(self) -> pulumi.Output[str]:
        """
        (Deprecated) First certificate ID used by the frontend.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> pulumi.Output[Optional[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the Load Balancer frontend.
        """
        return pulumi.get(self, "external_acls")

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> pulumi.Output[int]:
        """
        TCP port to listen to on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Output[str]:
        """
        The ID of the Load Balancer this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> pulumi.Output[Optional[str]]:
        """
        Maximum inactivity time on the client side. (e.g. `1s`)
        """
        return pulumi.get(self, "timeout_client")

