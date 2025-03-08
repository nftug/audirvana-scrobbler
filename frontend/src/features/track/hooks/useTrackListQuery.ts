import useErrorHandler from '@/lib/api/useErrorHandler'
import { ErrorResponse, TrackInfoResponse } from '@bindings/app/bindings'
import { GetTrackInfoList } from '@bindings/app/trackinfoservice'
import { useQuery } from '@tanstack/react-query'
import { useEffect } from 'react'

export const useTrackListQuery = () => {
  const { data, error, isPending } = useQuery<TrackInfoResponse[], ErrorResponse>({
    queryKey: ['trackList'],
    queryFn: async () => {
      const [data, error] = await GetTrackInfoList()
      if (error) throw error
      return data!
    }
  })

  const handleError = useErrorHandler()

  useEffect(() => {
    if (!error) return
    handleError(error)
  }, [error, handleError])

  return { data, error, isPending }
}
