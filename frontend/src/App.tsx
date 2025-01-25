import { HeaderProvider } from '@/lib/layout/components/HeaderContext'
import TheDrawer from '@/lib/layout/components/TheDrawer'
import TheHeader from '@/lib/layout/components/TheHeader'
import AboutPage from '@/pages/AboutPage'
import IndexPage from '@/pages/IndexPage'
import SettingsPage from '@/pages/SettingsPage'
import { Box, createTheme, CssBaseline, ThemeProvider, Toolbar } from '@mui/material'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ConfirmProvider } from 'material-ui-confirm'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import TrackEditModal from './features/track/components/TrackEditModal'

const App: React.FC = () => {
  const theme = createTheme({ colorSchemes: { dark: true } })
  const confirmOptions = { confirmationText: 'OK', cancellationText: 'Cancel' } as const
  const queryClient = new QueryClient()

  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <ConfirmProvider defaultOptions={confirmOptions}>
          <HeaderProvider>
            <TheHeader />
            <TheDrawer />
          </HeaderProvider>

          <QueryClientProvider client={queryClient}>
            <TrackEditModal.Root />

            <Box component="main">
              <Toolbar />
              <Routes>
                <Route index element={<IndexPage />} />
                <Route path="/about" element={<AboutPage />} />
                <Route path="/settings" element={<SettingsPage />} />
              </Routes>
            </Box>
          </QueryClientProvider>
        </ConfirmProvider>
      </ThemeProvider>
    </BrowserRouter>
  )
}

export default App
