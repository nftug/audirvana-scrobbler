import { Dialog, DialogProps } from '@mui/material'

interface CustomDialogProps extends Omit<DialogProps, 'onClose'> {
  backdropClick?: boolean
  escapeKeyDown?: boolean
  onCloseDialog: () => void
}

const CustomDialog = ({
  backdropClick = false,
  escapeKeyDown = true,
  onCloseDialog,
  ...restProps
}: CustomDialogProps) => {
  const handleClose = (_?: object, reason?: 'backdropClick' | 'escapeKeyDown') => {
    if (!backdropClick && reason === 'backdropClick') return
    if (!escapeKeyDown && reason == 'escapeKeyDown') return
    onCloseDialog()
  }

  return <Dialog onClose={handleClose} {...restProps} />
}

export default CustomDialog
