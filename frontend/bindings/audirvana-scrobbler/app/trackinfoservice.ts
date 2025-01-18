// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as bindings$0 from "./bindings/models.js";

export function DeleteTrackInfo(id: string): Promise<bindings$0.ErrorResponse | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4001489541, id) as any;
    return $resultPromise;
}

export function GetTrackInfo(): Promise<[bindings$0.TrackInfo[] | null, bindings$0.ErrorResponse | null]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4142243114) as any;
    return $resultPromise;
}

export function SaveTrackInfo(id: string, form: bindings$0.TrackInfoForm): Promise<bindings$0.ErrorResponse | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1423142159, id, form) as any;
    return $resultPromise;
}

export function ScrobbleAll(): Promise<bindings$0.ErrorResponse | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3577888866) as any;
    return $resultPromise;
}
