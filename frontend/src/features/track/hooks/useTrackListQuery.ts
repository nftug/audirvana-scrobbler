import { ErrorResponse, TrackInfoResponse } from '@bindings/app/bindings'
import { GetTrackInfoList } from '@bindings/app/trackinfoservice'
import { useQuery } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'
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

  const confirm = useConfirm()

  useEffect(() => {
    if (!error) return
    confirm({
      title: 'Error',
      description: error.data?.at(0)?.message,
      hideCancelButton: true
    })
  }, [error, confirm])

  return { data, error, isPending }
}
