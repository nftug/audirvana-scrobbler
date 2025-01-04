import { ApiError, handleApiError } from '@/lib/api/errors'
import { useQuery } from '@tanstack/react-query'
import { GetTrackInfo } from '@wailsjs/go/app/App'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'

export const useTrackListQuery = () => {
  const { data, error, isPending } = useQuery({
    queryKey: ['trackList'],
    queryFn: async () => await handleApiError(async () => await GetTrackInfo())
  })

  const confirm = useConfirm()

  useEffect(() => {
    if (!error) return
    const message = error instanceof ApiError ? error.data?.message : error.message
    confirm({ title: 'Error', description: message, hideCancelButton: true })
  }, [error])

  return { data, error, isPending }
}
