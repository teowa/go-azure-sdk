
## `github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/experiments` Documentation

The `experiments` SDK allows for interaction with the Azure Resource Manager Service `chaosstudio` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2022-07-01-preview/experiments"
```


### Client Initialization

```go
client := experiments.NewExperimentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExperimentsClient.Cancel`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

payload := experiments.Experiment{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.Delete`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.Get`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.GetExecutionDetails`

```go
ctx := context.TODO()
id := experiments.NewExecutionDetailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue", "executionDetailsIdValue")

read, err := client.GetExecutionDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.GetStatus`

```go
ctx := context.TODO()
id := experiments.NewStatusID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue", "statusIdValue")

read, err := client.GetStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentsClient.List`

```go
ctx := context.TODO()
id := experiments.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id, experiments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, experiments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExperimentsClient.ListAll`

```go
ctx := context.TODO()
id := experiments.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListAll(ctx, id, experiments.DefaultListAllOperationOptions())` can be used to do batched pagination
items, err := client.ListAllComplete(ctx, id, experiments.DefaultListAllOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExperimentsClient.ListAllStatuses`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

// alternatively `client.ListAllStatuses(ctx, id)` can be used to do batched pagination
items, err := client.ListAllStatusesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExperimentsClient.ListExecutionDetails`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

// alternatively `client.ListExecutionDetails(ctx, id)` can be used to do batched pagination
items, err := client.ListExecutionDetailsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExperimentsClient.Start`

```go
ctx := context.TODO()
id := experiments.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentValue")

read, err := client.Start(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
