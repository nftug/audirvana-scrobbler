import { IconButton, IconButtonProps, Menu, MenuProps } from '@mui/material'
import { useState } from 'react'

interface IconButtonMenuProps {
  icon: React.ReactNode
  children?: React.ReactNode
  buttonProps?: IconButtonProps
  menuProps?: Omit<MenuProps, 'open' | 'anchorEl' | 'onClose'>
}

const IconButtonMenu = ({ icon, children, buttonProps, menuProps }: IconButtonMenuProps) => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>()
  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    event.stopPropagation()
    setAnchorEl(event.currentTarget)
  }
  const handleClose = () => {
    setAnchorEl(null)
  }

  return (
    <>
      <IconButton onClick={handleMenu} {...buttonProps}>
        {icon}
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleClose}
        onClick={handleClose}
        {...menuProps}
      >
        {children}
      </Menu>
    </>
  )
}

export default IconButtonMenu
