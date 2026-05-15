<?php
declare(strict_types=1);

// Nominatim SDK utility: result_body

class NominatimResultBody
{
    public static function call(NominatimContext $ctx): ?NominatimResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
