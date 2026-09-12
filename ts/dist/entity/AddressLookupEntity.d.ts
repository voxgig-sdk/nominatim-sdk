import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { AddressLookup, AddressLookupListMatch } from '../NominatimTypes';
declare class AddressLookupEntity extends NominatimEntityBase<AddressLookup> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: AddressLookupEntity): AddressLookupEntity;
    list(this: any, reqmatch?: AddressLookupListMatch, ctrl?: Control): Promise<AddressLookupEntity[]>;
}
export { AddressLookupEntity };
