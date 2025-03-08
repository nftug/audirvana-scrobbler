import { atom, PrimitiveAtom, useAtomValue, useSetAtom } from 'jotai'
import { useCallback, useEffect, useState } from 'react'

const actionPortalElementAtom = atom<HTMLDivElement | null>(null)

const useSetElementAtom = (elementAtom: PrimitiveAtom<HTMLDivElement | null>) => {
  const setElement = useSetAtom(elementAtom)
  const [element, setElementState] = useState<HTMLDivElement | null>(null)

  const refCallback = useCallback((node: HTMLDivElement | null) => {
    if (node !== null) setElementState(node)
  }, [])

  useEffect(() => {
    setElement(element)
  }, [element, setElement])

  return refCallback
}

export const useActionElement = () => useAtomValue(actionPortalElementAtom)

export const useActionElementRef = () => useSetElementAtom(actionPortalElementAtom)
