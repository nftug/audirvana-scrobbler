import { Theme } from '@emotion/react'
import { Box, CircularProgress, List, ListItem, Stack, SxProps, Typography } from '@mui/material'
import { useTrackListQuery } from '../hooks/useTrackListQuery'
import TrackItem from './TrackItem'

type TrackListProps = {
  sx?: SxProps<Theme>
}

const TrackList = ({ sx }: TrackListProps) => {
  const { data, isPending } = useTrackListQuery()

  return (
    <List sx={sx}>
      {!data?.length || isPending ? (
        <Box display="flex" justifyContent="center" alignItems="center" height={1}>
          {isPending ? (
            <Stack display="flex" justifyContent="center" alignItems="center" spacing={3}>
              <CircularProgress size={80} />
              <Typography variant="h6" color="textSecondary">
                Loading track list...
              </Typography>
            </Stack>
          ) : (
            <Typography variant="h6" color="textSecondary">
              No tracks to scrobble.
            </Typography>
          )}
        </Box>
      ) : (
        data.map((item) => (
          <ListItem key={item.id}>
            <TrackItem track={item} />
          </ListItem>
        ))
      )}
    </List>
  )
}

export default TrackList
