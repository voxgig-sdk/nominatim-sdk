# Nominatim PHP SDK



The PHP SDK for the Nominatim API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->AddressLookup()` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nominatim-sdk/releases](https://github.com/voxgig-sdk/nominatim-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'nominatim_sdk.php';

$client = new NominatimSDK();
```

### 2. List addresslookup records

```php
try {
    // list() returns an array of AddressLookup records — iterate directly.
    $addresslookups = $client->AddressLookup()->list();
    foreach ($addresslookups as $item) {
        echo $item["address"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $reverses = $client->Reverse()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = NominatimSDK::test();

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$reverse = $client->Reverse()->list();
print_r($reverse);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new NominatimSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOMINATIM_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### NominatimSDK

```php
require_once 'nominatim_sdk.php';
$client = new NominatimSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = NominatimSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### NominatimSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `AddressLookup` | `($data): AddressLookupEntity` | Create an AddressLookup entity instance. |
| `Administrative` | `($data): AdministrativeEntity` | Create an Administrative entity instance. |
| `Debug` | `($data): DebugEntity` | Create a Debug entity instance. |
| `Reverse` | `($data): ReverseEntity` | Create a Reverse entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |
| `ServerStatus` | `($data): ServerStatusEntity` | Create a ServerStatus entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$address_lookup = $client->AddressLookup();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `array` | Address breakdown |
| `boundingbox` | `array` | Bounding box coordinates |
| `class` | `string` | Main OSM tag key |
| `display_name` | `string` | Full comma-separated address |
| `importance` | `float` | Computed importance rank |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |

#### Example: List

```php
// list() returns an array of AddressLookup records (throws on error).
$address_lookups = $client->AddressLookup()->list();
```


### Administrative

Create an instance: `$administrative = $client->Administrative();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `class` | `string` | Main OSM tag key |
| `country_code` | `string` | Country code |
| `errormessage` | `string` | Error message describing the polygon issue |
| `name` | `string` | Name of the object |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `string` | OSM type (way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |
| `updated` | `string` | Last update timestamp |

#### Example: List

```php
// list() returns an array of Administrative records (throws on error).
$administratives = $client->Administrative()->list();
```


### Debug

Create an instance: `$debug = $client->Debug();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresstags` | `array` | Address tags |
| `admin_level` | `int` | Administrative level |
| `calculated_importance` | `float` | Calculated importance |
| `calculated_postcode` | `string` | Calculated postcode |
| `calculated_wikipedia` | `string` | Wikipedia reference |
| `category` | `string` | Main OSM tag key |
| `centroid` | `array` | Centroid coordinates |
| `country_code` | `string` | Country code |
| `extratags` | `array` | Extra OSM tags |
| `geometry` | `array` | Geometry information |
| `housenumber` | `string` | House number |
| `importance` | `float` | Computed importance rank |
| `indexed_date` | `string` | Date when the object was indexed |
| `isarea` | `bool` | Whether the object is an area |
| `localname` | `string` | Local name |
| `names` | `array` | All available names |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `parent_place_id` | `int` | Parent place ID |
| `place_id` | `int` | Unique identifier for the place |
| `rank_address` | `int` | Address rank |
| `rank_search` | `int` | Search rank |
| `type` | `string` | Main OSM tag value |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Debug record (throws on error).
$debug = $client->Debug()->load();
```


### Reverse

Create an instance: `$reverse = $client->Reverse();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `array` | Address breakdown |
| `boundingbox` | `array` | Bounding box coordinates |
| `display_name` | `string` | Full comma-separated address |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |

#### Example: List

```php
// list() returns an array of Reverse records (throws on error).
$reverses = $client->Reverse()->list();
```


### Search

Create an instance: `$search = $client->Search();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `array` | Address breakdown |
| `boundingbox` | `array` | Bounding box coordinates |
| `class` | `string` | Main OSM tag key |
| `display_name` | `string` | Full comma-separated address |
| `icon` | `string` | URL of icon representing the place |
| `importance` | `float` | Computed importance rank |
| `lat` | `string` | Latitude |
| `licence` | `string` | License information |
| `lon` | `string` | Longitude |
| `osm_id` | `int` | OSM object ID |
| `osm_type` | `string` | OSM type (node, way, relation) |
| `place_id` | `int` | Unique identifier for the place |
| `type` | `string` | Main OSM tag value |

#### Example: List

```php
// list() returns an array of Search records (throws on error).
$searchs = $client->Search()->list();
```


### ServerStatus

Create an instance: `$server_status = $client->ServerStatus();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data_updated` | `string` | Timestamp when the database was last updated |
| `database_version` | `string` | Database version |
| `message` | `string` | Status message |
| `software_version` | `string` | Nominatim software version |
| `status` | `int` | Status code (0 = OK) |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ServerStatus record (throws on error).
$server_status = $client->ServerStatus()->load();
```


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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── nominatim_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`nominatim_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$reverse = $client->Reverse();
$reverse->list();

// $reverse->data_get() now returns the reverse data from the last list
// $reverse->match_get() returns the last match criteria
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
