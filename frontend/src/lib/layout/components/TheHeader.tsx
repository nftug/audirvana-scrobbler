import { DrawerDispatchContext } from '@/lib/layout/components/HeaderContext'
import { ScrobbleAll } from '@bindings/app/trackinfoservice'
import { Send } from '@mui/icons-material'
import MenuIcon from '@mui/icons-material/Menu'
import { AppBar, IconButton, Toolbar, Typography } from '@mui/material'
import { useQueryClient } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'
import { useContext } from 'react'

const TheHeader = () => {
  const setDrawerOpened = useContext(DrawerDispatchContext)
  const queryClient = useQueryClient()
  const confirm = useConfirm()

  const toggleDrawer = () => {
    setDrawerOpened((x) => !x)
  }

  const scrobbleAll = async () => {
    const error = await ScrobbleAll()
    if (error) {
      confirm({
        title: 'Error',
        description: error.data?.at(0)?.message ?? error.code,
        hideCancelButton: true
      })
    }

    queryClient.invalidateQueries({ queryKey: ['trackList'] })
  }

  return (
    <AppBar>
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2 }}
          onClick={toggleDrawer}
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h5" sx={{ flexGrow: 1 }}>
          Audirvana Scrobbler
        </Typography>

        <IconButton onClick={scrobbleAll}>
          <Send />
        </IconButton>
      </Toolbar>
    </AppBar>
  )
}

export default TheHeader
