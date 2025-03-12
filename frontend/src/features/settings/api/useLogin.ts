import { Login } from '@bindings/app/trackinfoservice'
import { useQuery } from '@tanstack/react-query'
import { Window } from '@wailsio/runtime'
import { useSnackbar } from 'notistack'
import { useEffect } from 'react'

export const getLoginQueryKey = () => ['login'] as const

export const useLoginState = () =>
  useQuery({
    queryFn: async () => {
      const loggedIn = await Login()
      return loggedIn
    },
    queryKey: getLoginQueryKey(),
    initialData: undefined
  })

export const useLoginRoot = () => {
  const { enqueueSnackbar } = useSnackbar()
  const { data: loggedIn } = useLoginState()

  useEffect(() => {
    if (loggedIn !== false) return
    enqueueSnackbar({
      message: 'You are not logged in Last.fm. Login on the setting page.',
      variant: 'warning'
    })
    Window.Show()
  }, [loggedIn, enqueueSnackbar])
}
