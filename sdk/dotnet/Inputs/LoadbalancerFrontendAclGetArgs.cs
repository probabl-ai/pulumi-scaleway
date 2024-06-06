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

    public sealed class LoadbalancerFrontendAclGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.LoadbalancerFrontendAclActionGetArgs> Action { get; set; } = null!;

        /// <summary>
        /// IsDate and time of ACL's creation (RFC 3339 format)
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the ACL
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.LoadbalancerFrontendAclMatchGetArgs> Match { get; set; } = null!;

        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IsDate and time of ACL's update (RFC 3339 format)
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public LoadbalancerFrontendAclGetArgs()
        {
        }
        public static new LoadbalancerFrontendAclGetArgs Empty => new LoadbalancerFrontendAclGetArgs();
    }
}
