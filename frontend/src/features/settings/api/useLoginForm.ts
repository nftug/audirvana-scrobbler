import useErrorHandler from '@/lib/api/useErrorHandler'
import { ErrorResponse } from '@bindings/app/bindings'
import { GetSessionKey } from '@bindings/app/trackinfoservice'
import { yupResolver } from '@hookform/resolvers/yup'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useSnackbar } from 'notistack'
import { useForm } from 'react-hook-form'
import { loginFieldSchema, LoginFieldSchemaType } from './loginFieldSchema'
import { getLoginQueryKey } from './useLogin'

export const useLoginForm = ({ onSuccess }: { onSuccess: () => void }) => {
  const form = useForm({ resolver: yupResolver(loginFieldSchema) })

  const { enqueueSnackbar } = useSnackbar()
  const queryClient = useQueryClient()
  const handleError = useErrorHandler()

  const mutation = useMutation({
    mutationFn: async (data: LoginFieldSchemaType) => {
      const error = await GetSessionKey(data.username, data.password)
      if (error) throw error
    },
    onSuccess: () => {
      enqueueSnackbar({
        message: 'Login succeeded.',
        variant: 'success'
      })
      queryClient.invalidateQueries({ queryKey: getLoginQueryKey() })
      onSuccess()
    },
    onError: (error: ErrorResponse) => {
      handleError(error)
    }
  })

  return { form, mutation }
}
