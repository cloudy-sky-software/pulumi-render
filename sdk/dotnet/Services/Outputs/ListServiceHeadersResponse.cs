// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class ListServiceHeadersResponse
    {
        public readonly string? Cursor;
        /// <summary>
        /// A service header object
        /// </summary>
        public readonly Outputs.ServiceHeader? Header;

        [OutputConstructor]
        private ListServiceHeadersResponse(
            string? cursor,

            Outputs.ServiceHeader? header)
        {
            Cursor = cursor;
            Header = header;
        }
    }
}