import { formatDateTime } from '@/lib/common/dayjsUtils'
import IconButtonMenu from '@/lib/menu/IconButtonMenu'
import { TrackInfoResponse } from '@bindings/app/bindings'
import { Audiotrack, Delete, Edit, MoreVert } from '@mui/icons-material'
import { ListItem, ListItemIcon, ListItemText, MenuItem, Stack, Typography } from '@mui/material'
import { useDeleteTrackDialog } from '../hooks/useDeleteTrack'
import TrackEditDialog from './TrackEditDialog'

type TrackItemProps = {
  track: TrackInfoResponse
}

const TrackItem = ({ track }: TrackItemProps) => {
  const showDeleteTrackDialog = useDeleteTrackDialog()
  const scrobbled = !!track.scrobbledAt

  return (
    <ListItem
      component="div"
      disablePadding
      sx={{ color: scrobbled ? 'text.disabled' : 'text.primary' }}
      secondaryAction={
        <IconButtonMenu icon={<MoreVert />} buttonProps={{ edge: 'end' }}>
          <MenuItem onClick={() => TrackEditDialog.call({ item: track })} disabled={scrobbled}>
            <ListItemIcon>
              <Edit />
            </ListItemIcon>
            <ListItemText>Edit</ListItemText>
          </MenuItem>
          <MenuItem onClick={() => showDeleteTrackDialog(track)}>
            <ListItemIcon>
              <Delete />
            </ListItemIcon>
            <ListItemText>Delete</ListItemText>
          </MenuItem>
        </IconButtonMenu>
      }
    >
      <ListItemIcon>
        <Audiotrack color={scrobbled ? 'disabled' : 'inherit'} />
      </ListItemIcon>
      <ListItemText
        primary={track.track}
        secondary={
          <Stack spacing={0.2} sx={{ color: scrobbled ? 'text.disabled' : 'text.body2' }}>
            <Typography variant="body2">{`${track.artist ?? 'No artits'} ― ${track.album ?? 'No album'}`}</Typography>
            <Typography variant="body2">{formatDateTime(track.playedAt)}</Typography>
          </Stack>
        }
        slotProps={{ secondary: { component: 'div', mt: 1 } }}
      />
    </ListItem>
  )
}

export default TrackItem
