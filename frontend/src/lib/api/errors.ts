import { ErrorCode } from '@bindings/app/shared/enums'

interface ApiErrorData {
  field: string
  message: string
}

interface ApiErrorJson {
  code: ErrorCode
  data?: ApiErrorData
}

export class ApiError extends Error {
  private constructor(
    public readonly code: ErrorCode,
    public readonly data?: ApiErrorData
  ) {
    super(code)
  }

  static create(errJson: ApiErrorJson) {
    return new ApiError(errJson.code, errJson.data)
  }
}
