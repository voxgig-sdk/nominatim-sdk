<?php
declare(strict_types=1);

// Nominatim SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class NominatimMakeContext
{
    public static function call(array $ctxmap, ?NominatimContext $basectx): NominatimContext
    {
        return new NominatimContext($ctxmap, $basectx);
    }
}
