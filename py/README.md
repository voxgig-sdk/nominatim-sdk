# Nominatim Python SDK



The Python SDK for the Nominatim API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.AddressLookup()` — each
carrying a small, uniform set of operations (`list`, `load`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK()
```

### 2. List addresslookup records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    addresslookups = client.AddressLookup().list({"osm_id": "example"})
    for addresslookup in addresslookups:
        print(addresslookup)
except Exception as err:
    print(f"list failed: {err}")
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    reverses = client.Reverse().list()
    print(reverses)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = NominatimSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
reverse = client.Reverse().list()
# reverse contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = NominatimSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### NominatimSDK

```python
from nominatim_sdk import NominatimSDK

client = NominatimSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = NominatimSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `AddressLookup` | `(data) -> AddressLookupEntity` | Create an AddressLookup entity instance. |
| `Administrative` | `(data) -> AdministrativeEntity` | Create an Administrative entity instance. |
| `Debug` | `(data) -> DebugEntity` | Create a Debug entity instance. |
| `Reverse` | `(data) -> ReverseEntity` | Create a Reverse entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |
| `ServerStatus` | `(data) -> ServerStatusEntity` | Create a ServerStatus entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### AddressLookup

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `class` | Main OSM tag key |
| `display_name` | Full comma-separated address |
| `importance` | Computed importance rank |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |

Operations: List.

API path: `/lookup`

#### Administrative

| Field | Description |
| --- | --- |
| `class` | Main OSM tag key |
| `country_code` | Country code |
| `errormessage` | Error message describing the polygon issue |
| `name` | Name of the object |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |
| `updated` | Last update timestamp |

Operations: List.

API path: `/polygons`

#### Debug

| Field | Description |
| --- | --- |
| `addresstags` | Address tags |
| `admin_level` | Administrative level |
| `calculated_importance` | Calculated importance |
| `calculated_postcode` | Calculated postcode |
| `calculated_wikipedia` | Wikipedia reference |
| `category` | Main OSM tag key |
| `centroid` | Centroid coordinates |
| `country_code` | Country code |
| `extratags` | Extra OSM tags |
| `geometry` | Geometry information |
| `housenumber` | House number |
| `importance` | Computed importance rank |
| `indexed_date` | Date when the object was indexed |
| `isarea` | Whether the object is an area |
| `localname` | Local name |
| `names` | All available names |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `parent_place_id` | Parent place ID |
| `place_id` | Unique identifier for the place |
| `rank_address` | Address rank |
| `rank_search` | Search rank |
| `type` | Main OSM tag value |

Operations: Load.

API path: `/details`

#### Reverse

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `display_name` | Full comma-separated address |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |

Operations: List.

API path: `/reverse`

#### Search

| Field | Description |
| --- | --- |
| `address` | Address breakdown |
| `boundingbox` | Bounding box coordinates |
| `class` | Main OSM tag key |
| `display_name` | Full comma-separated address |
| `icon` | URL of icon representing the place |
| `importance` | Computed importance rank |
| `lat` | Latitude |
| `licence` | License information |
| `lon` | Longitude |
| `osm_id` | OSM object ID |
| `osm_type` | OSM type (node, way, relation) |
| `place_id` | Unique identifier for the place |
| `type` | Main OSM tag value |

Operations: List.

API path: `/search`

#### ServerStatus

| Field | Description |
| --- | --- |
| `data_updated` | Timestamp when the database was last updated |
| `database_version` | Database version |
| `message` | Status message |
| `software_version` | Nominatim software version |
| `status` | Status code (0 = OK) |

Operations: Load.

API path: `/status`



## Entities


### AddressLookup

Create an instance: `address_lookup = client.AddressLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `dict` | Address breakdown |
| `boundingbox` | `list` | Bounding box coordinates |
| `class` | `str` | Main OSM tag key |
| `display_name` | `str` | Full comma-separated address |
| `importance` | `float` | Computed importance rank |
| `lat` | `str` | Latitude |
| `licence` | `str` | License information |
| `lon` | `str` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `str` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `str` | Main OSM tag value |

#### Example: List

```python
address_lookups = client.AddressLookup().list({"osm_id": "example"})
```


### Administrative

Create an instance: `administrative = client.Administrative()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | `str` | Main OSM tag key |
| `country_code` | `str` | Country code |
| `errormessage` | `str` | Error message describing the polygon issue |
| `name` | `str` | Name of the object |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `str` | OSM type (way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `str` | Main OSM tag value |
| `updated` | `str` | Last update timestamp |

#### Example: List

```python
administratives = client.Administrative().list()
```


### Debug

Create an instance: `debug = client.Debug()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstags` | `dict` | Address tags |
| `admin_level` | `int` | Administrative level |
| `calculated_importance` | `float` | Calculated importance |
| `calculated_postcode` | `str` | Calculated postcode |
| `calculated_wikipedia` | `str` | Wikipedia reference |
| `category` | `str` | Main OSM tag key |
| `centroid` | `dict` | Centroid coordinates |
| `country_code` | `str` | Country code |
| `extratags` | `dict` | Extra OSM tags |
| `geometry` | `dict` | Geometry information |
| `housenumber` | `str` | House number |
| `importance` | `float` | Computed importance rank |
| `indexed_date` | `str` | Date when the object was indexed |
| `isarea` | `bool` | Whether the object is an area |
| `localname` | `str` | Local name |
| `names` | `dict` | All available names |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `str` | OSM type (node, way, relation) |
| `parent_place_id` | `int` | Parent place ID |
| `place_id` | `int` | Unique identifier for the place |
| `rank_address` | `int` | Address rank |
| `rank_search` | `int` | Search rank |
| `type` | `str` | Main OSM tag value |

#### Example: Load

```python
debug = client.Debug().load()
```


### Reverse

Create an instance: `reverse = client.Reverse()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `dict` | Address breakdown |
| `boundingbox` | `list` | Bounding box coordinates |
| `display_name` | `str` | Full comma-separated address |
| `lat` | `str` | Latitude |
| `licence` | `str` | License information |
| `lon` | `str` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `str` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |

#### Example: List

```python
reverses = client.Reverse().list({"lat": 1, "lon": 1})
```


### Search

Create an instance: `search = client.Search()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `dict` | Address breakdown |
| `boundingbox` | `list` | Bounding box coordinates |
| `class` | `str` | Main OSM tag key |
| `display_name` | `str` | Full comma-separated address |
| `icon` | `str` | URL of icon representing the place |
| `importance` | `float` | Computed importance rank |
| `lat` | `str` | Latitude |
| `licence` | `str` | License information |
| `lon` | `str` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `str` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `str` | Main OSM tag value |

#### Example: List

```python
searchs = client.Search().list()
```


### ServerStatus

Create an instance: `server_status = client.ServerStatus()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | `str` | Timestamp when the database was last updated |
| `database_version` | `str` | Database version |
| `message` | `str` | Status message |
| `software_version` | `str` | Nominatim software version |
| `status` | `int` | Status code (0 = OK) |

#### Example: Load

```python
server_status = client.ServerStatus().load()
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── nominatim_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`nominatim_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
reverse = client.Reverse()
reverse.list()

# reverse.data_get() now returns the reverse data from the last list
# reverse.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
