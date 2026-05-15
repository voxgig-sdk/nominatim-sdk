# Nominatim SDK utility: make_context

from core.context import NominatimContext


def make_context_util(ctxmap, basectx):
    return NominatimContext(ctxmap, basectx)
