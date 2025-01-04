import { createContext, Dispatch, SetStateAction, useState } from 'react'

export const DrawerContext = createContext<boolean>(false)
export const DrawerDispatchContext = createContext<Dispatch<SetStateAction<boolean>>>(() => {})

export const HeaderProvider = ({ children }: { children?: React.ReactNode }) => {
  const [drawerOpened, setDrawerOpened] = useState(false)

  return (
    <DrawerContext.Provider value={drawerOpened}>
      <DrawerDispatchContext.Provider value={setDrawerOpened}>
        {children}
      </DrawerDispatchContext.Provider>
    </DrawerContext.Provider>
  )
}
