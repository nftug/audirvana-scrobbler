import { formatDateTime } from '@/lib/common/dayjsUtils'
import { overflowEllipsisStyle } from '@/lib/layout/styles'
import { TrackInfoResponse } from '@bindings/app/bindings'
import { Delete, Edit } from '@mui/icons-material'
import { Box, Button, Card, CardActions, CardContent, Stack, Typography } from '@mui/material'
import { useDeleteTrackDialog } from '../hooks/useDeleteTrack'
import TrackEditDialog from './TrackEditDialog'

type TrackItemProps = {
  track: TrackInfoResponse
}

const TrackItem = ({ track }: TrackItemProps) => {
  const showDeleteTrackDialog = useDeleteTrackDialog()

  return (
    <Card sx={{ marginBottom: '10px', width: 1 }}>
      <CardContent>
        <Box sx={{ display: 'flex', width: 1, alignContent: 'center' }}>
          <Stack sx={{ flexGrow: 1 }}>
            <Typography variant="body2" sx={overflowEllipsisStyle}>
              {track.track}
            </Typography>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {track.artist ?? 'No artist'}
            </Typography>
          </Stack>

          <Stack sx={{ width: 150 }}>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {formatDateTime(track.playedAt)}
            </Typography>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {track.album ?? 'No album'}
            </Typography>
          </Stack>
        </Box>
      </CardContent>

      <CardActions>
        <Button
          size="small"
          color="primary"
          onClick={() => TrackEditDialog.call({ item: track })}
          startIcon={<Edit />}
          disabled={!!track.scrobbledAt}
        >
          Edit
        </Button>
        <Button
          size="small"
          color="error"
          onClick={() => showDeleteTrackDialog(track)}
          startIcon={<Delete />}
        >
          Delete
        </Button>
      </CardActions>
    </Card>
  )
}

export default TrackItem
