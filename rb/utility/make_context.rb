# Nominatim SDK utility: make_context
require_relative '../core/context'
module NominatimUtilities
  MakeContext = ->(ctxmap, basectx) {
    NominatimContext.new(ctxmap, basectx)
  }
end
