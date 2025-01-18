import { TrackInfoForm } from '@bindings/app/bindings/models'
import * as yup from 'yup'

export const trackInfoFieldSchema = yup.object<TrackInfoForm>().shape({
  artist: yup.string().required(),
  album: yup.string().required(),
  track: yup.string().required()
})
