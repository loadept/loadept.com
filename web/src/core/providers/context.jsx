import { useState } from 'preact/hooks';
import { DataContext } from './dataContext';

export const DataProvider = ({ children }) => {
  const [data, setData] = useState('')

  return (
    <DataContext.Provider value={{ data, setData }} >
      {children}
    </DataContext.Provider>
  )
}
