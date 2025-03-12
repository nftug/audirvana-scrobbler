import CustomDialog from '@/lib/dialog/CustomDialog'
import FormTextField from '@/lib/form/FormTextField'
import { Close, Login, Replay } from '@mui/icons-material'
import { Box, Button, DialogActions, DialogContent, DialogTitle, IconButton } from '@mui/material'
import { createCallable } from 'react-call'
import { FormProvider } from 'react-hook-form'
import { useLoginForm } from '../api/useLoginForm'

const LoginDialog = createCallable(({ call }) => {
  const [open, closeDialog] = [!call.ended, call.end]

  const { form, mutation } = useLoginForm({ onSuccess: closeDialog })

  const onReset = (e: React.FormEvent) => {
    e.preventDefault()
    form.reset()
  }

  return (
    <FormProvider {...form}>
      <CustomDialog
        open={open}
        onCloseDialog={closeDialog}
        slotProps={{
          paper: {
            component: 'form',
            onSubmit: form.handleSubmit((data) => mutation.mutate(data)),
            onReset
          }
        }}
      >
        <DialogTitle sx={{ m: 0, p: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
            Login
            <IconButton
              aria-label="close"
              onClick={() => closeDialog()}
              sx={(theme) => ({ color: theme.palette.grey[500] })}
              children={<Close />}
              title="Close"
            />
          </Box>
        </DialogTitle>

        <DialogContent>
          <FormTextField name="username" label="User name" fullWidth margin="normal" />
          <FormTextField
            name="password"
            label="Password"
            fullWidth
            margin="normal"
            type="password"
          />
        </DialogContent>

        <DialogActions>
          <Button type="reset" startIcon={<Replay />} disabled={!form.formState.isDirty}>
            Reset
          </Button>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            startIcon={<Login />}
            disabled={!form.formState.isValid || mutation.isPending}
          >
            Login
          </Button>
        </DialogActions>
      </CustomDialog>
    </FormProvider>
  )
}, 500)

export default LoginDialog
