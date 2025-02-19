
## `github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/hosts` Documentation

The `hosts` SDK allows for interaction with Azure Resource Manager `datadog` (API Version `2023-10-20`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/hosts"
```


### Client Initialization

```go
client := hosts.NewHostsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HostsClient.MonitorsListHosts`

```go
ctx := context.TODO()
id := hosts.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.MonitorsListHosts(ctx, id)` can be used to do batched pagination
items, err := client.MonitorsListHostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
