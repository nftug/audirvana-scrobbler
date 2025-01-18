import { ErrorCode, ErrorResponse, TrackInfo, TrackInfoForm } from '@bindings/app/bindings'
import { SaveTrackInfo } from '@bindings/app/trackinfoservice'
import { yupResolver } from '@hookform/resolvers/yup'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import { trackInfoFieldSchema } from '../trackInfoFieldSchema'

interface UseTrackInfoFormOptions {
  item: TrackInfo | undefined
  onSuccess: () => void
  dialogOpened?: boolean
}

export const useTrackEditForm = ({ item, onSuccess, dialogOpened }: UseTrackInfoFormOptions) => {
  const confirm = useConfirm()
  const queryClient = useQueryClient()

  const form = useForm({ resolver: yupResolver(trackInfoFieldSchema), mode: 'onChange' })

  useEffect(() => {
    form.reset(item)
  }, [form, item, dialogOpened])

  const mutation = useMutation({
    mutationFn: async (data: TrackInfoForm) => {
      if (!item) return
      const error = await SaveTrackInfo(item.id, data)
      if (error) throw error
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['trackList'] })
      onSuccess()
      setTimeout(() => mutation.reset(), 100)
    },
    onError: (error: ErrorResponse) => {
      if (error.code === ErrorCode.ValidationError && error.data) {
        for (const errorItem of error.data) {
          form.setError(errorItem.field as any, { message: errorItem.message })
        }
      } else {
        confirm({
          title: 'Error',
          description: error.data?.at(0)?.message ?? error.code,
          hideCancelButton: true
        })
      }
    }
  })

  return { form, mutation }
}
