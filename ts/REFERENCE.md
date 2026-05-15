# Nominatim TypeScript SDK Reference

Complete API reference for the Nominatim TypeScript SDK.


## NominatimSDK

### Constructor

```ts
new NominatimSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NominatimSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NominatimSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NominatimSDK` instance in test mode.


### Instance Methods

#### `AddressLookup(data?: object)`

Create a new `AddressLookup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AddressLookupEntity` instance.

#### `Administrative(data?: object)`

Create a new `Administrative` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AdministrativeEntity` instance.

#### `Debug(data?: object)`

Create a new `Debug` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DebugEntity` instance.

#### `Reverse(data?: object)`

Create a new `Reverse` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReverseEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `ServerStatus(data?: object)`

Create a new `ServerStatus` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ServerStatusEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NominatimSDK.test()`.

**Returns:** `NominatimSDK` instance in test mode.


---

## AddressLookupEntity

```ts
const address_lookup = client.AddressLookup()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.AddressLookup().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AddressLookupEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AdministrativeEntity

```ts
const administrative = client.Administrative()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Administrative().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AdministrativeEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DebugEntity

```ts
const debug = client.Debug()
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Debug().load({ id: 'debug_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DebugEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReverseEntity

```ts
const reverse = client.Reverse()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Reverse().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReverseEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ServerStatusEntity

```ts
const server_status = client.ServerStatus()
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ServerStatus().load({ id: 'server_status_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ServerStatusEntity` instance with the same client and
options.

#### `client()`

Return the parent `NominatimSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NominatimSDK({
  feature: {
    test: { active: true },
  }
})
```

