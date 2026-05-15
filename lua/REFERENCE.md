# Nominatim Lua SDK Reference

Complete API reference for the Nominatim Lua SDK.


## NominatimSDK

### Constructor

```lua
local sdk = require("nominatim_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts, sdkopts)`

Create a test client with mock features active. Both arguments may be `nil`.

```lua
local client = sdk.test(nil, nil)
```


### Instance Methods

#### `AddressLookup(data)`

Create a new `AddressLookup` entity instance. Pass `nil` for no initial data.

#### `Administrative(data)`

Create a new `Administrative` entity instance. Pass `nil` for no initial data.

#### `Debug(data)`

Create a new `Debug` entity instance. Pass `nil` for no initial data.

#### `Reverse(data)`

Create a new `Reverse` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `ServerStatus(data)`

Create a new `ServerStatus` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AddressLookupEntity

```lua
local address_lookup = client:AddressLookup(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$OBJECT`` | No |  |
| `boundingbox` | ``$ARRAY`` | No |  |
| `class` | ``$STRING`` | No |  |
| `display_name` | ``$STRING`` | No |  |
| `importance` | ``$NUMBER`` | No |  |
| `lat` | ``$STRING`` | No |  |
| `licence` | ``$STRING`` | No |  |
| `lon` | ``$STRING`` | No |  |
| `osm_id` | ``$INTEGER`` | No |  |
| `osm_type` | ``$STRING`` | No |  |
| `place_id` | ``$INTEGER`` | No |  |
| `type` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:AddressLookup(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AdministrativeEntity

```lua
local administrative = client:Administrative(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | ``$STRING`` | No |  |
| `country_code` | ``$STRING`` | No |  |
| `errormessage` | ``$STRING`` | No |  |
| `name` | ``$STRING`` | No |  |
| `osm_id` | ``$INTEGER`` | No |  |
| `osm_type` | ``$STRING`` | No |  |
| `place_id` | ``$INTEGER`` | No |  |
| `type` | ``$STRING`` | No |  |
| `updated` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Administrative(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DebugEntity

```lua
local debug = client:Debug(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstag` | ``$OBJECT`` | No |  |
| `admin_level` | ``$INTEGER`` | No |  |
| `calculated_importance` | ``$NUMBER`` | No |  |
| `calculated_postcode` | ``$STRING`` | No |  |
| `calculated_wikipedia` | ``$STRING`` | No |  |
| `category` | ``$STRING`` | No |  |
| `centroid` | ``$OBJECT`` | No |  |
| `country_code` | ``$STRING`` | No |  |
| `extratag` | ``$OBJECT`` | No |  |
| `geometry` | ``$OBJECT`` | No |  |
| `housenumber` | ``$STRING`` | No |  |
| `importance` | ``$NUMBER`` | No |  |
| `indexed_date` | ``$STRING`` | No |  |
| `isarea` | ``$BOOLEAN`` | No |  |
| `localname` | ``$STRING`` | No |  |
| `name` | ``$OBJECT`` | No |  |
| `osm_id` | ``$INTEGER`` | No |  |
| `osm_type` | ``$STRING`` | No |  |
| `parent_place_id` | ``$INTEGER`` | No |  |
| `place_id` | ``$INTEGER`` | No |  |
| `rank_address` | ``$INTEGER`` | No |  |
| `rank_search` | ``$INTEGER`` | No |  |
| `type` | ``$STRING`` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Debug(nil):load({ id = "debug_id" }, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DebugEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReverseEntity

```lua
local reverse = client:Reverse(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$OBJECT`` | No |  |
| `boundingbox` | ``$ARRAY`` | No |  |
| `display_name` | ``$STRING`` | No |  |
| `lat` | ``$STRING`` | No |  |
| `licence` | ``$STRING`` | No |  |
| `lon` | ``$STRING`` | No |  |
| `osm_id` | ``$INTEGER`` | No |  |
| `osm_type` | ``$STRING`` | No |  |
| `place_id` | ``$INTEGER`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Reverse(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | ``$OBJECT`` | No |  |
| `boundingbox` | ``$ARRAY`` | No |  |
| `class` | ``$STRING`` | No |  |
| `display_name` | ``$STRING`` | No |  |
| `icon` | ``$STRING`` | No |  |
| `importance` | ``$NUMBER`` | No |  |
| `lat` | ``$STRING`` | No |  |
| `licence` | ``$STRING`` | No |  |
| `lon` | ``$STRING`` | No |  |
| `osm_id` | ``$INTEGER`` | No |  |
| `osm_type` | ``$STRING`` | No |  |
| `place_id` | ``$INTEGER`` | No |  |
| `type` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search(nil):list(nil, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ServerStatusEntity

```lua
local server_status = client:ServerStatus(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | ``$STRING`` | No |  |
| `database_version` | ``$STRING`` | No |  |
| `message` | ``$STRING`` | No |  |
| `software_version` | ``$STRING`` | No |  |
| `status` | ``$INTEGER`` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ServerStatus(nil):load({ id = "server_status_id" }, nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

