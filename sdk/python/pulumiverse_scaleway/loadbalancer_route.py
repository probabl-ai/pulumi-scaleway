# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LoadbalancerRouteArgs', 'LoadbalancerRoute']

@pulumi.input_type
class LoadbalancerRouteArgs:
    def __init__(__self__, *,
                 backend_id: pulumi.Input[str],
                 frontend_id: pulumi.Input[str],
                 match_host_header: Optional[pulumi.Input[str]] = None,
                 match_sni: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoadbalancerRoute resource.
        :param pulumi.Input[str] backend_id: The ID of the backend the route is associated with.
        :param pulumi.Input[str] frontend_id: The ID of the frontend the route is associated with.
        :param pulumi.Input[str] match_host_header: The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on HTTP Load Balancers.
        :param pulumi.Input[str] match_sni: The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on TCP Load Balancers.
        """
        pulumi.set(__self__, "backend_id", backend_id)
        pulumi.set(__self__, "frontend_id", frontend_id)
        if match_host_header is not None:
            pulumi.set(__self__, "match_host_header", match_host_header)
        if match_sni is not None:
            pulumi.set(__self__, "match_sni", match_sni)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Input[str]:
        """
        The ID of the backend the route is associated with.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> pulumi.Input[str]:
        """
        The ID of the frontend the route is associated with.
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter(name="matchHostHeader")
    def match_host_header(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on HTTP Load Balancers.
        """
        return pulumi.get(self, "match_host_header")

    @match_host_header.setter
    def match_host_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_host_header", value)

    @property
    @pulumi.getter(name="matchSni")
    def match_sni(self) -> Optional[pulumi.Input[str]]:
        """
        The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on TCP Load Balancers.
        """
        return pulumi.get(self, "match_sni")

    @match_sni.setter
    def match_sni(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_sni", value)


@pulumi.input_type
class _LoadbalancerRouteState:
    def __init__(__self__, *,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 match_host_header: Optional[pulumi.Input[str]] = None,
                 match_sni: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoadbalancerRoute resources.
        :param pulumi.Input[str] backend_id: The ID of the backend the route is associated with.
        :param pulumi.Input[str] created_at: The date on which the route was created.
        :param pulumi.Input[str] frontend_id: The ID of the frontend the route is associated with.
        :param pulumi.Input[str] match_host_header: The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on HTTP Load Balancers.
        :param pulumi.Input[str] match_sni: The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on TCP Load Balancers.
        :param pulumi.Input[str] updated_at: The date on which the route was last updated.
        """
        if backend_id is not None:
            pulumi.set(__self__, "backend_id", backend_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if frontend_id is not None:
            pulumi.set(__self__, "frontend_id", frontend_id)
        if match_host_header is not None:
            pulumi.set(__self__, "match_host_header", match_host_header)
        if match_sni is not None:
            pulumi.set(__self__, "match_sni", match_sni)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the backend the route is associated with.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date on which the route was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the frontend the route is associated with.
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter(name="matchHostHeader")
    def match_host_header(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on HTTP Load Balancers.
        """
        return pulumi.get(self, "match_host_header")

    @match_host_header.setter
    def match_host_header(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_host_header", value)

    @property
    @pulumi.getter(name="matchSni")
    def match_sni(self) -> Optional[pulumi.Input[str]]:
        """
        The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on TCP Load Balancers.
        """
        return pulumi.get(self, "match_sni")

    @match_sni.setter
    def match_sni(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_sni", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date on which the route was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class LoadbalancerRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 match_host_header: Optional[pulumi.Input[str]] = None,
                 match_sni: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Load Balancer routes.

        For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).

        ## Example Usage

        ### With SNI for direction to TCP backends

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        ip01 = scaleway.LoadbalancerIp("ip01")
        lb01 = scaleway.Loadbalancer("lb01",
            ip_id=ip01.id,
            type="lb-s")
        bkd01 = scaleway.LoadbalancerBackend("bkd01",
            lb_id=lb01.id,
            forward_protocol="tcp",
            forward_port=80,
            proxy_protocol="none")
        frt01 = scaleway.LoadbalancerFrontend("frt01",
            lb_id=lb01.id,
            backend_id=bkd01.id,
            inbound_port=80)
        rt01 = scaleway.LoadbalancerRoute("rt01",
            frontend_id=frt01.id,
            backend_id=bkd01.id,
            match_sni="sni.scaleway.com")
        ```

        ### With host-header for direction to HTTP backends

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        ip01 = scaleway.LoadbalancerIp("ip01")
        lb01 = scaleway.Loadbalancer("lb01",
            ip_id=ip01.id,
            type="lb-s")
        bkd01 = scaleway.LoadbalancerBackend("bkd01",
            lb_id=lb01.id,
            forward_protocol="http",
            forward_port=80,
            proxy_protocol="none")
        frt01 = scaleway.LoadbalancerFrontend("frt01",
            lb_id=lb01.id,
            backend_id=bkd01.id,
            inbound_port=80)
        rt01 = scaleway.LoadbalancerRoute("rt01",
            frontend_id=frt01.id,
            backend_id=bkd01.id,
            match_host_header="host.scaleway.com")
        ```

        ## Import

        Load Balancer frontends can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/loadbalancerRoute:LoadbalancerRoute main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_id: The ID of the backend the route is associated with.
        :param pulumi.Input[str] frontend_id: The ID of the frontend the route is associated with.
        :param pulumi.Input[str] match_host_header: The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on HTTP Load Balancers.
        :param pulumi.Input[str] match_sni: The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on TCP Load Balancers.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadbalancerRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Load Balancer routes.

        For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).

        ## Example Usage

        ### With SNI for direction to TCP backends

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        ip01 = scaleway.LoadbalancerIp("ip01")
        lb01 = scaleway.Loadbalancer("lb01",
            ip_id=ip01.id,
            type="lb-s")
        bkd01 = scaleway.LoadbalancerBackend("bkd01",
            lb_id=lb01.id,
            forward_protocol="tcp",
            forward_port=80,
            proxy_protocol="none")
        frt01 = scaleway.LoadbalancerFrontend("frt01",
            lb_id=lb01.id,
            backend_id=bkd01.id,
            inbound_port=80)
        rt01 = scaleway.LoadbalancerRoute("rt01",
            frontend_id=frt01.id,
            backend_id=bkd01.id,
            match_sni="sni.scaleway.com")
        ```

        ### With host-header for direction to HTTP backends

        ```python
        import pulumi
        import pulumiverse_scaleway as scaleway

        ip01 = scaleway.LoadbalancerIp("ip01")
        lb01 = scaleway.Loadbalancer("lb01",
            ip_id=ip01.id,
            type="lb-s")
        bkd01 = scaleway.LoadbalancerBackend("bkd01",
            lb_id=lb01.id,
            forward_protocol="http",
            forward_port=80,
            proxy_protocol="none")
        frt01 = scaleway.LoadbalancerFrontend("frt01",
            lb_id=lb01.id,
            backend_id=bkd01.id,
            inbound_port=80)
        rt01 = scaleway.LoadbalancerRoute("rt01",
            frontend_id=frt01.id,
            backend_id=bkd01.id,
            match_host_header="host.scaleway.com")
        ```

        ## Import

        Load Balancer frontends can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/loadbalancerRoute:LoadbalancerRoute main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param LoadbalancerRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadbalancerRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 match_host_header: Optional[pulumi.Input[str]] = None,
                 match_sni: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadbalancerRouteArgs.__new__(LoadbalancerRouteArgs)

            if backend_id is None and not opts.urn:
                raise TypeError("Missing required property 'backend_id'")
            __props__.__dict__["backend_id"] = backend_id
            if frontend_id is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_id'")
            __props__.__dict__["frontend_id"] = frontend_id
            __props__.__dict__["match_host_header"] = match_host_header
            __props__.__dict__["match_sni"] = match_sni
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(LoadbalancerRoute, __self__).__init__(
            'scaleway:index/loadbalancerRoute:LoadbalancerRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            frontend_id: Optional[pulumi.Input[str]] = None,
            match_host_header: Optional[pulumi.Input[str]] = None,
            match_sni: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'LoadbalancerRoute':
        """
        Get an existing LoadbalancerRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend_id: The ID of the backend the route is associated with.
        :param pulumi.Input[str] created_at: The date on which the route was created.
        :param pulumi.Input[str] frontend_id: The ID of the frontend the route is associated with.
        :param pulumi.Input[str] match_host_header: The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on HTTP Load Balancers.
        :param pulumi.Input[str] match_sni: The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
               Only one of `match_sni` and `match_host_header` should be specified.
               
               > **Important:** This field should be set for routes on TCP Load Balancers.
        :param pulumi.Input[str] updated_at: The date on which the route was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoadbalancerRouteState.__new__(_LoadbalancerRouteState)

        __props__.__dict__["backend_id"] = backend_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["frontend_id"] = frontend_id
        __props__.__dict__["match_host_header"] = match_host_header
        __props__.__dict__["match_sni"] = match_sni
        __props__.__dict__["updated_at"] = updated_at
        return LoadbalancerRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Output[str]:
        """
        The ID of the backend the route is associated with.
        """
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date on which the route was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> pulumi.Output[str]:
        """
        The ID of the frontend the route is associated with.
        """
        return pulumi.get(self, "frontend_id")

    @property
    @pulumi.getter(name="matchHostHeader")
    def match_host_header(self) -> pulumi.Output[Optional[str]]:
        """
        The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on HTTP Load Balancers.
        """
        return pulumi.get(self, "match_host_header")

    @property
    @pulumi.getter(name="matchSni")
    def match_sni(self) -> pulumi.Output[Optional[str]]:
        """
        The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
        Only one of `match_sni` and `match_host_header` should be specified.

        > **Important:** This field should be set for routes on TCP Load Balancers.
        """
        return pulumi.get(self, "match_sni")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date on which the route was last updated.
        """
        return pulumi.get(self, "updated_at")

