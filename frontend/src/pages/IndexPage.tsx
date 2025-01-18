import TrackEditModal from '@/features/track/components/TrackEditModal'
import TrackList from '@/features/track/components/TrackList'
import { fullViewHeightStyle } from '@/lib/layout/styles'
import { TrackInfo } from '@bindings/app/bindings'
import { useState } from 'react'

const IndexPage = () => {
  const [selectedTrack, setSelectedTrack] = useState<TrackInfo>()
  const [openEdit, setOpenEdit] = useState(false)

  const onClickEdit = (item: TrackInfo) => {
    setSelectedTrack(item)
    setOpenEdit(true)
  }

  const onClickDelete = (itemId: string) => {}

  return (
    <>
      <TrackList sx={fullViewHeightStyle} onClickEdit={onClickEdit} onClickDelete={onClickDelete} />
      <TrackEditModal item={selectedTrack} open={openEdit} onClose={() => setOpenEdit(false)} />
    </>
  )
}

export default IndexPage
