import { formatDateTime } from '@/lib/common/datetimeUtils'
import { overflowEllipsisStyle } from '@/lib/layout/styles'
import { TrackInfo } from '@bindings/app/bindings'
import { Delete, Edit } from '@mui/icons-material'
import { Box, Button, Card, CardActions, CardContent, Stack, Typography } from '@mui/material'

type TrackItemProps = {
  item: TrackInfo
  onClickEdit: (itemId: string) => void
  onClickDelete: (itemId: string) => void
}

const TrackItem = ({ item, onClickEdit, onClickDelete }: TrackItemProps) => {
  return (
    <Card sx={{ marginBottom: '10px', width: 1 }}>
      <CardContent>
        <Box sx={{ display: 'flex', width: 1, alignContent: 'center' }}>
          <Stack sx={{ flexGrow: 1 }}>
            <Typography variant="body2" sx={overflowEllipsisStyle}>
              {item.track}
            </Typography>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {item.artist ?? 'No artist'}
            </Typography>
          </Stack>

          <Stack sx={{ width: 150 }}>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {formatDateTime(item.playedAt)}
            </Typography>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
              {item.album ?? 'No album'}
            </Typography>
          </Stack>
        </Box>
      </CardContent>

      <CardActions>
        <Button
          size="small"
          color="primary"
          onClick={() => onClickEdit(item.id)}
          startIcon={<Edit />}
        >
          Edit
        </Button>
        <Button
          size="small"
          color="error"
          onClick={() => onClickDelete(item.id)}
          startIcon={<Delete />}
        >
          Delete
        </Button>
      </CardActions>
    </Card>
  )
}

export default TrackItem
