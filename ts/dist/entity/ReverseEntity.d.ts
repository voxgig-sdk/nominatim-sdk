import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { Reverse, ReverseListMatch } from '../NominatimTypes';
declare class ReverseEntity extends NominatimEntityBase<Reverse> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: ReverseEntity): ReverseEntity;
    list(this: any, reqmatch?: ReverseListMatch, ctrl?: Control): Promise<ReverseEntity[]>;
}
export { ReverseEntity };
