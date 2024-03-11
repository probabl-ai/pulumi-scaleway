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

    public sealed class ObjectBucketAclAccessControlPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("grants")]
        private InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantGetArgs>? _grants;
        public InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantGetArgs> Grants
        {
            get => _grants ?? (_grants = new InputList<Inputs.ObjectBucketAclAccessControlPolicyGrantGetArgs>());
            set => _grants = value;
        }

        /// <summary>
        /// Configuration block of the bucket project owner's display organization ID.
        /// </summary>
        [Input("owner", required: true)]
        public Input<Inputs.ObjectBucketAclAccessControlPolicyOwnerGetArgs> Owner { get; set; } = null!;

        public ObjectBucketAclAccessControlPolicyGetArgs()
        {
        }
        public static new ObjectBucketAclAccessControlPolicyGetArgs Empty => new ObjectBucketAclAccessControlPolicyGetArgs();
    }
}
