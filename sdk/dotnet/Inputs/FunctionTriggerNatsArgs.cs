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

    public sealed class FunctionTriggerNatsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the mnq nats account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// ID of the project that contain the mnq nats account, defaults to provider's project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Region where the mnq nats account is, defaults to provider's region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The subject to listen to
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        public FunctionTriggerNatsArgs()
        {
        }
        public static new FunctionTriggerNatsArgs Empty => new FunctionTriggerNatsArgs();
    }
}
