// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Inputs
{

    public sealed class DatabaseInstancePrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, the IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// 
        /// &gt; **NOTE:** Please calculate your host IP using cidrhost. Otherwise, let IPAM service
        /// handle the host IP on the network.
        /// 
        /// &gt; **Important:** Updates to `private_network` will recreate the Instance's endpoint
        /// </summary>
        [Input("enableIpam")]
        public Input<bool>? EnableIpam { get; set; }

        /// <summary>
        /// The ID of the endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a CIDR notation. If not set, The IP network address within the private subnet is determined by the IP Address Management (IPAM) service.
        /// </summary>
        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DatabaseInstancePrivateNetworkArgs()
        {
        }
        public static new DatabaseInstancePrivateNetworkArgs Empty => new DatabaseInstancePrivateNetworkArgs();
    }
}
