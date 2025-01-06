import { ApiError } from '@/lib/api/errors'
import { GetTrackInfo } from '@bindings/app/app'
import { useQuery } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'

export const useTrackListQuery = () => {
  const { data, error, isPending } = useQuery({
    queryKey: ['trackList'],
    queryFn: async () => {
      const [data, error] = await GetTrackInfo()
      if (error) throw error
      return data
    }
  })

  const confirm = useConfirm()

  useEffect(() => {
    if (!error) return
    const message = error instanceof ApiError ? error.data?.message : error.message
    confirm({ title: 'Error', description: message, hideCancelButton: true })
  }, [error])

  return { data, error, isPending }
}
