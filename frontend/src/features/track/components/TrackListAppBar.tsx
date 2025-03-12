import { useLoginState } from '@/features/settings/api/useLogin'
import useErrorHandler from '@/lib/api/useErrorHandler'
import { useActionElement } from '@/lib/layout/atoms/portalAtom'
import { ScrobbleAll } from '@bindings/app/trackinfoservice'
import { Send } from '@mui/icons-material'
import { IconButton } from '@mui/material'
import { useQueryClient } from '@tanstack/react-query'
import { createPortal } from 'react-dom'
import { getTrackListQueryKey, useTrackListQuery } from '../hooks/useTrackListQuery'

const TrackListAppBar = () => {
  const actionElement = useActionElement()
  const queryClient = useQueryClient()
  const handleError = useErrorHandler()
  const { data } = useTrackListQuery()
  const { data: loggedIn } = useLoginState()

  const scrobbleAll = async () => {
    const error = await ScrobbleAll()
    if (error) handleError(error)
    queryClient.invalidateQueries({ queryKey: getTrackListQueryKey() })
  }

  return (
    <>
      {actionElement &&
        createPortal(
          <IconButton
            onClick={scrobbleAll}
            disabled={!loggedIn || !data?.some((x) => !x.scrobbledAt)}
          >
            <Send />
          </IconButton>,
          actionElement
        )}
    </>
  )
}

export default TrackListAppBar
