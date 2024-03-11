// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway.Outputs
{

    [OutputType]
    public sealed class RedisClusterAcl
    {
        /// <summary>
        /// A text describing this rule. Default description: `Allow IP`
        /// 
        /// &gt; The `acl` conflict with `private_network`. Only one should be specified.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The UUID of the Private Network resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The ip range to whitelist
        /// in [CIDR notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation)
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private RedisClusterAcl(
            string? description,

            string? id,

            string ip)
        {
            Description = description;
            Id = id;
            Ip = ip;
        }
    }
}
