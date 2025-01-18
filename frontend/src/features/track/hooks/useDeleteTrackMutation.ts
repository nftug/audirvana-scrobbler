import { ErrorResponse } from '@bindings/app/bindings/models'
import { DeleteTrackInfo } from '@bindings/app/trackinfoservice'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'

export const useDeleteTrackMutation = () => {
  const confirm = useConfirm()
  const queryClient = useQueryClient()

  const mutation = useMutation({
    mutationFn: async (id: string) => {
      const error = await DeleteTrackInfo(id)
      if (error) throw error
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['trackList'] })
    },
    onError: (error: ErrorResponse) => {
      confirm({
        title: 'Error',
        description: error.data?.at(0)?.message ?? error.code,
        hideCancelButton: true
      })
    }
  })

  return mutation
}
