import { useLoginRoot } from '@/features/settings/api/useLogin'
import TrackList from '@/features/track/components/TrackList'
import TrackListAppBar from '@/features/track/components/TrackListAppBar'
import TrackNowPlaying from '@/features/track/components/TrackNowPlaying'
import { fullViewHeightStyle } from '@/lib/layout/styles'
import { Stack } from '@mui/material'

const IndexPage = () => {
  useLoginRoot()

  return (
    <>
      <TrackListAppBar />

      <Stack sx={fullViewHeightStyle} spacing={1}>
        <TrackNowPlaying />
        <TrackList />
      </Stack>
    </>
  )
}

export default IndexPage
