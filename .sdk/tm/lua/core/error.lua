-- Nominatim SDK error

local NominatimError = {}
NominatimError.__index = NominatimError


function NominatimError.new(code, msg, ctx)
  local self = setmetatable({}, NominatimError)
  self.is_sdk_error = true
  self.sdk = "Nominatim"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function NominatimError:error()
  return self.msg
end


function NominatimError:__tostring()
  return self.msg
end


return NominatimError
