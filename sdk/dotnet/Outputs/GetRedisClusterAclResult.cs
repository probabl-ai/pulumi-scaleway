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
    public sealed class GetRedisClusterAclResult
    {
        /// <summary>
        /// Description of the rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the Redis cluster.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv4 network address of the rule (IP network in a CIDR format).
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private GetRedisClusterAclResult(
            string description,

            string id,

            string ip)
        {
            Description = description;
            Id = id;
            Ip = ip;
        }
    }
}
