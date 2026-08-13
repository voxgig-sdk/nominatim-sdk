# Nominatim SDK utility: make_context

from nominatim_sdk.core.context import NominatimContext


def make_context_util(ctxmap, basectx):
    return NominatimContext(ctxmap, basectx)
