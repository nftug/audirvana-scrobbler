import { TextField, TextFieldProps } from '@mui/material'
import { FieldPath, FieldValues, useFormContext, useWatch } from 'react-hook-form'

interface Props<TFieldValues extends FieldValues>
  extends Omit<TextFieldProps, 'onChange' | 'value'> {
  name: FieldPath<TFieldValues>
}

const FormTextField = <TFieldValues extends FieldValues>({
  name,
  ...textFieldProps
}: Props<TFieldValues>) => {
  const { register, formState, control } = useFormContext()
  const value = useWatch({ control, name })

  return (
    <TextField
      {...register(name)}
      {...textFieldProps}
      error={!!formState.errors[name]}
      helperText={formState.errors[name]?.message as string}
      slotProps={{ inputLabel: { shrink: !!value } }}
    />
  )
}

export default FormTextField
