// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT


export enum ErrorCode {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero = "",

    ValidationError = "ValidationError",
    NotFoundError = "NotFound",
    InternalError = "InternalError",
};

export interface ErrorData {
    "field": string;
    "message": string;
}

export interface ErrorResponse {
    "code": ErrorCode;
    "data": ErrorData[] | null;
}

export interface TrackInfo {
    "id": string;
    "artist": string;
    "album": string;
    "track": string;
    "playedAt": string;
}

export interface TrackInfoForm {
    "artist": string;
    "album": string;
    "track": string;
}
