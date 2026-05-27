import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import axios from 'axios'
import './Auth.css'

function Register({ onLogin }) {
  const [formularz, setFormularz] = useState({ name: '', email: '', password: '' })
  const [blad, setBlad] = useState(null)
  const [ladowanie, setLadowanie] = useState(false)
  const navigate = useNavigate()

  const handleChange = (e) => {
    setFormularz(prev => ({ ...prev, [e.target.name]: e.target.value }))
  }

  const handleSubmit = (e) => {
    e.preventDefault()
    setLadowanie(true)
    setBlad(null)

   axios.post('http://localhost:8080/auth/register', formularz)
    .then(res => {
        onLogin(res.data.user, res.data.token)
        navigate('/')
    })
      .catch(() => {
        setBlad('Błąd rejestracji — email może być już zajęty')
      })
      .finally(() => setLadowanie(false))
  }

  return (
    <div className="auth-wrapper">
      <div className="auth-card">
        <h1 className="auth-title">Rejestracja</h1>

        <form onSubmit={handleSubmit}>
          <div className="auth-pole">
            <label className="auth-label">Imię i nazwisko</label>
            <input
              className="auth-input"
              name="name"
              placeholder="Jan Kowalski"
              value={formularz.name}
              onChange={handleChange}
              required
            />
          </div>

          <div className="auth-pole">
            <label className="auth-label">Email</label>
            <input
              className="auth-input"
              name="email"
              type="email"
              placeholder="jan@example.com"
              value={formularz.email}
              onChange={handleChange}
              required
            />
          </div>

          <div className="auth-pole">
            <label className="auth-label">Hasło</label>
            <input
              className="auth-input"
              name="password"
              type="password"
              placeholder="••••••••"
              value={formularz.password}
              onChange={handleChange}
              required
            />
          </div>

          {blad && <p className="auth-blad">{blad}</p>}

          <button className="auth-btn" type="submit" disabled={ladowanie}>
            {ladowanie ? 'Rejestrowanie...' : 'Zarejestruj się'}
          </button>
        </form>

        <p className="auth-link">
          Masz już konto? <Link to="/login">Zaloguj się</Link>
        </p>
      </div>
    </div>
  )
}

export default Register