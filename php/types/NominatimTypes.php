<?php
declare(strict_types=1);

// Typed models for the Nominatim SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** AddressLookup entity data model. */
class AddressLookup
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $class = null;
    public ?string $display_name = null;
    public ?float $importance = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
}

/** Request payload for AddressLookup#list. */
class AddressLookupListMatch
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $class = null;
    public ?string $display_name = null;
    public ?float $importance = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
}

/** Administrative entity data model. */
class Administrative
{
    public ?string $class = null;
    public ?string $country_code = null;
    public ?string $errormessage = null;
    public ?string $name = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
    public ?string $updated = null;
}

/** Request payload for Administrative#list. */
class AdministrativeListMatch
{
    public ?string $class = null;
    public ?string $country_code = null;
    public ?string $errormessage = null;
    public ?string $name = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
    public ?string $updated = null;
}

/** Debug entity data model. */
class Debug
{
    public ?array $addresstag = null;
    public ?int $admin_level = null;
    public ?float $calculated_importance = null;
    public ?string $calculated_postcode = null;
    public ?string $calculated_wikipedia = null;
    public ?string $category = null;
    public ?array $centroid = null;
    public ?string $country_code = null;
    public ?array $extratag = null;
    public ?array $geometry = null;
    public ?string $housenumber = null;
    public ?float $importance = null;
    public ?string $indexed_date = null;
    public ?bool $isarea = null;
    public ?string $localname = null;
    public ?array $name = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $parent_place_id = null;
    public ?int $place_id = null;
    public ?int $rank_address = null;
    public ?int $rank_search = null;
    public ?string $type = null;
}

/** Request payload for Debug#load. */
class DebugLoadMatch
{
    public ?array $addresstag = null;
    public ?int $admin_level = null;
    public ?float $calculated_importance = null;
    public ?string $calculated_postcode = null;
    public ?string $calculated_wikipedia = null;
    public ?string $category = null;
    public ?array $centroid = null;
    public ?string $country_code = null;
    public ?array $extratag = null;
    public ?array $geometry = null;
    public ?string $housenumber = null;
    public ?float $importance = null;
    public ?string $indexed_date = null;
    public ?bool $isarea = null;
    public ?string $localname = null;
    public ?array $name = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $parent_place_id = null;
    public ?int $place_id = null;
    public ?int $rank_address = null;
    public ?int $rank_search = null;
    public ?string $type = null;
}

/** Reverse entity data model. */
class Reverse
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $display_name = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
}

/** Request payload for Reverse#list. */
class ReverseListMatch
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $display_name = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
}

/** Search entity data model. */
class Search
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $class = null;
    public ?string $display_name = null;
    public ?string $icon = null;
    public ?float $importance = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?array $address = null;
    public ?array $boundingbox = null;
    public ?string $class = null;
    public ?string $display_name = null;
    public ?string $icon = null;
    public ?float $importance = null;
    public ?string $lat = null;
    public ?string $licence = null;
    public ?string $lon = null;
    public ?int $osm_id = null;
    public ?string $osm_type = null;
    public ?int $place_id = null;
    public ?string $type = null;
}

/** ServerStatus entity data model. */
class ServerStatus
{
    public ?string $data_updated = null;
    public ?string $database_version = null;
    public ?string $message = null;
    public ?string $software_version = null;
    public ?int $status = null;
}

/** Request payload for ServerStatus#load. */
class ServerStatusLoadMatch
{
    public ?string $data_updated = null;
    public ?string $database_version = null;
    public ?string $message = null;
    public ?string $software_version = null;
    public ?int $status = null;
}

