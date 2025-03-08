import TrackEditDialog from '@/features/track/components/TrackEditDialog'
import MessageDialog from '@/lib/dialog/MessageDialog'

const DialogProvider = () => {
  return (
    <>
      <MessageDialog.Root />
      <TrackEditDialog.Root />
    </>
  )
}

export default DialogProvider
