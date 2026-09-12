import { AddressLookupEntity } from './entity/AddressLookupEntity';
import { AdministrativeEntity } from './entity/AdministrativeEntity';
import { DebugEntity } from './entity/DebugEntity';
import { ReverseEntity } from './entity/ReverseEntity';
import { SearchEntity } from './entity/SearchEntity';
import { ServerStatusEntity } from './entity/ServerStatusEntity';
export type * from './NominatimTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { NominatimEntityBase } from './NominatimEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class NominatimSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    AddressLookup(entopts?: Record<string, any>): AddressLookupEntity;
    Administrative(entopts?: Record<string, any>): AdministrativeEntity;
    Debug(entopts?: Record<string, any>): DebugEntity;
    Reverse(entopts?: Record<string, any>): ReverseEntity;
    Search(entopts?: Record<string, any>): SearchEntity;
    ServerStatus(entopts?: Record<string, any>): ServerStatusEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): NominatimSDK;
    tester(testopts?: any, sdkopts?: any): NominatimSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof NominatimSDK;
export { stdutil, config, BaseFeature, NominatimEntityBase, NominatimSDK, SDK, };
