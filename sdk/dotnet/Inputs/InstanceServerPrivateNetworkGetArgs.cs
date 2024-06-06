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

    public sealed class InstanceServerPrivateNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// MAC address of the NIC
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The Private Network ID
        /// </summary>
        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        /// <summary>
        /// The ID of the NIC
        /// </summary>
        [Input("pnicId")]
        public Input<string>? PnicId { get; set; }

        /// <summary>
        /// The private NIC state
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceServerPrivateNetworkGetArgs()
        {
        }
        public static new InstanceServerPrivateNetworkGetArgs Empty => new InstanceServerPrivateNetworkGetArgs();
    }
}
