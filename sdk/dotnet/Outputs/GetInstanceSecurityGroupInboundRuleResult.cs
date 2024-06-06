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
    public sealed class GetInstanceSecurityGroupInboundRuleResult
    {
        /// <summary>
        /// The action to take when rule match. Possible values are: `accept` or `drop`.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The ip this rule apply to.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The ip range (e.g `192.168.1.0/24`) this rule apply to.
        /// </summary>
        public readonly string IpRange;
        /// <summary>
        /// The port this rule apply to. If no port is specified, rule will apply to all port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The port range (e.g `22-23`) this rule applies to.
        /// If no `port` nor `port_range` are specified, rule will apply to all port.
        /// Only one of `port` and `port_range` should be specified.
        /// </summary>
        public readonly string PortRange;
        /// <summary>
        /// The protocol this rule apply to. Possible values are: `TCP`, `UDP`, `ICMP` or `ANY`.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetInstanceSecurityGroupInboundRuleResult(
            string action,

            string ip,

            string ipRange,

            int port,

            string portRange,

            string protocol)
        {
            Action = action;
            Ip = ip;
            IpRange = ipRange;
            Port = port;
            PortRange = portRange;
            Protocol = protocol;
        }
    }
}
