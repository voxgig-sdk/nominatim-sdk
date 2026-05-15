<?php
declare(strict_types=1);

// Nominatim SDK utility: feature_add

class NominatimFeatureAdd
{
    public static function call(NominatimContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
