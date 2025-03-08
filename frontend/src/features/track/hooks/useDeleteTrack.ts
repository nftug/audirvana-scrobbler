import useErrorHandler from '@/lib/api/useErrorHandler'
import MessageDialog from '@/lib/dialog/MessageDialog'
import { TrackInfoResponse } from '@bindings/app/bindings'
import { DeleteTrackInfo } from '@bindings/app/trackinfoservice'
import { Delete } from '@mui/icons-material'
import { Stack, Table, TableBody, TableCell, TableRow, Typography } from '@mui/material'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { createElement } from 'react'

const useDeleteTrack = () => {
  const handleError = useErrorHandler()
  const queryClient = useQueryClient()

  const mutation = useMutation({
    mutationFn: async (id: number) => {
      const error = await DeleteTrackInfo(id)
      if (error) throw error
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['trackList'] })
    },
    onError: handleError
  })

  return mutation
}

export const useDeleteTrackDialog = () => {
  const deleteTrack = useDeleteTrack()

  return async (track: TrackInfoResponse) => {
    const rows = [
      { name: 'Artist', value: track.artist },
      { name: 'Album', value: track.album },
      { name: 'Track', value: track.track }
    ]

    const ok = await MessageDialog.call({
      message: createElement(
        Stack,
        { spacing: 2 },
        createElement(Typography, null, 'Delete this track?'),
        createElement(
          Table,
          null,
          createElement(
            TableBody,
            null,
            rows.map((t) =>
              createElement(
                TableRow,
                { key: t.name },
                createElement(TableCell, null, t.name),
                createElement(TableCell, null, t.value)
              )
            )
          )
        )
      ),
      buttonType: 'okCancel',
      confirmText: 'Delete',
      confirmButtonProps: {
        variant: 'contained',
        color: 'error',
        startIcon: createElement(Delete)
      }
    })
    if (!ok) return

    deleteTrack.mutate(track.id)
  }
}
