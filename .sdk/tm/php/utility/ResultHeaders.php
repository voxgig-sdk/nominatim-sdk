<?php
declare(strict_types=1);

// Nominatim SDK utility: result_headers

class NominatimResultHeaders
{
    public static function call(NominatimContext $ctx): ?NominatimResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
