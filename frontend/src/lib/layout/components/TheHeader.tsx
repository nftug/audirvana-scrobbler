import { DrawerDispatchContext } from '@/lib/layout/components/HeaderContext'
import MenuIcon from '@mui/icons-material/Menu'
import { AppBar, IconButton, Stack, Toolbar, Typography } from '@mui/material'
import { useContext } from 'react'
import { useActionElementRef } from '../atoms/portalAtom'

const TheHeader = () => {
  const setDrawerOpened = useContext(DrawerDispatchContext)
  const actionElementRef = useActionElementRef()

  const toggleDrawer = () => {
    setDrawerOpened((x) => !x)
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

        <Stack ref={actionElementRef} direction="row" useFlexGap />
      </Toolbar>
    </AppBar>
  )
}

export default TheHeader
