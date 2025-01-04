import {
  Container,
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
  Stack,
  Typography,
  useColorScheme
} from '@mui/material'

const SettingsPage: React.FC = () => {
  const { mode, setMode } = useColorScheme()
  if (!mode) {
    return null
  }

  return (
    <Container sx={{ marginTop: 5 }}>
      <Stack spacing={3}>
        <Typography variant="h4">Settings</Typography>

        <FormControl>
          <FormLabel>Theme</FormLabel>
          <RadioGroup
            row
            value={mode}
            onChange={(event) => setMode(event.target.value as 'system' | 'light' | 'dark')}
          >
            <FormControlLabel value="system" control={<Radio />} label="System" />
            <FormControlLabel value="light" control={<Radio />} label="Light mode" />
            <FormControlLabel value="dark" control={<Radio />} label="Dark mode" />
          </RadioGroup>
        </FormControl>
      </Stack>
    </Container>
  )
}

export default SettingsPage
