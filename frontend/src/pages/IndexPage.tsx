import TrackList from '@/features/track/components/TrackList'
import TrackListAppBar from '@/features/track/components/TrackListAppBar'
import useNowPlaying from '@/features/track/hooks/useNowPlaying'
import { fullViewHeightStyle, overflowEllipsisStyle } from '@/lib/layout/styles'
import { Box, Card, CardContent, Typography } from '@mui/material'
import { useEffect } from 'react'

const IndexPage = () => {
  const { nowPlaying, error } = useNowPlaying()

  useEffect(() => {
    if (!error) return
    console.error(error)
  }, [error])

  return (
    <Box sx={fullViewHeightStyle}>
      <TrackListAppBar />

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

      <TrackList sx={{ height: 0.8 }} />
    </Box>
  )
}

export default IndexPage
