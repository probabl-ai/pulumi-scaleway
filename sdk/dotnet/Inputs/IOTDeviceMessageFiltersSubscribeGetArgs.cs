// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Inputs
{

    public sealed class IOTDeviceMessageFiltersSubscribeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Same as publish rules.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// Same as publish rules.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public IOTDeviceMessageFiltersSubscribeGetArgs()
        {
        }
    }
}
