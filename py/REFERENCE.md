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
address_lookup = client.address_lookup
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.address_lookup.list({})
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
administrative = client.administrative
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.administrative.list({})
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
debug = client.debug
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.debug.load({"id": "debug_id"})
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
reverse = client.reverse
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.reverse.list({})
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
search = client.search
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.search.list({})
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
server_status = client.server_status
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.server_status.load({"id": "server_status_id"})
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

