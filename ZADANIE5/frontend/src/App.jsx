import { useState, useEffect } from 'react'
import { Routes, Route, Link, NavLink, useLocation, Navigate } from 'react-router-dom'
import Produkty from './components/Produkty'
import Koszyk from './components/Koszyk'
import Platnosci from './components/Platnosci'
import Login from './components/Login'
import Register from './components/Register'
import './App.css'

function App() {
  // React Hook useState na 4.0
  // koszyk przesalany jako props do komponentow
  const [koszyk, setKoszyk] = useState([])
  const [user, setUser] = useState(null)
  const location = useLocation()

  // odczyt tokenu z localStorage przy starcie aplikacji
  useEffect(() => {
    const token = localStorage.getItem('token')
    const userData = localStorage.getItem('user')
    if (token && userData) {
      setUser(JSON.parse(userData))
    }
  }, [])

  const authPages = ['/login', '/register']
  const isAuthPage = authPages.includes(location.pathname)

  const dodajDoKoszyka = (produkt) => {
    setKoszyk(prev => {
      const istniejacy = prev.find(p => p.id === produkt.id)
      if (istniejacy) {
        return prev.map(p =>
          p.id === produkt.id ? { ...p, quantity: p.quantity + 1 } : p
        )
      }
      return [...prev, { ...produkt, quantity: 1 }]
    })
  }

  const usunZKoszyka = (id) => {
    setKoszyk(prev => prev.filter(p => p.id !== id))
  }

  const handleLogin = (userData, token) => {
    localStorage.setItem('token', token)
    localStorage.setItem('user', JSON.stringify(userData))
    setUser(userData)
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    setUser(null)
  }

  const ChronionaTrasa = ({ children }) => {
    if (!user)
      return <Navigate to="/login" replace />
    
    return children
  }

  const suma = koszyk.reduce((sum, p) => sum + p.price * p.quantity, 0)

  return (
    <>
      {!isAuthPage && (
        <nav className="nav">
          <Link to="/" className="nav-logo">Sklep React JS + GO</Link>
          <div className="nav-links">
            <NavLink to="/" end>Produkty</NavLink>
            <NavLink to="/koszyk">Koszyk</NavLink>
            <NavLink to="/platnosci">Płatności</NavLink>
          </div>
          <div className="nav-auth">
            <span className="nav-user">{user?.name}</span>
            <button className="nav-logout" onClick={handleLogout}>Wyloguj</button>
          </div>
          <Link to="/koszyk" className="nav-cart">
            Koszyk ({koszyk.length}) · {suma.toFixed(2)} zł
          </Link>
        </nav>
      )}

      <Routes>
        <Route path="/login" element={<Login onLogin={handleLogin} />} />
        <Route path="/register" element={<Register onLogin={handleLogin} />} />

        <Route path="/" element={
          <ChronionaTrasa>
            <Produkty dodajDoKoszyka={dodajDoKoszyka} />
          </ChronionaTrasa>
        } />
        <Route path="/koszyk" element={
          <ChronionaTrasa>
            <Koszyk koszyk={koszyk} usunZKoszyka={usunZKoszyka} />
          </ChronionaTrasa>
        } />
        <Route path="/platnosci" element={
          <ChronionaTrasa>
            <Platnosci koszyk={koszyk} />
          </ChronionaTrasa>
        } />
      </Routes>
    </>
  )
}

export default App