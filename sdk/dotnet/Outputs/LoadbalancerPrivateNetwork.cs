// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Outputs
{

    [OutputType]
    public sealed class LoadbalancerPrivateNetwork
    {
        public readonly bool? DhcpConfig;
        public readonly string PrivateNetworkId;
        public readonly ImmutableArray<string> StaticConfigs;
        /// <summary>
        /// The Private Network attachment status
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private LoadbalancerPrivateNetwork(
            bool? dhcpConfig,

            string privateNetworkId,

            ImmutableArray<string> staticConfigs,

            string? status,

            string? zone)
        {
            DhcpConfig = dhcpConfig;
            PrivateNetworkId = privateNetworkId;
            StaticConfigs = staticConfigs;
            Status = status;
            Zone = zone;
        }
    }
}
