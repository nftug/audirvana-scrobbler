import * as yup from 'yup'

export const loginFieldSchema = yup.object().shape({
  username: yup.string().required(),
  password: yup.string().required()
})

export type LoginFieldSchemaType = yup.InferType<typeof loginFieldSchema>
