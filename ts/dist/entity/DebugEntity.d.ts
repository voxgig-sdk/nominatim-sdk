import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { Debug, DebugLoadMatch } from '../NominatimTypes';
declare class DebugEntity extends NominatimEntityBase<Debug> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: DebugEntity): DebugEntity;
    load(this: any, reqmatch?: DebugLoadMatch, ctrl?: Control): Promise<DebugEntity>;
}
export { DebugEntity };
