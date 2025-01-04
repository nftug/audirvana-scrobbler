import { fullViewHeightStyle } from '@/lib/layout/styles'
import { Box, Typography } from '@mui/material'

const AboutPage = () => {
  return (
    <Box sx={fullViewHeightStyle}>
      <Box display="flex" justifyContent="center" alignItems="center" height={1}>
        <Typography variant="h2" color="textSecondary">
          Audirvana Scrobbler
        </Typography>
      </Box>
    </Box>
  )
}

export default AboutPage
