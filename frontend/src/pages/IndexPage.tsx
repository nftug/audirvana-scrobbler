import TrackList from '@/features/track/components/TrackList'
import { fullViewHeightStyle } from '@/lib/layout/styles'

const IndexPage = () => {
  const onClickEdit = (itemId: string) => {}
  const onClickDelete = (itemId: string) => {}

  return (
    <>
      <TrackList sx={fullViewHeightStyle} onClickEdit={onClickEdit} onClickDelete={onClickDelete} />
    </>
  )
}

export default IndexPage
