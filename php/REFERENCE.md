# Nominatim PHP SDK Reference

Complete API reference for the Nominatim PHP SDK.


## NominatimSDK

### Constructor

```php
require_once __DIR__ . '/nominatim_sdk.php';

$client = new NominatimSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = NominatimSDK::test();
```


### Instance Methods

#### `AddressLookup($data = null)`

Create a new `AddressLookupEntity` instance. Pass `null` for no initial data.

#### `Administrative($data = null)`

Create a new `AdministrativeEntity` instance. Pass `null` for no initial data.

#### `Debug($data = null)`

Create a new `DebugEntity` instance. Pass `null` for no initial data.

#### `Reverse($data = null)`

Create a new `ReverseEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `ServerStatus($data = null)`

Create a new `ServerStatusEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): NominatimUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AddressLookupEntity

```php
$address_lookup = $client->AddressLookup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `array` | No |  |
| `boundingbox` | `array` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `importance` | `float` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->AddressLookup()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AddressLookupEntity`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## AdministrativeEntity

```php
$administrative = $client->Administrative();
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Administrative()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AdministrativeEntity`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DebugEntity

```php
$debug = $client->Debug();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresstag` | `array` | No |  |
| `admin_level` | `int` | No |  |
| `calculated_importance` | `float` | No |  |
| `calculated_postcode` | `string` | No |  |
| `calculated_wikipedia` | `string` | No |  |
| `category` | `string` | No |  |
| `centroid` | `array` | No |  |
| `country_code` | `string` | No |  |
| `extratag` | `array` | No |  |
| `geometry` | `array` | No |  |
| `housenumber` | `string` | No |  |
| `importance` | `float` | No |  |
| `indexed_date` | `string` | No |  |
| `isarea` | `bool` | No |  |
| `localname` | `string` | No |  |
| `name` | `array` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `parent_place_id` | `int` | No |  |
| `place_id` | `int` | No |  |
| `rank_address` | `int` | No |  |
| `rank_search` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Debug()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DebugEntity`

Create a new `DebugEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReverseEntity

```php
$reverse = $client->Reverse();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `array` | No |  |
| `boundingbox` | `array` | No |  |
| `display_name` | `string` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Reverse()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReverseEntity`

Create a new `ReverseEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `array` | No |  |
| `boundingbox` | `array` | No |  |
| `class` | `string` | No |  |
| `display_name` | `string` | No |  |
| `icon` | `string` | No |  |
| `importance` | `float` | No |  |
| `lat` | `string` | No |  |
| `licence` | `string` | No |  |
| `lon` | `string` | No |  |
| `osm_id` | `int` | No |  |
| `osm_type` | `string` | No |  |
| `place_id` | `int` | No |  |
| `type` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ServerStatusEntity

```php
$server_status = $client->ServerStatus();
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ServerStatus()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ServerStatusEntity`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new NominatimSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

