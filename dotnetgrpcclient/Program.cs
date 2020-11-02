using System;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;

namespace dotnetgrpcclient
{
	class Program
	{
		static async Task Main(string[] args)
		{
			var httpClientHandler = new HttpClientHandler
			{
				ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator
			};

			var subdirectoryHandler = new SubdirectoryHandler(httpClientHandler, "/grpc");
			var httpClient = new HttpClient(subdirectoryHandler);

			using var channel = GrpcChannel.ForAddress("https://localhost.test:443/grpc/", new GrpcChannelOptions
			{
				HttpClient = httpClient
			});

			var client = new global::going.Greeter.GreeterClient(channel);

			var response = await client.SayHelloAsync(new going.HelloRequest() { Name = "dotnet" }).ResponseAsync.ConfigureAwait(false);

			Console.WriteLine(response.ToString());
		}
	}

	public class SubdirectoryHandler : DelegatingHandler
	{
		private readonly string path;

		public SubdirectoryHandler(HttpMessageHandler innerHandler, string path)
		 : base(innerHandler)
		{
			this.path = path;
		}

		protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken cancellationToken)
		{
			var url = $"{request.RequestUri.Scheme}://{request.RequestUri.Host}{path}{request.RequestUri.AbsolutePath}";
			request.RequestUri = new Uri(url, UriKind.Absolute);
			return base.SendAsync(request, cancellationToken);
		}
	}
}