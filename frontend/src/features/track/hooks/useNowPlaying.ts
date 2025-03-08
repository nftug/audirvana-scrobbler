import { useQueryClient } from '@tanstack/react-query'
import { Events } from '@wailsio/runtime'
import { useEffect, useState } from 'react'
import { ErrorResponse } from 'react-router-dom'
import { AppEvent, NowPlayingResponse } from '../api/trackTypes'

type NowPlayingEventData = [NowPlayingResponse | null, ErrorResponse | null]

const useNowPlaying = () => {
  const [nowPlaying, setNowPlaying] = useState<NowPlayingResponse | null>(null)
  const [error, setError] = useState<ErrorResponse | null>(null)
  const queryClient = useQueryClient()

  useEffect(() => {
    const disposeNowPlaying = Events.On(
      AppEvent.NotifyNowPlaying,
      ({ data }: { data: NowPlayingEventData }) => {
        const [nowPlaying, error] = data

        if (error) {
          setError(error)
          return
        }
        setNowPlaying(nowPlaying)
      }
    )

    const disposeAdded = Events.On(AppEvent.NotifyAdded, () => {
      queryClient.refetchQueries({ queryKey: ['trackList'] })
    })

    return () => {
      disposeNowPlaying()
      disposeAdded()
    }
  }, [queryClient])

  return { nowPlaying, error }
}

export default useNowPlaying
