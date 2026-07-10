# Nominatim Golang SDK Reference

Complete API reference for the Nominatim Golang SDK.


## NominatimSDK

### Constructor

```go
func NewNominatimSDK(options map[string]any) *NominatimSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *NominatimSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *NominatimSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `AddressLookup(data map[string]any) NominatimEntity`

Create a new `AddressLookup` entity instance. Pass `nil` for no initial data.

#### `Administrative(data map[string]any) NominatimEntity`

Create a new `Administrative` entity instance. Pass `nil` for no initial data.

#### `Debug(data map[string]any) NominatimEntity`

Create a new `Debug` entity instance. Pass `nil` for no initial data.

#### `Reverse(data map[string]any) NominatimEntity`

Create a new `Reverse` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) NominatimEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `ServerStatus(data map[string]any) NominatimEntity`

Create a new `ServerStatus` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AddressLookupEntity

```go
addressLookup := client.AddressLookup(nil)
fmt.Println(addressLookup.GetName()) // "address_lookup"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `map[string]any` | No |  |
| `boundingbox` | `[]any` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `importance` | `float64` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.AddressLookup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AdministrativeEntity

```go
administrative := client.Administrative(nil)
fmt.Println(administrative.GetName()) // "administrative"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | `string` | No |  |
| `country_code` | `string` | No |  |
| `errormessage` | `string` | No |  |
| `name` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |
| `type` | `string` | No |  |
| `updated` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Administrative(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DebugEntity

```go
debug := client.Debug(nil)
fmt.Println(debug.GetName()) // "debug"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstag` | `map[string]any` | No |  |
| `admin_level` | `int` | No |  |
| `calculated_importance` | `float64` | No |  |
| `calculated_postcode` | `string` | No |  |
| `calculated_wikipedia` | `string` | No |  |
| `category` | `string` | No |  |
| `centroid` | `map[string]any` | No |  |
| `country_code` | `string` | No |  |
| `extratag` | `map[string]any` | No |  |
| `geometry` | `map[string]any` | No |  |
| `housenumber` | `string` | No |  |
| `importance` | `float64` | No |  |
| `indexed_date` | `string` | No |  |
| `isarea` | `bool` | No |  |
| `localname` | `string` | No |  |
| `name` | `map[string]any` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `parent_place_id` | `int` | No |  |
| `place_id` | `int` | No |  |
| `rank_address` | `int` | No |  |
| `rank_search` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Debug(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DebugEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReverseEntity

```go
reverse := client.Reverse(nil)
fmt.Println(reverse.GetName()) // "reverse"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `map[string]any` | No |  |
| `boundingbox` | `[]any` | No |  |
| `display_name` | `string` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Reverse(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
fmt.Println(search.GetName()) // "search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `map[string]any` | No |  |
| `boundingbox` | `[]any` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `icon` | `string` | No |  |
| `importance` | `float64` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ServerStatusEntity

```go
serverStatus := client.ServerStatus(nil)
fmt.Println(serverStatus.GetName()) // "server_status"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | `string` | No |  |
| `database_version` | `string` | No |  |
| `message` | `string` | No |  |
| `software_version` | `string` | No |  |
| `status` | `int` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ServerStatus(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewNominatimSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

