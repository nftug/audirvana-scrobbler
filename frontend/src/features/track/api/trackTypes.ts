export type NowPlayingResponse = {
  appName: string
  track: string
  artist: string
  album: string
  duration: number
  position: number
  isAdded: boolean
}

export enum AppEvent {
  NotifyNowPlaying = 'NotifyNowPlaying',
  NotifyAdded = 'NotifyAdded'
}
