// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetWebhostingOptionResult
    {
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetWebhostingOptionResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
