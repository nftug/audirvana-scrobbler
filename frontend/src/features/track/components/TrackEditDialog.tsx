import CustomDialog from '@/lib/dialog/CustomDialog'
import MessageDialog from '@/lib/dialog/MessageDialog'
import FormTextField from '@/lib/form/FormTextField'
import { TrackInfoResponse } from '@bindings/app/bindings'
import { Close, Replay, Save } from '@mui/icons-material'
import { Box, Button, DialogActions, DialogContent, DialogTitle, IconButton } from '@mui/material'
import { createCallable } from 'react-call'
import { FormProvider } from 'react-hook-form'
import { useTrackEditForm } from '../hooks/useTrackEditForm'

interface TrackEditDialogProps {
  item: TrackInfoResponse | undefined
}

const TrackEditDialog = createCallable<TrackEditDialogProps>(({ call, item }) => {
  const [open, closeDialog] = [!call.ended, call.end]
  const { form, mutation } = useTrackEditForm({ item, onSuccess: () => closeDialog() })

  const handleClose = async () => {
    if (form.formState.isDirty) {
      const ok = await MessageDialog.call({
        message: 'Discard all changes?',
        buttonType: 'okCancel'
      })
      if (!ok) return
    }
    closeDialog()
  }

  const onReset = (e: React.FormEvent) => {
    e.preventDefault()
    form.reset()
  }

  return (
    <FormProvider {...form}>
      <CustomDialog
        open={open}
        onCloseDialog={handleClose}
        PaperProps={{
          component: 'form',
          onSubmit: form.handleSubmit((data) => mutation.mutate(data)),
          onReset
        }}
      >
        <DialogTitle sx={{ m: 0, p: 2 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
            Edit the track
            <IconButton
              aria-label="close"
              onClick={handleClose}
              sx={(theme) => ({ color: theme.palette.grey[500] })}
              children={<Close />}
              title="Close"
            />
          </Box>
        </DialogTitle>

        <DialogContent dividers>
          <FormTextField name="artist" label="Artist" fullWidth margin="normal" />
          <FormTextField name="album" label="Album" fullWidth margin="normal" />
          <FormTextField name="track" label="Track" fullWidth margin="normal" />
        </DialogContent>

        <DialogActions>
          <Button type="reset" startIcon={<Replay />} disabled={!form.formState.isDirty}>
            Reset
          </Button>
          <Button
            type="submit"
            variant="contained"
            color="primary"
            startIcon={<Save />}
            disabled={!form.formState.isValid || mutation.isPending}
          >
            Save
          </Button>
        </DialogActions>
      </CustomDialog>
    </FormProvider>
  )
}, 500)

export default TrackEditDialog
