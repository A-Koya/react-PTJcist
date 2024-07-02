import { Route, Routes } from "react-router-dom"
import { Login } from "./components/page/Login"
import { AddUser } from "./components/page/AddUser"

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/addUser" element={<AddUser />} />
    </Routes>
  )
}

export default App
