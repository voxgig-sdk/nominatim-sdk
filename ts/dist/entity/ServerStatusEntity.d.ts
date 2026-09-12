import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { ServerStatus, ServerStatusLoadMatch } from '../NominatimTypes';
declare class ServerStatusEntity extends NominatimEntityBase<ServerStatus> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: ServerStatusEntity): ServerStatusEntity;
    load(this: any, reqmatch?: ServerStatusLoadMatch, ctrl?: Control): Promise<ServerStatusEntity>;
}
export { ServerStatusEntity };
