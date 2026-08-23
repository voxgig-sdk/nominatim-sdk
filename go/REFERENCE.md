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
| `address` | `map[string]any` | No | Address breakdown |
| `boundingbox` | `[]any` | No | Bounding box coordinates |
| `class` | `string` | No | Main OSM tag key |
| `display_name` | `string` | No | Full comma-separated address |
| `importance` | `float64` | No | Computed importance rank |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `int` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `int` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |

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
| `class` | `string` | No | Main OSM tag key |
| `country_code` | `string` | No | Country code |
| `errormessage` | `string` | No | Error message describing the polygon issue |
| `name` | `string` | No | Name of the object |
| `osm_id` | `int` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (way, relation) |
| `place_id` | `int` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |
| `updated` | `string` | No | Last update timestamp |

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
| `addresstags` | `map[string]any` | No | Address tags |
| `admin_level` | `int` | No | Administrative level |
| `calculated_importance` | `float64` | No | Calculated importance |
| `calculated_postcode` | `string` | No | Calculated postcode |
| `calculated_wikipedia` | `string` | No | Wikipedia reference |
| `category` | `string` | No | Main OSM tag key |
| `centroid` | `map[string]any` | No | Centroid coordinates |
| `country_code` | `string` | No | Country code |
| `extratags` | `map[string]any` | No | Extra OSM tags |
| `geometry` | `map[string]any` | No | Geometry information |
| `housenumber` | `string` | No | House number |
| `importance` | `float64` | No | Computed importance rank |
| `indexed_date` | `string` | No | Date when the object was indexed |
| `isarea` | `bool` | No | Whether the object is an area |
| `localname` | `string` | No | Local name |
| `names` | `map[string]any` | No | All available names |
| `osm_id` | `int` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `parent_place_id` | `int` | No | Parent place ID |
| `place_id` | `int` | No | Unique identifier for the place |
| `rank_address` | `int` | No | Address rank |
| `rank_search` | `int` | No | Search rank |
| `type` | `string` | No | Main OSM tag value |

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
| `address` | `map[string]any` | No | Address breakdown |
| `boundingbox` | `[]any` | No | Bounding box coordinates |
| `display_name` | `string` | No | Full comma-separated address |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `int` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `int` | No | Unique identifier for the place |

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
| `address` | `map[string]any` | No | Address breakdown |
| `boundingbox` | `[]any` | No | Bounding box coordinates |
| `class` | `string` | No | Main OSM tag key |
| `display_name` | `string` | No | Full comma-separated address |
| `icon` | `string` | No | URL of icon representing the place |
| `importance` | `float64` | No | Computed importance rank |
| `lat` | `string` | No | Latitude |
| `licence` | `string` | No | License information |
| `lon` | `string` | No | Longitude |
| `osm_id` | `int` | No | OSM object ID |
| `osm_type` | `string` | No | OSM type (node, way, relation) |
| `place_id` | `int` | No | Unique identifier for the place |
| `type` | `string` | No | Main OSM tag value |

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
| `data_updated` | `string` | No | Timestamp when the database was last updated |
| `database_version` | `string` | No | Database version |
| `message` | `string` | No | Status message |
| `software_version` | `string` | No | Nominatim software version |
| `status` | `int` | No | Status code (0 = OK) |

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

