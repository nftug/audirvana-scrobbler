import { Box, CircularProgress, Divider, Stack, Typography } from '@mui/material'
import AutoSizer from 'react-virtualized-auto-sizer'
import { FixedSizeList } from 'react-window'
import { useTrackListQuery } from '../hooks/useTrackListQuery'
import TrackItem from './TrackItem'

const TrackList = () => {
  const { data, isPending } = useTrackListQuery()

  return (
    <Box sx={{ height: 1 }}>
      {!data?.length || isPending || !data ? (
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
        <AutoSizer>
          {({ width, height }) => (
            <FixedSizeList width={width} height={height} itemSize={100} itemCount={data.length}>
              {({ index, style }) => (
                <Stack key={index} style={style} spacing={1}>
                  {<TrackItem track={data[index]} />}
                  {data && index < data?.length - 1 && <Divider />}
                </Stack>
              )}
            </FixedSizeList>
          )}
        </AutoSizer>
      )}
    </Box>
  )
}

export default TrackList
