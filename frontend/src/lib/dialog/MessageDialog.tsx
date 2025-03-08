import {
  Button,
  ButtonProps,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogContentTextProps,
  DialogProps,
  DialogTitle
} from '@mui/material'
import { createCallable } from 'react-call'
import CustomDialog from './CustomDialog'

interface MessageDialogProps {
  title?: React.ReactNode
  message?: React.ReactNode
  contentProps?: DialogContentTextProps
  confirmText?: React.ReactNode
  cancelText?: React.ReactNode
  confirmButtonProps?: ButtonProps
  cancelButtonProps?: ButtonProps
  buttonType: 'ok' | 'okCancel'
  backdropClick?: boolean
  escapeKeyDown?: boolean
  dialogProps?: Omit<DialogProps, 'open'>
}

const MessageDialog = createCallable<MessageDialogProps, boolean>(
  ({
    call,
    buttonType = 'ok',
    title = buttonType == 'ok' ? 'Message' : 'Confirm',
    message: content,
    contentProps,
    confirmText = 'OK',
    cancelText = 'Cancel',
    confirmButtonProps,
    cancelButtonProps,
    backdropClick = false,
    escapeKeyDown = true,
    dialogProps
  }) => {
    const [open, closeDialog] = [!call.ended, call.end]

    return (
      <CustomDialog
        open={open}
        onCloseDialog={() => closeDialog(false)}
        backdropClick={backdropClick}
        escapeKeyDown={escapeKeyDown}
        fullWidth
        {...dialogProps}
      >
        <DialogTitle>{title}</DialogTitle>
        <DialogContent>
          <DialogContentText component="div" {...contentProps}>
            {content}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          {buttonType === 'okCancel' && (
            <Button onClick={() => closeDialog(false)} {...cancelButtonProps}>
              {cancelText}
            </Button>
          )}
          <Button onClick={() => closeDialog(true)} {...confirmButtonProps}>
            {confirmText}
          </Button>
        </DialogActions>
      </CustomDialog>
    )
  },
  500
)

export default MessageDialog
