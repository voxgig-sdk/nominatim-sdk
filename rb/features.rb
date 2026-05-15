# Nominatim SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module NominatimFeatures
  def self.make_feature(name)
    case name
    when "base"
      NominatimBaseFeature.new
    when "test"
      NominatimTestFeature.new
    else
      NominatimBaseFeature.new
    end
  end
end
