import TrackList from '@/features/track/components/TrackList'
import TrackListAppBar from '@/features/track/components/TrackListAppBar'
import useNowPlaying from '@/features/track/hooks/useNowPlaying'
import { fullViewHeightStyle, overflowEllipsisStyle } from '@/lib/layout/styles'
import { Card, CardContent, Stack, Typography } from '@mui/material'
import { useEffect } from 'react'

const IndexPage = () => {
  const { nowPlaying, error } = useNowPlaying()

  useEffect(() => {
    if (!error) return
    console.error(error)
  }, [error])

  return (
    <>
      <TrackListAppBar />

      <Stack sx={fullViewHeightStyle} spacing={1}>
        <Card sx={{ width: 1, overflowX: 'hidden' }}>
          <CardContent component={Stack} spacing={1} sx={{ mb: 1.5 }}>
            <Typography gutterBottom sx={{ color: 'text.secondary', fontSize: 14 }}>
              Now playing
            </Typography>

            <Typography variant="h5" component="div" sx={overflowEllipsisStyle}>
              {!nowPlaying ? 'No track playing' : nowPlaying.track}
            </Typography>
            <Typography sx={{ color: 'text.secondary', ...overflowEllipsisStyle }}>
              {!nowPlaying
                ? 'N/A'
                : `${nowPlaying.artist ?? 'No artist'} - ${nowPlaying.album ?? 'No album'}`}
            </Typography>
          </CardContent>
        </Card>

        <TrackList sx={{ height: 1 }} />
      </Stack>
    </>
  )
}

export default IndexPage
