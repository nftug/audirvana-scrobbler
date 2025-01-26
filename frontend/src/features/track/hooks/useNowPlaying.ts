import { Events } from '@wailsio/runtime'
import { useEffect, useState } from 'react'
import { ErrorResponse } from 'react-router-dom'

type NowPlayingEventData = [NowPlayingResponse | null, ErrorResponse | null]

const useNowPlaying = () => {
  const [nowPlaying, setNowPlaying] = useState<NowPlayingResponse | null>(null)
  const [error, setError] = useState<ErrorResponse | null>(null)

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

  return { nowPlaying, error }
}

export default useNowPlaying
