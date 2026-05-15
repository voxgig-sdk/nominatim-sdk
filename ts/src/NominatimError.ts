
import { Context } from './Context'


class NominatimError extends Error {

  isNominatimError = true

  sdk = 'Nominatim'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  NominatimError
}

