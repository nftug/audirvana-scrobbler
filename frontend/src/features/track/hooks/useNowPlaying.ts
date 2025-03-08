import { useQueryClient } from '@tanstack/react-query'
import { Events } from '@wailsio/runtime'
import { useEffect, useState } from 'react'
import { ErrorResponse } from 'react-router-dom'
import { usePreviousDistinct } from 'react-use'
import { NowPlayingResponse } from '../api/trackTypes'

type NowPlayingEventData = [NowPlayingResponse | null, ErrorResponse | null]

const useNowPlaying = () => {
  const [nowPlaying, setNowPlaying] = useState<NowPlayingResponse | null>(null)
  const [error, setError] = useState<ErrorResponse | null>(null)
  const queryClient = useQueryClient()

  const nowPlayingJson = JSON.stringify({
    track: nowPlaying?.track,
    album: nowPlaying?.album,
    artist: nowPlaying?.artist
  })
  const prevNowPlayingJson = usePreviousDistinct(nowPlayingJson)

  useEffect(() => {
    const dispose = Events.On('NotifyNowPlaying', ({ data }: { data: NowPlayingEventData }) => {
      const [nowPlaying, error] = data

      if (error) {
        setError(error)
        return
      }
      setNowPlaying(nowPlaying)
    })

    return () => dispose()
  }, [])

  useEffect(() => {
    if (nowPlayingJson !== prevNowPlayingJson) {
      queryClient.invalidateQueries({ queryKey: ['trackList'] })
    }
  }, [nowPlayingJson, prevNowPlayingJson, queryClient])

  return { nowPlaying, error }
}

export default useNowPlaying
