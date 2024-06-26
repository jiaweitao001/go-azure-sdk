
## `github.com/hashicorp/go-azure-sdk/resource-manager/devcenter/2024-02-01/devboxdefinitions` Documentation

The `devboxdefinitions` SDK allows for interaction with the Azure Resource Manager Service `devcenter` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devcenter/2024-02-01/devboxdefinitions"
```


### Client Initialization

```go
client := devboxdefinitions.NewDevBoxDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DevBoxDefinitionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevCenterDevBoxDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "devBoxDefinitionValue")

payload := devboxdefinitions.DevBoxDefinition{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevBoxDefinitionsClient.Delete`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevCenterDevBoxDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "devBoxDefinitionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DevBoxDefinitionsClient.Get`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevCenterDevBoxDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "devBoxDefinitionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevBoxDefinitionsClient.GetByProject`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevBoxDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "projectValue", "devBoxDefinitionValue")

read, err := client.GetByProject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevBoxDefinitionsClient.ListByDevCenter`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue")

// alternatively `client.ListByDevCenter(ctx, id)` can be used to do batched pagination
items, err := client.ListByDevCenterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevBoxDefinitionsClient.ListByProject`

```go
ctx := context.TODO()
id := devboxdefinitions.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "projectValue")

// alternatively `client.ListByProject(ctx, id)` can be used to do batched pagination
items, err := client.ListByProjectComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevBoxDefinitionsClient.Update`

```go
ctx := context.TODO()
id := devboxdefinitions.NewDevCenterDevBoxDefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "devCenterValue", "devBoxDefinitionValue")

payload := devboxdefinitions.DevBoxDefinitionUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
