// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Scaleway
{
    public static class GetConfig
    {
        public static Task<GetConfigResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigResult>("scaleway:index/getConfig:getConfig", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetConfigResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigResult>("scaleway:index/getConfig:getConfig", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetConfigResult
    {
        public readonly string AccessKey;
        public readonly string AccessKeySource;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProjectId;
        public readonly string ProjectIdSource;
        public readonly string Region;
        public readonly string RegionSource;
        public readonly string SecretKey;
        public readonly string SecretKeySource;
        public readonly string Zone;
        public readonly string ZoneSource;

        [OutputConstructor]
        private GetConfigResult(
            string accessKey,

            string accessKeySource,

            string id,

            string projectId,

            string projectIdSource,

            string region,

            string regionSource,

            string secretKey,

            string secretKeySource,

            string zone,

            string zoneSource)
        {
            AccessKey = accessKey;
            AccessKeySource = accessKeySource;
            Id = id;
            ProjectId = projectId;
            ProjectIdSource = projectIdSource;
            Region = region;
            RegionSource = regionSource;
            SecretKey = secretKey;
            SecretKeySource = secretKeySource;
            Zone = zone;
            ZoneSource = zoneSource;
        }
    }
}
