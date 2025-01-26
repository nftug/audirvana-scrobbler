import TrackList from '@/features/track/components/TrackList'
import { useDeleteTrackMutation } from '@/features/track/hooks/useDeleteTrackMutation'
import useNowPlaying from '@/features/track/hooks/useNowPlaying'
import { fullViewHeightStyle, overflowEllipsisStyle } from '@/lib/layout/styles'
import { TrackInfoResponse } from '@bindings/app/bindings'
import {
  Box,
  Card,
  CardContent,
  DialogContentText,
  List,
  ListItem,
  Typography
} from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'

const IndexPage = () => {
  const confirm = useConfirm()
  const deleteTrack = useDeleteTrackMutation()

  const { nowPlaying, error } = useNowPlaying()
  useEffect(() => {
    console.error(error)
  }, [error])

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
    <Box sx={fullViewHeightStyle}>
      <Card sx={{ width: 1, overflowX: 'hidden', mb: 5 }}>
        <CardContent>
          {nowPlaying ? (
            <>
              <Typography gutterBottom sx={{ color: 'text.secondary', fontSize: 14 }}>
                Now playing
              </Typography>

              <Typography variant="h5" component="div" sx={overflowEllipsisStyle}>
                {nowPlaying.track}
              </Typography>
              <Typography sx={{ color: 'text.secondary' }}>
                {nowPlaying.artist ?? 'No artist'}
              </Typography>
            </>
          ) : (
            <Typography gutterBottom sx={{ color: 'text.secondary' }}>
              No track playing
            </Typography>
          )}
        </CardContent>
      </Card>

      <TrackList sx={{ height: 0.8 }} onClickDelete={onClickDelete} />
    </Box>
  )
}

export default IndexPage
