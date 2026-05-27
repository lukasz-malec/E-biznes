import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import axios from 'axios'
import './Auth.css'

function Login({ onLogin }) {
  const [formularz, setFormularz] = useState({ email: '', password: '' })
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

    axios.post('http://localhost:8080/auth/login', formularz)
        .then(res => {
            onLogin(res.data.user, res.data.token)
            navigate('/')
        })
      .catch(() => {
        setBlad('Nieprawidłowy email lub hasło')
      })
      .finally(() => setLadowanie(false))
  }

  return (
    <div className="auth-wrapper">
      <div className="auth-card">
        <h1 className="auth-title">Logowanie</h1>

        <form onSubmit={handleSubmit}>
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
            {ladowanie ? 'Logowanie...' : 'Zaloguj się'}
          </button>
        </form>

        <p className="auth-link">
          Nie masz konta? <Link to="/register">Zarejestruj się</Link>
        </p>
      </div>
    </div>
  )
}

export default Login