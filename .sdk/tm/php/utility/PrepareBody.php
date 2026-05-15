<?php
declare(strict_types=1);

// Nominatim SDK utility: prepare_body

class NominatimPrepareBody
{
    public static function call(NominatimContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
