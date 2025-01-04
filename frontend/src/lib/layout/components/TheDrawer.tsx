import { DrawerContext, DrawerDispatchContext } from '@/lib/layout/components/HeaderContext'
import { Home, Info, Settings } from '@mui/icons-material'
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText
} from '@mui/material'
import { useContext } from 'react'
import { Link } from 'react-router-dom'

type DrawerItem = {
  name: string
  href: string
  icon?: React.JSX.Element
}

const TheDrawer = () => {
  const [drawerOpened, setDrawerOpened] = [
    useContext(DrawerContext),
    useContext(DrawerDispatchContext)
  ]

  const menuItems: DrawerItem[] = [
    { name: 'Home', href: '/', icon: <Home /> },
    { name: 'Settings', href: '/settings', icon: <Settings /> },
    { name: 'About', href: '/about', icon: <Info /> }
  ]

  return (
    <Drawer open={drawerOpened} onClose={() => setDrawerOpened(false)}>
      <Box sx={{ width: 250 }} role="presentation" onClick={() => setDrawerOpened(false)}>
        <List>
          {menuItems.map((item) => (
            <ListItem key={item.name} disablePadding>
              <ListItemButton component={Link} to={item.href}>
                <ListItemIcon children={item.icon} />
                <ListItemText primary={item.name} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Box>
    </Drawer>
  )
}

export default TheDrawer
