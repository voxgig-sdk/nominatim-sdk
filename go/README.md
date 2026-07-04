# Nominatim Golang SDK



The Golang SDK for the Nominatim API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/nominatim-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/nominatim-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/nominatim-sdk/go=../nominatim-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/nominatim-sdk/go"
)

func main() {
    client := sdk.New()

    // List addresslookup records — the value is the array of records itself.
    addresslookups, err := client.AddressLookup(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range addresslookups.([]any) {
        fmt.Println(item)
    }
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

addresslookup, err := client.AddressLookup(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(addresslookup) // the loaded mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewNominatimSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOMINATIM_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewNominatimSDK

```go
func NewNominatimSDK(options map[string]any) *NominatimSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *NominatimSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `AddressLookup` | `(data map[string]any) NominatimEntity` | Create an AddressLookup entity instance. |
| `Administrative` | `(data map[string]any) NominatimEntity` | Create an Administrative entity instance. |
| `Debug` | `(data map[string]any) NominatimEntity` | Create a Debug entity instance. |
| `Reverse` | `(data map[string]any) NominatimEntity` | Create a Reverse entity instance. |
| `Search` | `(data map[string]any) NominatimEntity` | Create a Search entity instance. |
| `ServerStatus` | `(data map[string]any) NominatimEntity` | Create a ServerStatus entity instance. |

### Entity interface (NominatimEntity)

All entities implement the `NominatimEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    addresslookup, err := client.AddressLookup(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil { /* handle */ }
    // addresslookup is the loaded record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### AddressLookup

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"boundingbox"` |  |
| `"class"` |  |
| `"display_name"` |  |
| `"importance"` |  |
| `"lat"` |  |
| `"licence"` |  |
| `"lon"` |  |
| `"osm_id"` |  |
| `"osm_type"` |  |
| `"place_id"` |  |
| `"type"` |  |

Operations: List.

API path: `/lookup`

#### Administrative

| Field | Description |
| --- | --- |
| `"class"` |  |
| `"country_code"` |  |
| `"errormessage"` |  |
| `"name"` |  |
| `"osm_id"` |  |
| `"osm_type"` |  |
| `"place_id"` |  |
| `"type"` |  |
| `"updated"` |  |

Operations: List.

API path: `/polygons`

#### Debug

| Field | Description |
| --- | --- |
| `"addresstag"` |  |
| `"admin_level"` |  |
| `"calculated_importance"` |  |
| `"calculated_postcode"` |  |
| `"calculated_wikipedia"` |  |
| `"category"` |  |
| `"centroid"` |  |
| `"country_code"` |  |
| `"extratag"` |  |
| `"geometry"` |  |
| `"housenumber"` |  |
| `"importance"` |  |
| `"indexed_date"` |  |
| `"isarea"` |  |
| `"localname"` |  |
| `"name"` |  |
| `"osm_id"` |  |
| `"osm_type"` |  |
| `"parent_place_id"` |  |
| `"place_id"` |  |
| `"rank_address"` |  |
| `"rank_search"` |  |
| `"type"` |  |

Operations: Load.

API path: `/details`

#### Reverse

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"boundingbox"` |  |
| `"display_name"` |  |
| `"lat"` |  |
| `"licence"` |  |
| `"lon"` |  |
| `"osm_id"` |  |
| `"osm_type"` |  |
| `"place_id"` |  |

Operations: List.

API path: `/reverse`

#### Search

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"boundingbox"` |  |
| `"class"` |  |
| `"display_name"` |  |
| `"icon"` |  |
| `"importance"` |  |
| `"lat"` |  |
| `"licence"` |  |
| `"lon"` |  |
| `"osm_id"` |  |
| `"osm_type"` |  |
| `"place_id"` |  |
| `"type"` |  |

Operations: List.

API path: `/search`

#### ServerStatus

| Field | Description |
| --- | --- |
| `"data_updated"` |  |
| `"database_version"` |  |
| `"message"` |  |
| `"software_version"` |  |
| `"status"` |  |

Operations: Load.

API path: `/status`



## Entities


### AddressLookup

Create an instance: `address_lookup := client.AddressLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `class` | ``$STRING`` |  |
| `display_name` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: List

```go
address_lookups, err := client.AddressLookup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(address_lookups) // the array of records
```


### Administrative

Create an instance: `administrative := client.Administrative(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | ``$STRING`` |  |
| `country_code` | ``$STRING`` |  |
| `errormessage` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |
| `updated` | ``$STRING`` |  |

#### Example: List

```go
administratives, err := client.Administrative(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(administratives) // the array of records
```


### Debug

Create an instance: `debug := client.Debug(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstag` | ``$OBJECT`` |  |
| `admin_level` | ``$INTEGER`` |  |
| `calculated_importance` | ``$NUMBER`` |  |
| `calculated_postcode` | ``$STRING`` |  |
| `calculated_wikipedia` | ``$STRING`` |  |
| `category` | ``$STRING`` |  |
| `centroid` | ``$OBJECT`` |  |
| `country_code` | ``$STRING`` |  |
| `extratag` | ``$OBJECT`` |  |
| `geometry` | ``$OBJECT`` |  |
| `housenumber` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `indexed_date` | ``$STRING`` |  |
| `isarea` | ``$BOOLEAN`` |  |
| `localname` | ``$STRING`` |  |
| `name` | ``$OBJECT`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `parent_place_id` | ``$INTEGER`` |  |
| `place_id` | ``$INTEGER`` |  |
| `rank_address` | ``$INTEGER`` |  |
| `rank_search` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: Load

```go
debug, err := client.Debug(nil).Load(map[string]any{"id": "debug_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(debug) // the loaded record
```


### Reverse

Create an instance: `reverse := client.Reverse(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `display_name` | ``$STRING`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |

#### Example: List

```go
reverses, err := client.Reverse(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(reverses) // the array of records
```


### Search

Create an instance: `search := client.Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | ``$OBJECT`` |  |
| `boundingbox` | ``$ARRAY`` |  |
| `class` | ``$STRING`` |  |
| `display_name` | ``$STRING`` |  |
| `icon` | ``$STRING`` |  |
| `importance` | ``$NUMBER`` |  |
| `lat` | ``$STRING`` |  |
| `licence` | ``$STRING`` |  |
| `lon` | ``$STRING`` |  |
| `osm_id` | ``$INTEGER`` |  |
| `osm_type` | ``$STRING`` |  |
| `place_id` | ``$INTEGER`` |  |
| `type` | ``$STRING`` |  |

#### Example: List

```go
searchs, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(searchs) // the array of records
```


### ServerStatus

Create an instance: `server_status := client.ServerStatus(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | ``$STRING`` |  |
| `database_version` | ``$STRING`` |  |
| `message` | ``$STRING`` |  |
| `software_version` | ``$STRING`` |  |
| `status` | ``$INTEGER`` |  |

#### Example: Load

```go
server_status, err := client.ServerStatus(nil).Load(map[string]any{"id": "server_status_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(server_status) // the loaded record
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/nominatim-sdk/go/
├── nominatim.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/nominatim-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
addresslookup := client.AddressLookup(nil)
addresslookup.Load(map[string]any{"id": "example_id"}, nil)

// addresslookup.Data() now returns the loaded addresslookup data
// addresslookup.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
