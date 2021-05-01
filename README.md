# vmaas-cmp-go-sdk

This package provides the official [Go](https://golang.org/) library for the [CMP API](https://supreme-happiness-313dad4f.pages.github.io/).

This is being developed in conjunction with the [VMaaS Terraform Provider](https://github.com/hpe-hcss/vmaas-terraform-resources).

## Setup

Install Go, export environment variables, go get the morpheus package and begin executing requests.

## Requirements

* [Go](https://golang.org/dl/) | 1.13

## Usage

Here are some examples of how to use `cmp.Client`.

### List Instances

Fetch a list of instances.

```go
import "github.com/hpe-hcss/vmaas-cmp-go-sdk"
cmpClient := cmp.NewClient("https://<short-uuid>.privatecloud.greenlake.hpe.com")
cmpClient.SetAccessToken("048f1eaa-****-447b-****-7f19055b64fd", "", 0, "write")
resp, err := cmpClient.ListInstances(&cmp.Request{})
// parse JSON and fetch the first one by ID
listInstancesResult := resp.Result.(*cmp.ListInstancesResult)
instancesCount := listInstancesResult.Meta.Total
fmt.Sprintf("Found %d Instances.", instancesCount)
```

**NOTE:** This may be simplified so that typecasting the result is not always needed.

## Testing

You can execute the latest tests using:

```sh
go test
```

The above command will (ideally) print results like this:

```
Initializing test client for tfplugin @ https://<short-uuid>.privatecloud.greenlake.hpe.com
PASS
ok      github.com/hpe-hcss/vmaas-cmp-go-sdk   1.098s
```

Running `go test` will fail with a panic right away if you have not yet setup your test environment variables.  

```bash
export CMP_TEST_URL=https://<short-uuid>.privatecloud.greenlake.hpe.com
export CMP_TEST_USERNAME=gotest
export CMP_TEST_PASSWORD=19830B3f489
```
**Be Careful running this test suite**. It creates and destroys data. Never point at any URL other than a test environment. Although, in reality, tests will not modify or destroy any pre-existing data. It could still orphan some test some data, or cause otherwise undesired effects.

You can run an individual test like this:

```sh
go test -run TestGroupsCRUD
```


```bash
go test -v
```

### Code Structure

The main type this package exposes is [Client](../blob/master/client.go), implemented in client.go.  

Each resource is defined in its own file eg. [instances.go](../blob/master/instances.go)  which extends the `Client` type by defining a function for each endpoint the resource has, such as GetInstance(), ListInstances(), CreateInstance(), UpdateInstance, DeleteInstance(), etc. The request and response payload types used by those methods are also defined here.

#### Test Files

Be sure to add a `_test.go` file with unit tests for each new resource that is implemented.