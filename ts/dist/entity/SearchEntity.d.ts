import { NominatimEntityBase } from '../NominatimEntityBase';
import type { NominatimSDK } from '../NominatimSDK';
import type { Control } from '../types';
import type { Search, SearchListMatch } from '../NominatimTypes';
declare class SearchEntity extends NominatimEntityBase<Search> {
    constructor(client: NominatimSDK, entopts: any);
    make(this: SearchEntity): SearchEntity;
    list(this: any, reqmatch?: SearchListMatch, ctrl?: Control): Promise<SearchEntity[]>;
}
export { SearchEntity };
