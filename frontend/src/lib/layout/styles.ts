import { Theme } from '@emotion/react'
import { SxProps } from '@mui/material'

export const overflowEllipsisStyle: SxProps<Theme> = {
  overflow: 'hidden',
  textOverflow: 'ellipsis',
  whiteSpace: 'nowrap',
  maxWidth: 1
} as const

export const fullViewHeightStyle: SxProps<Theme> = {
  height: 'calc(100vh - 64px)'
} as const
