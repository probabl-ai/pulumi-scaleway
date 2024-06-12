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

    public sealed class LoadbalancerAclActionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("redirects")]
        private InputList<Inputs.LoadbalancerAclActionRedirectGetArgs>? _redirects;

        /// <summary>
        /// Redirect parameters when using an ACL with `redirect` action.
        /// </summary>
        public InputList<Inputs.LoadbalancerAclActionRedirectGetArgs> Redirects
        {
            get => _redirects ?? (_redirects = new InputList<Inputs.LoadbalancerAclActionRedirectGetArgs>());
            set => _redirects = value;
        }

        /// <summary>
        /// The action type. Possible values are: `allow` or `deny` or `redirect`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LoadbalancerAclActionGetArgs()
        {
        }
        public static new LoadbalancerAclActionGetArgs Empty => new LoadbalancerAclActionGetArgs();
    }
}
