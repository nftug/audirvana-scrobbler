import TrackList from '@/features/track/components/TrackList'
import { useDeleteTrackMutation } from '@/features/track/hooks/useDeleteTrackMutation'
import { fullViewHeightStyle, overflowEllipsisStyle } from '@/lib/layout/styles'
import { TrackInfoResponse } from '@bindings/app/bindings'
import { DialogContentText, List, ListItem } from '@mui/material'
import { useConfirm } from 'material-ui-confirm'

const IndexPage = () => {
  const confirm = useConfirm()
  const deleteTrack = useDeleteTrackMutation()

  const onClickDelete = async (item: TrackInfoResponse) => {
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
      <TrackList sx={fullViewHeightStyle} onClickDelete={onClickDelete} />
    </>
  )
}

export default IndexPage
