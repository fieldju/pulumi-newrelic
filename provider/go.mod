module github.com/pulumi/pulumi-newrelic/provider/v3

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/newrelic/terraform-provider-newrelic v1.19.1
	github.com/newrelic/terraform-provider-newrelic/v2 v2.2.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.5-0.20200720120417-c472626db64b
	github.com/pulumi/pulumi/sdk/v2 v2.6.2-0.20200715231137-1a10e06588f6
	github.com/vmihailenco/msgpack v4.0.2+incompatible // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9
)
