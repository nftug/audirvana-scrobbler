import { handleApiError } from '@/lib/api/errors'
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
    confirm({ title: 'Error', description: error.message, hideCancelButton: true })
  }, [error])

  useEffect(() => {
    if (!data) return
    console.log(data)
  }, [data])

  return { data, error, isPending }
}
