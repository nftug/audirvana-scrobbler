import { overflowEllipsisStyle } from '@/lib/layout/styles'
import { Card, CardContent, Stack, Typography } from '@mui/material'
import { useEffect } from 'react'
import useNowPlaying from '../hooks/useNowPlaying'

const TrackNowPlaying = () => {
  const { nowPlaying, error } = useNowPlaying()

  useEffect(() => {
    if (!error) return
    console.error(error)
  }, [error])

  return (
    <Card sx={{ width: 1, overflowX: 'hidden', minHeight: 120, borderRadius: 0 }}>
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
  )
}

export default TrackNowPlaying
