// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway VPC Public Gateway Network.
    /// It allows attaching Private Networks to the VPC Public Gateway and your DHCP config
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway/#step-3-attach-private-networks-to-the-vpc-public-gateway).
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a gateway network with IPAM config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Vpc("vpc01");
    /// 
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01", new()
    ///     {
    ///         Ipv4Subnet = new Scaleway.Inputs.VpcPrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.64.0/22",
    ///         },
    ///         VpcId = vpc01.Id,
    ///     });
    /// 
    ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
    ///     {
    ///         Type = "VPC-GW-S",
    ///     });
    /// 
    ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         EnableMasquerade = true,
    ///         IpamConfigs = new[]
    ///         {
    ///             new Scaleway.Inputs.VpcGatewayNetworkIpamConfigArgs
    ///             {
    ///                 PushDefaultRoute = true,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create a gateway network with a booked IPAM IP
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var vpc01 = new Scaleway.Vpc("vpc01");
    /// 
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01", new()
    ///     {
    ///         Ipv4Subnet = new Scaleway.Inputs.VpcPrivateNetworkIpv4SubnetArgs
    ///         {
    ///             Subnet = "172.16.64.0/22",
    ///         },
    ///         VpcId = vpc01.Id,
    ///     });
    /// 
    ///     var ip01 = new Scaleway.IpamIp("ip01", new()
    ///     {
    ///         Address = "172.16.64.7",
    ///         Sources = new[]
    ///         {
    ///             new Scaleway.Inputs.IpamIpSourceArgs
    ///             {
    ///                 PrivateNetworkId = pn01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
    ///     {
    ///         Type = "VPC-GW-S",
    ///     });
    /// 
    ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         EnableMasquerade = true,
    ///         IpamConfigs = new[]
    ///         {
    ///             new Scaleway.Inputs.VpcGatewayNetworkIpamConfigArgs
    ///             {
    ///                 PushDefaultRoute = true,
    ///                 IpamIpId = ip01.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create a gateway network with DHCP
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01");
    /// 
    ///     var gw01 = new Scaleway.VpcPublicGatewayIp("gw01");
    /// 
    ///     var dhcp01 = new Scaleway.VpcPublicGatewayDhcp("dhcp01", new()
    ///     {
    ///         Subnet = "192.168.1.0/24",
    ///         PushDefaultRoute = true,
    ///     });
    /// 
    ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
    ///     {
    ///         Type = "VPC-GW-S",
    ///         IpId = gw01.Id,
    ///     });
    /// 
    ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         DhcpId = dhcp01.Id,
    ///         CleanupDhcp = true,
    ///         EnableMasquerade = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create a gateway network with a static IP address
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumiverse.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn01 = new Scaleway.VpcPrivateNetwork("pn01");
    /// 
    ///     var pg01 = new Scaleway.VpcPublicGateway("pg01", new()
    ///     {
    ///         Type = "VPC-GW-S",
    ///     });
    /// 
    ///     var main = new Scaleway.VpcGatewayNetwork("main", new()
    ///     {
    ///         GatewayId = pg01.Id,
    ///         PrivateNetworkId = pn01.Id,
    ///         EnableDhcp = false,
    ///         EnableMasquerade = true,
    ///         StaticAddress = "192.168.1.42/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gateway network can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork")]
    public partial class VpcGatewayNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Remove DHCP config on this network on destroy. It requires DHCP id.
        /// </summary>
        [Output("cleanupDhcp")]
        public Output<bool?> CleanupDhcp { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the gateway network.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the public gateway DHCP config. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Output("dhcpId")]
        public Output<string?> DhcpId { get; private set; } = null!;

        /// <summary>
        /// Enable DHCP config on this network. It requires DHCP id.
        /// </summary>
        [Output("enableDhcp")]
        public Output<bool?> EnableDhcp { get; private set; } = null!;

        /// <summary>
        /// Enable masquerade on this network
        /// </summary>
        [Output("enableMasquerade")]
        public Output<bool?> EnableMasquerade { get; private set; } = null!;

        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Output("ipamConfigs")]
        public Output<ImmutableArray<Outputs.VpcGatewayNetworkIpamConfig>> IpamConfigs { get; private set; } = null!;

        /// <summary>
        /// The mac address of the creation of the gateway network.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Output("privateNetworkId")]
        public Output<string> PrivateNetworkId { get; private set; } = null!;

        /// <summary>
        /// Enable DHCP config on this network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Output("staticAddress")]
        public Output<string> StaticAddress { get; private set; } = null!;

        /// <summary>
        /// The status of the Public Gateway's connection to the Private Network.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the gateway network.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the gateway network should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcGatewayNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcGatewayNetwork(string name, VpcGatewayNetworkArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, args ?? new VpcGatewayNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcGatewayNetwork(string name, Input<string> id, VpcGatewayNetworkState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcGatewayNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcGatewayNetwork Get(string name, Input<string> id, VpcGatewayNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcGatewayNetwork(name, id, state, options);
        }
    }

    public sealed class VpcGatewayNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remove DHCP config on this network on destroy. It requires DHCP id.
        /// </summary>
        [Input("cleanupDhcp")]
        public Input<bool>? CleanupDhcp { get; set; }

        /// <summary>
        /// The ID of the public gateway DHCP config. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Input("dhcpId")]
        public Input<string>? DhcpId { get; set; }

        /// <summary>
        /// Enable DHCP config on this network. It requires DHCP id.
        /// </summary>
        [Input("enableDhcp")]
        public Input<bool>? EnableDhcp { get; set; }

        /// <summary>
        /// Enable masquerade on this network
        /// </summary>
        [Input("enableMasquerade")]
        public Input<bool>? EnableMasquerade { get; set; }

        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        [Input("ipamConfigs")]
        private InputList<Inputs.VpcGatewayNetworkIpamConfigArgs>? _ipamConfigs;

        /// <summary>
        /// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        public InputList<Inputs.VpcGatewayNetworkIpamConfigArgs> IpamConfigs
        {
            get => _ipamConfigs ?? (_ipamConfigs = new InputList<Inputs.VpcGatewayNetworkIpamConfigArgs>());
            set => _ipamConfigs = value;
        }

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        /// <summary>
        /// Enable DHCP config on this network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Input("staticAddress")]
        public Input<string>? StaticAddress { get; set; }

        /// <summary>
        /// `zone`) The zone in which the gateway network should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcGatewayNetworkArgs()
        {
        }
        public static new VpcGatewayNetworkArgs Empty => new VpcGatewayNetworkArgs();
    }

    public sealed class VpcGatewayNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remove DHCP config on this network on destroy. It requires DHCP id.
        /// </summary>
        [Input("cleanupDhcp")]
        public Input<bool>? CleanupDhcp { get; set; }

        /// <summary>
        /// The date and time of the creation of the gateway network.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The ID of the public gateway DHCP config. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Input("dhcpId")]
        public Input<string>? DhcpId { get; set; }

        /// <summary>
        /// Enable DHCP config on this network. It requires DHCP id.
        /// </summary>
        [Input("enableDhcp")]
        public Input<bool>? EnableDhcp { get; set; }

        /// <summary>
        /// Enable masquerade on this network
        /// </summary>
        [Input("enableMasquerade")]
        public Input<bool>? EnableMasquerade { get; set; }

        /// <summary>
        /// The ID of the public gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        [Input("ipamConfigs")]
        private InputList<Inputs.VpcGatewayNetworkIpamConfigGetArgs>? _ipamConfigs;

        /// <summary>
        /// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        public InputList<Inputs.VpcGatewayNetworkIpamConfigGetArgs> IpamConfigs
        {
            get => _ipamConfigs ?? (_ipamConfigs = new InputList<Inputs.VpcGatewayNetworkIpamConfigGetArgs>());
            set => _ipamConfigs = value;
        }

        /// <summary>
        /// The mac address of the creation of the gateway network.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// Enable DHCP config on this network. Only one of `dhcp_id`, `static_address` and `ipam_config` should be specified.
        /// </summary>
        [Input("staticAddress")]
        public Input<string>? StaticAddress { get; set; }

        /// <summary>
        /// The status of the Public Gateway's connection to the Private Network.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date and time of the last update of the gateway network.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// `zone`) The zone in which the gateway network should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcGatewayNetworkState()
        {
        }
        public static new VpcGatewayNetworkState Empty => new VpcGatewayNetworkState();
    }
}
