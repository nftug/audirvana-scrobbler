import DialogProvider from '@/lib/layout/components/DialogProvider'
import { HeaderProvider } from '@/lib/layout/components/HeaderContext'
import TheDrawer from '@/lib/layout/components/TheDrawer'
import TheHeader from '@/lib/layout/components/TheHeader'
import AboutPage from '@/pages/AboutPage'
import IndexPage from '@/pages/IndexPage'
import SettingsPage from '@/pages/SettingsPage'
import { Box, createTheme, CssBaseline, ThemeProvider, Toolbar } from '@mui/material'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { SnackbarProvider } from 'notistack'
import { BrowserRouter, Route, Routes } from 'react-router-dom'

const theme = createTheme({ colorSchemes: { dark: true } })

const App = () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchOnWindowFocus: false
      }
    }
  })

  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <SnackbarProvider
          anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
          preventDuplicate
        >
          <QueryClientProvider client={queryClient}>
            <HeaderProvider>
              <TheHeader />
              <TheDrawer />
            </HeaderProvider>

            <DialogProvider />

            <Box component="main">
              <Toolbar />
              <Routes>
                <Route index element={<IndexPage />} />
                <Route path="/about" element={<AboutPage />} />
                <Route path="/settings" element={<SettingsPage />} />
              </Routes>
            </Box>
          </QueryClientProvider>
        </SnackbarProvider>
      </ThemeProvider>
    </BrowserRouter>
  )
}

export default App
