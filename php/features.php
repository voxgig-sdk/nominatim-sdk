<?php
declare(strict_types=1);

// Nominatim SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class NominatimFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new NominatimBaseFeature();
            case "test":
                return new NominatimTestFeature();
            default:
                return new NominatimBaseFeature();
        }
    }
}
