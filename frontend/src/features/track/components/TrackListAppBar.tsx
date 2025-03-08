import useErrorHandler from '@/lib/api/useErrorHandler'
import { useActionElement } from '@/lib/layout/atoms/portalAtom'
import { ScrobbleAll } from '@bindings/app/trackinfoservice'
import { Send } from '@mui/icons-material'
import { IconButton } from '@mui/material'
import { useQueryClient } from '@tanstack/react-query'
import { createPortal } from 'react-dom'
import { useTrackListQuery } from '../hooks/useTrackListQuery'

const TrackListAppBar = () => {
  const actionElement = useActionElement()
  const queryClient = useQueryClient()
  const handleError = useErrorHandler()
  const { data } = useTrackListQuery()

  const scrobbleAll = async () => {
    const error = await ScrobbleAll()
    if (error) handleError(error)
    queryClient.invalidateQueries({ queryKey: ['trackList'] })
  }

  return (
    <>
      {actionElement &&
        createPortal(
          <IconButton onClick={scrobbleAll} disabled={!data?.some((x) => !x.scrobbledAt)}>
            <Send />
          </IconButton>,
          actionElement
        )}
    </>
  )
}

export default TrackListAppBar
