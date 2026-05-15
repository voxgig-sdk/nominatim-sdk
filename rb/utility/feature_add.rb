# Nominatim SDK utility: feature_add
module NominatimUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
