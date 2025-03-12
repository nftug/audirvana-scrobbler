import MessageDialog from '@/lib/dialog/MessageDialog'
import { Logout } from '@bindings/app/trackinfoservice'
import { Login as LoginIcon, Logout as LogoutIcon } from '@mui/icons-material'
import { Button } from '@mui/material'
import { useQueryClient } from '@tanstack/react-query'
import { getLoginQueryKey, useLoginState } from '../api/useLogin'
import LoginDialog from './LoginDialog'

const ToggleLoginSettings = () => {
  const { data: loggedIn } = useLoginState()
  const queryClient = useQueryClient()

  const handleClickToggleLogin = async () => {
    if (loggedIn) {
      const ok = await MessageDialog.call({
        message: 'Are you sure you want to logout?',
        buttonType: 'okCancel'
      })
      if (ok) {
        await Logout()
        queryClient.invalidateQueries({ queryKey: getLoginQueryKey() })
      }
    } else {
      LoginDialog.call()
    }
  }

  return (
    <Button
      variant="contained"
      color={loggedIn ? 'error' : 'primary'}
      startIcon={loggedIn ? <LogoutIcon /> : <LoginIcon />}
      onClick={handleClickToggleLogin}
    >
      {loggedIn ? 'Logout' : 'Login'}
    </Button>
  )
}
export default ToggleLoginSettings
