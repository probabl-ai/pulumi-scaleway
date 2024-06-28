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

    public sealed class VpcPrivateNetworkIpv4SubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network address of the subnet in hexadecimal notation, e.g., '2001:db8::' for a '2001:db8::/64' subnet.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The date and time of the creation of the subnet.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The subnet ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The length of the network prefix, e.g., 64 for a 'ffff:ffff:ffff:ffff::' mask.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// The subnet CIDR.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// The subnet mask expressed in dotted decimal notation, e.g., '255.255.255.0' for a /24 subnet
        /// </summary>
        [Input("subnetMask")]
        public Input<string>? SubnetMask { get; set; }

        /// <summary>
        /// The date and time of the last update of the subnet.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public VpcPrivateNetworkIpv4SubnetArgs()
        {
        }
        public static new VpcPrivateNetworkIpv4SubnetArgs Empty => new VpcPrivateNetworkIpv4SubnetArgs();
    }
}
