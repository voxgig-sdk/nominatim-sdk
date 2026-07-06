# Nominatim Ruby SDK Reference

Complete API reference for the Nominatim Ruby SDK.


## NominatimSDK

### Constructor

```ruby
require_relative 'Nominatim_sdk'

client = NominatimSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = NominatimSDK.test
```


### Instance Methods

#### `AddressLookup(data = nil)`

Create a new `AddressLookup` entity instance. Pass `nil` for no initial data.

#### `Administrative(data = nil)`

Create a new `Administrative` entity instance. Pass `nil` for no initial data.

#### `Debug(data = nil)`

Create a new `Debug` entity instance. Pass `nil` for no initial data.

#### `Reverse(data = nil)`

Create a new `Reverse` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `ServerStatus(data = nil)`

Create a new `ServerStatus` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AddressLookupEntity

```ruby
address_lookup = client.AddressLookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No |  |
| `boundingbox` | `Array` | No |  |
| `class` | `String` | No |  |
| `display_name` | `String` | No |  |
| `importance` | `Float` | No |  |
| `lat` | `String` | No |  |
| `licence` | `String` | No |  |
| `lon` | `String` | No |  |
| `osm_id` | `Integer` | No |  |
| `osm_type` | `String` | No |  |
| `place_id` | `Integer` | No |  |
| `type` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.AddressLookup.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AdministrativeEntity

```ruby
administrative = client.Administrative
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | `String` | No |  |
| `country_code` | `String` | No |  |
| `errormessage` | `String` | No |  |
| `name` | `String` | No |  |
| `osm_id` | `Integer` | No |  |
| `osm_type` | `String` | No |  |
| `place_id` | `Integer` | No |  |
| `type` | `String` | No |  |
| `updated` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Administrative.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DebugEntity

```ruby
debug = client.Debug
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstag` | `Hash` | No |  |
| `admin_level` | `Integer` | No |  |
| `calculated_importance` | `Float` | No |  |
| `calculated_postcode` | `String` | No |  |
| `calculated_wikipedia` | `String` | No |  |
| `category` | `String` | No |  |
| `centroid` | `Hash` | No |  |
| `country_code` | `String` | No |  |
| `extratag` | `Hash` | No |  |
| `geometry` | `Hash` | No |  |
| `housenumber` | `String` | No |  |
| `importance` | `Float` | No |  |
| `indexed_date` | `String` | No |  |
| `isarea` | `Boolean` | No |  |
| `localname` | `String` | No |  |
| `name` | `Hash` | No |  |
| `osm_id` | `Integer` | No |  |
| `osm_type` | `String` | No |  |
| `parent_place_id` | `Integer` | No |  |
| `place_id` | `Integer` | No |  |
| `rank_address` | `Integer` | No |  |
| `rank_search` | `Integer` | No |  |
| `type` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Debug.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DebugEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ReverseEntity

```ruby
reverse = client.Reverse
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No |  |
| `boundingbox` | `Array` | No |  |
| `display_name` | `String` | No |  |
| `lat` | `String` | No |  |
| `licence` | `String` | No |  |
| `lon` | `String` | No |  |
| `osm_id` | `Integer` | No |  |
| `osm_type` | `String` | No |  |
| `place_id` | `Integer` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Reverse.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Hash` | No |  |
| `boundingbox` | `Array` | No |  |
| `class` | `String` | No |  |
| `display_name` | `String` | No |  |
| `icon` | `String` | No |  |
| `importance` | `Float` | No |  |
| `lat` | `String` | No |  |
| `licence` | `String` | No |  |
| `lon` | `String` | No |  |
| `osm_id` | `Integer` | No |  |
| `osm_type` | `String` | No |  |
| `place_id` | `Integer` | No |  |
| `type` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ServerStatusEntity

```ruby
server_status = client.ServerStatus
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | `String` | No |  |
| `database_version` | `String` | No |  |
| `message` | `String` | No |  |
| `software_version` | `String` | No |  |
| `status` | `Integer` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ServerStatus.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = NominatimSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

