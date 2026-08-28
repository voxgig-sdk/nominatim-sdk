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
    public ?string $accept_language = null;
    public ?int $addressdetail = null;
    public ?int $extratag = null;
    public ?string $format = null;
    public ?int $namedetail = null;
    public string $osm_id;
    public ?int $polygon_geojson = null;
    public ?int $polygon_kml = null;
    public ?int $polygon_svg = null;
    public ?int $polygon_text = null;
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
    public ?int $day = null;
    public ?string $format = null;
}

/** Debug entity data model. */
class Debug
{
    public ?array $addresstags = null;
    public ?int $admin_level = null;
    public ?float $calculated_importance = null;
    public ?string $calculated_postcode = null;
    public ?string $calculated_wikipedia = null;
    public ?string $category = null;
    public ?array $centroid = null;
    public ?string $country_code = null;
    public ?array $extratags = null;
    public ?array $geometry = null;
    public ?string $housenumber = null;
    public ?float $importance = null;
    public ?string $indexed_date = null;
    public ?bool $isarea = null;
    public ?string $localname = null;
    public ?array $names = null;
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
    public ?int $addressdetail = null;
    public ?string $class = null;
    public ?string $format = null;
    public ?int $group_hierarchy = null;
    public ?int $keyword = null;
    public ?int $osmid = null;
    public ?string $osmtype = null;
    public ?int $place_id = null;
    public ?int $polygon_geojson = null;
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
    public ?string $accept_language = null;
    public ?int $addressdetail = null;
    public ?int $extratag = null;
    public ?string $format = null;
    public float $lat;
    public float $lon;
    public ?int $namedetail = null;
    public ?int $polygon_geojson = null;
    public ?int $polygon_kml = null;
    public ?int $polygon_svg = null;
    public ?int $polygon_text = null;
    public ?int $zoom = null;
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
    public ?string $accept_language = null;
    public ?int $addressdetail = null;
    public ?int $bounded = null;
    public ?string $city = null;
    public ?string $country = null;
    public ?string $countrycode = null;
    public ?string $county = null;
    public ?int $dedupe = null;
    public ?int $extratag = null;
    public ?string $format = null;
    public ?int $limit = null;
    public ?int $namedetail = null;
    public ?int $polygon_geojson = null;
    public ?int $polygon_kml = null;
    public ?int $polygon_svg = null;
    public ?int $polygon_text = null;
    public ?string $postalcode = null;
    public ?string $q = null;
    public ?string $state = null;
    public ?string $street = null;
    public ?string $viewbox = null;
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
    public ?string $format = null;
}

