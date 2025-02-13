import { AuthProvider } from './context/AuthContext'
import { BrowserRouter as Router, Routes, Route } from 'react-router'
import { SnackbarProvider } from 'notistack'
import ProtectedRoute from './utils/ProtectedRoutes'
import Home from './views/Home'
import Login from './views/Login'
import Register from './views/Register'
import FileBrowser from './views/Folder'
import { TreeView } from './components/tree-view'

const data = [
  {
    id: '1',
    name: 'Item 1',
    children: [
      {
        id: '2',
        name: 'Item 1.1',
        children: [
          {
            id: '3',
            name: 'Item 1.1.1',
          },
          {
            id: '4',
            name: 'Item 1.1.2',
          },
        ],
      },
      {
        id: '5',
        name: 'Item 1.2',
      },
    ],
  },
  {
    id: '6',
    name: 'Item 2',
  },
];


function App() {
  return (
    <SnackbarProvider>
      <Router>
        <AuthProvider>
          <Routes>
            <Route
              path="/"
              element={
                <ProtectedRoute>
                  <Home />
                </ProtectedRoute>
              }
            />
            <Route path="/login" element={<Login />} />
            <Route path='/register' element={<Register />} />
            <Route path='/test' element={<FileBrowser />} />
            <Route path='/test2' element={<TreeView data={data} />} />
          </Routes>
        </AuthProvider>
      </Router>
    </SnackbarProvider>
  )
}

export default App
