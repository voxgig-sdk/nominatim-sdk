# Nominatim Python SDK Reference

Complete API reference for the Nominatim Python SDK.


## NominatimSDK

### Constructor

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = NominatimSDK.test()
```


### Instance Methods

#### `AddressLookup(data=None)`

Create a new `AddressLookupEntity` instance. Pass `None` for no initial data.

#### `Administrative(data=None)`

Create a new `AdministrativeEntity` instance. Pass `None` for no initial data.

#### `Debug(data=None)`

Create a new `DebugEntity` instance. Pass `None` for no initial data.

#### `Reverse(data=None)`

Create a new `ReverseEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `ServerStatus(data=None)`

Create a new `ServerStatusEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AddressLookupEntity

```python
address_lookup = client.AddressLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `dict` | No |  |
| `boundingbox` | `list` | No |  |
| `class` | `str` | No |  |
| `display_name` | `str` | No |  |
| `importance` | `float` | No |  |
| `lat` | `str` | No |  |
| `licence` | `str` | No |  |
| `lon` | `str` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `str` | No |  |
| `place_id` | `int` | No |  |
| `type` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.AddressLookup().list()
for address_lookup in results:
    print(address_lookup)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AddressLookupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AdministrativeEntity

```python
administrative = client.Administrative()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `class` | `str` | No |  |
| `country_code` | `str` | No |  |
| `errormessage` | `str` | No |  |
| `name` | `str` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `str` | No |  |
| `place_id` | `int` | No |  |
| `type` | `str` | No |  |
| `updated` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Administrative().list()
for administrative in results:
    print(administrative)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AdministrativeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DebugEntity

```python
debug = client.Debug()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstag` | `dict` | No |  |
| `admin_level` | `int` | No |  |
| `calculated_importance` | `float` | No |  |
| `calculated_postcode` | `str` | No |  |
| `calculated_wikipedia` | `str` | No |  |
| `category` | `str` | No |  |
| `centroid` | `dict` | No |  |
| `country_code` | `str` | No |  |
| `extratag` | `dict` | No |  |
| `geometry` | `dict` | No |  |
| `housenumber` | `str` | No |  |
| `importance` | `float` | No |  |
| `indexed_date` | `str` | No |  |
| `isarea` | `bool` | No |  |
| `localname` | `str` | No |  |
| `name` | `dict` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `str` | No |  |
| `parent_place_id` | `int` | No |  |
| `place_id` | `int` | No |  |
| `rank_address` | `int` | No |  |
| `rank_search` | `int` | No |  |
| `type` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Debug().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DebugEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReverseEntity

```python
reverse = client.Reverse()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `dict` | No |  |
| `boundingbox` | `list` | No |  |
| `display_name` | `str` | No |  |
| `lat` | `str` | No |  |
| `licence` | `str` | No |  |
| `lon` | `str` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `str` | No |  |
| `place_id` | `int` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Reverse().list()
for reverse in results:
    print(reverse)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `dict` | No |  |
| `boundingbox` | `list` | No |  |
| `class` | `str` | No |  |
| `display_name` | `str` | No |  |
| `icon` | `str` | No |  |
| `importance` | `float` | No |  |
| `lat` | `str` | No |  |
| `licence` | `str` | No |  |
| `lon` | `str` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `str` | No |  |
| `place_id` | `int` | No |  |
| `type` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ServerStatusEntity

```python
server_status = client.ServerStatus()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data_updated` | `str` | No |  |
| `database_version` | `str` | No |  |
| `message` | `str` | No |  |
| `software_version` | `str` | No |  |
| `status` | `int` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ServerStatus().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ServerStatusEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = NominatimSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

