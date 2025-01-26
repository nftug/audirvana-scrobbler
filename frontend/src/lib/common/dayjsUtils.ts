import dayjs from 'dayjs'
import duration from 'dayjs/plugin/duration'

dayjs.extend(duration)

export const formatDateTime = (datetime?: string, useFullNotation = false) => {
  if (!datetime) return 'N/A'

  const dt = dayjs(datetime)
  const now = dayjs()

  if (useFullNotation || dt.year() !== now.year()) {
    return dt.format('YYYY/MM/DD HH:mm:ss')
  } else if (dt.date() === now.date()) {
    return dt.format('[Today] HH:mm:ss')
  } else if (now.diff(dt, 'day') === 1) {
    return dt.format('[Yesterday] HH:mm:ss')
  } else if (now.diff(dt, 'day') < 7) {
    return dt.format('dddd HH:mm:ss')
  } else {
    return dt.format('MM/DD HH:mm:ss')
  }
}

export const formatDuration = (seconds: number): string => {
  const duration = dayjs.duration(seconds, 'seconds')
  return duration.format('mm:ss')
}
