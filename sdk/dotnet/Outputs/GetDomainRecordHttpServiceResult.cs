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
    public sealed class GetDomainRecordHttpServiceResult
    {
        /// <summary>
        /// IPs to check
        /// </summary>
        public readonly ImmutableArray<string> Ips;
        /// <summary>
        /// Text to search
        /// </summary>
        public readonly string MustContain;
        /// <summary>
        /// Strategy to return an IP from the IPs list
        /// </summary>
        public readonly string Strategy;
        /// <summary>
        /// URL to match the must_contain text to validate an IP
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// User-agent used when checking the URL
        /// </summary>
        public readonly string UserAgent;

        [OutputConstructor]
        private GetDomainRecordHttpServiceResult(
            ImmutableArray<string> ips,

            string mustContain,

            string strategy,

            string url,

            string userAgent)
        {
            Ips = ips;
            MustContain = mustContain;
            Strategy = strategy;
            Url = url;
            UserAgent = userAgent;
        }
    }
}
