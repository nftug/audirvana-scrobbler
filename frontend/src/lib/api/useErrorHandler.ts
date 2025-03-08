import { ErrorResponse } from '@bindings/app/bindings/models'
import MessageDialog from '../dialog/MessageDialog'

const useErrorHandler = () => {
  return (error: ErrorResponse) => {
    MessageDialog.call({
      title: 'Error',
      message: error.data?.at(0)?.message ?? error.code,
      buttonType: 'ok'
    })
  }
}

export default useErrorHandler
