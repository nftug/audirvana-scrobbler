import useErrorHandler from '@/lib/api/useErrorHandler'
import { ErrorCode, ErrorResponse, TrackInfoForm, TrackInfoResponse } from '@bindings/app/bindings'
import { SaveTrackInfo } from '@bindings/app/trackinfoservice'
import { yupResolver } from '@hookform/resolvers/yup'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useEffect } from 'react'
import { FieldPath, useForm } from 'react-hook-form'
import { TrackFieldSchemaType, trackInfoFieldSchema } from '../api/trackInfoFieldSchema'

interface UseTrackInfoFormOptions {
  item: TrackInfoResponse | undefined
  onSuccess: () => void
  dialogOpened?: boolean
}

export const useTrackEditForm = ({ item, onSuccess, dialogOpened }: UseTrackInfoFormOptions) => {
  const queryClient = useQueryClient()
  const handleError = useErrorHandler()

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
          form.setError(errorItem.field as FieldPath<TrackFieldSchemaType>, {
            message: errorItem.message
          })
        }
      } else {
        handleError(error)
      }
    }
  })

  return { form, mutation }
}
