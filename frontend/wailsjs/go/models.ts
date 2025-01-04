export namespace enums {
	
	export enum ErrorCode {
	    ValidationError = "ValidationError",
	    NotFound = "NotFound",
	    InternalError = "InternalError",
	}

}

export namespace response {
	
	export interface TrackInfo {
	    id: string;
	    artist: string;
	    album: string;
	    track: string;
	    playedAt: string;
	}

}

