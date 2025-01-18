import TrackEditModal from '@/features/track/components/TrackEditModal'
import TrackList from '@/features/track/components/TrackList'
import { useDeleteTrackMutation } from '@/features/track/hooks/useDeleteTrackMutation'
import { fullViewHeightStyle, overflowEllipsisStyle } from '@/lib/layout/styles'
import { TrackInfo } from '@bindings/app/bindings'
import { DialogContentText, List, ListItem } from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { useState } from 'react'

const IndexPage = () => {
  const [selectedTrack, setSelectedTrack] = useState<TrackInfo>()
  const [openEdit, setOpenEdit] = useState(false)
  const confirm = useConfirm()
  const deleteTrack = useDeleteTrackMutation()

  const onClickEdit = (item: TrackInfo) => {
    setSelectedTrack(item)
    setOpenEdit(true)
  }

  const onClickDelete = async (item: TrackInfo) => {
    try {
      await confirm({
        title: 'Confirm',
        content: (
          <DialogContentText sx={overflowEllipsisStyle}>
            Delete this track?
            <List>
              <ListItem>Artist: {item.artist}</ListItem>
              <ListItem>Album: {item.album}</ListItem>
              <ListItem>Track: {item.track}</ListItem>
            </List>
          </DialogContentText>
        )
      })
    } catch {
      return
    }

    deleteTrack.mutate(item.id)
  }

  return (
    <>
      <TrackList sx={fullViewHeightStyle} onClickEdit={onClickEdit} onClickDelete={onClickDelete} />
      <TrackEditModal item={selectedTrack} open={openEdit} onClose={() => setOpenEdit(false)} />
    </>
  )
}

export default IndexPage
