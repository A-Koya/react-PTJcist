import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { AddUser } from './components/page/AddUser.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AddUser />
  </React.StrictMode>,
)
