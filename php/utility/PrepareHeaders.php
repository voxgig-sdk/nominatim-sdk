<?php
declare(strict_types=1);

// Nominatim SDK utility: prepare_headers

class NominatimPrepareHeaders
{
    public static function call(NominatimContext $ctx): array
    {
        $options = $ctx->client->options_map();
        $headers = \Voxgig\Struct\Struct::getprop($options, 'headers');
        if (!$headers) {
            return [];
        }
        $out = \Voxgig\Struct\Struct::clone($headers);
        return is_array($out) ? $out : [];
    }
}
