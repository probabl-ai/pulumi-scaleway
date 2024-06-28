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

    public sealed class DocumentdbPrivateNetworkEndpointPrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hostname of your endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The private network ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The IP of your private network service.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM) service if not set.
        /// </summary>
        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// The name of your private service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of your private service.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The zone of your endpoint.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DocumentdbPrivateNetworkEndpointPrivateNetworkGetArgs()
        {
        }
        public static new DocumentdbPrivateNetworkEndpointPrivateNetworkGetArgs Empty => new DocumentdbPrivateNetworkEndpointPrivateNetworkGetArgs();
    }
}
