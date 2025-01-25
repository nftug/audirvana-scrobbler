import FormTextField from '@/lib/form/FormTextField'
import { TrackInfo } from '@bindings/app/bindings'
import { Close, Replay, Save } from '@mui/icons-material'
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton
} from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { createCallable } from 'react-call'
import { FormProvider } from 'react-hook-form'
import { useTrackEditForm } from '../hooks/useTrackEditForm'

interface TrackEditModalProps {
  item: TrackInfo | undefined
}

const TrackEditModal = createCallable<TrackEditModalProps>(({ call, item }) => {
  const [open, closeDialog] = [!call.ended, call.end]
  const { form, mutation } = useTrackEditForm({ item, onSuccess: () => call.end() })
  const confirm = useConfirm()

  const handleClose = async (_?: object, reason?: 'backdropClick' | 'escapeKeyDown') => {
    if (reason === 'backdropClick') return
    if (form.formState.isDirty) {
      try {
        await confirm({ title: 'Confirm', description: 'Discard all changes?' })
      } catch {
        return
      }
    }
    closeDialog()
  }

  const onReset = (e: React.FormEvent) => {
    e.preventDefault()
    form.reset()
  }

  return (
    <FormProvider {...form}>
      <Dialog
        open={open}
        onClose={handleClose}
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
      </Dialog>
    </FormProvider>
  )
}, 500)

export default TrackEditModal
