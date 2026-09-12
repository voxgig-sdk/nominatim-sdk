import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { Administrative, AdministrativeListMatch } from '../NominatimTypes';
declare class AdministrativeEntity extends NominatimEntityBase<Administrative> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: AdministrativeEntity): AdministrativeEntity;
    list(this: any, reqmatch?: AdministrativeListMatch, ctrl?: Control): Promise<AdministrativeEntity[]>;
}
export { AdministrativeEntity };
