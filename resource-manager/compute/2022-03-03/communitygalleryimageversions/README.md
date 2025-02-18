
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/communitygalleryimageversions` Documentation

The `communitygalleryimageversions` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2022-03-03`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-03/communitygalleryimageversions"
```


### Client Initialization

```go
client := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CommunityGalleryImageVersionsClient.Get`

```go
ctx := context.TODO()
id := communitygalleryimageversions.NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CommunityGalleryImageVersionsClient.List`

```go
ctx := context.TODO()
id := communitygalleryimageversions.NewImageID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
