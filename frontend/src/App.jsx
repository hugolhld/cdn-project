import { AuthProvider } from './context/AuthContext'
import { BrowserRouter as Router, Routes, Route } from 'react-router'
import { SnackbarProvider } from 'notistack'
import ProtectedRoute from './utils/ProtectedRoutes'
import Home from './views/Home'
import Login from './views/Login'
import Register from './views/Register'

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
          </Routes>
        </AuthProvider>
      </Router>
    </SnackbarProvider>
  )
}

export default App
