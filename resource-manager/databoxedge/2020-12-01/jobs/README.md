
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/jobs` Documentation

The `jobs` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2020-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/jobs"
```


### Client Initialization

```go
client := jobs.NewJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobsClient.Get`

```go
ctx := context.TODO()
id := jobs.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "jobValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
