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

    public sealed class LoadbalancerCertificateLetsencryptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Main domain of the certificate
        /// </summary>
        [Input("commonName", required: true)]
        public Input<string> CommonName { get; set; } = null!;

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;

        /// <summary>
        /// The alternative domain names of the certificate
        /// </summary>
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        public LoadbalancerCertificateLetsencryptArgs()
        {
        }
        public static new LoadbalancerCertificateLetsencryptArgs Empty => new LoadbalancerCertificateLetsencryptArgs();
    }
}
