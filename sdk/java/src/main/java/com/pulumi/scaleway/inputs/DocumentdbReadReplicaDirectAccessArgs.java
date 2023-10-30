// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DocumentdbReadReplicaDirectAccessArgs extends com.pulumi.resources.ResourceArgs {

    public static final DocumentdbReadReplicaDirectAccessArgs Empty = new DocumentdbReadReplicaDirectAccessArgs();

    /**
     * The ID of the endpoint of the read replica.
     * 
     */
    @Import(name="endpointId")
    private @Nullable Output<String> endpointId;

    /**
     * @return The ID of the endpoint of the read replica.
     * 
     */
    public Optional<Output<String>> endpointId() {
        return Optional.ofNullable(this.endpointId);
    }

    /**
     * Hostname of the endpoint. Only one of ip and hostname may be set.
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return Hostname of the endpoint. Only one of ip and hostname may be set.
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * Name of the endpoint.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the endpoint.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * TCP port of the endpoint.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return TCP port of the endpoint.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private DocumentdbReadReplicaDirectAccessArgs() {}

    private DocumentdbReadReplicaDirectAccessArgs(DocumentdbReadReplicaDirectAccessArgs $) {
        this.endpointId = $.endpointId;
        this.hostname = $.hostname;
        this.ip = $.ip;
        this.name = $.name;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DocumentdbReadReplicaDirectAccessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DocumentdbReadReplicaDirectAccessArgs $;

        public Builder() {
            $ = new DocumentdbReadReplicaDirectAccessArgs();
        }

        public Builder(DocumentdbReadReplicaDirectAccessArgs defaults) {
            $ = new DocumentdbReadReplicaDirectAccessArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpointId The ID of the endpoint of the read replica.
         * 
         * @return builder
         * 
         */
        public Builder endpointId(@Nullable Output<String> endpointId) {
            $.endpointId = endpointId;
            return this;
        }

        /**
         * @param endpointId The ID of the endpoint of the read replica.
         * 
         * @return builder
         * 
         */
        public Builder endpointId(String endpointId) {
            return endpointId(Output.of(endpointId));
        }

        /**
         * @param hostname Hostname of the endpoint. Only one of ip and hostname may be set.
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname Hostname of the endpoint. Only one of ip and hostname may be set.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param ip IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IPv4 address of the endpoint (IP address). Only one of ip and hostname may be set.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param name Name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param port TCP port of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port TCP port of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public DocumentdbReadReplicaDirectAccessArgs build() {
            return $;
        }
    }

}
