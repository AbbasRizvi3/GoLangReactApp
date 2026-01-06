import React, { useState } from 'react'
import './Auth.css'
import { Link} from 'react-router-dom'
import { useNavigate } from 'react-router-dom'

export default function Login() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  function handleLogin({ email, password }) {
    fetch("http://localhost:8000/Login",{
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    })
    .then((r) => {
      if (r.ok) {
        r.json().then((user) => {
          console.log(user)
          localStorage.setItem("token", user.token);
        });
        navigate("/Dashboard");
      }
      else {
        r.json().then((err) => setError(err.error));
      } 
    })
  }

  function handleSubmit(e) {
    e.preventDefault()
    setError('')
    if (!email) return setError('Please enter your email')
    if (!password || password.length < 6) return setError('Password must be at least 6 characters')
      handleLogin({ email, password })
    console.log('Logging in', { email, password })
  }

  return (
    <div className="auth-page">
      <div className="auth-card">
        <h2 className="auth-title">Sign in</h2>
        <form onSubmit={handleSubmit} className="auth-form">
          <label className="auth-label">Email</label>
          <input className="auth-input" type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="you@example.com" />

          <label className="auth-label">Password</label>
          <input className="auth-input" type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="••••••••" />

          {error && <div className="auth-error">{error}</div>}

          <button className="auth-button" type="submit">Login</button>
        </form>

        <div className="auth-footer">
          <span>New here? <Link to="/signup">Create an account</Link></span>
        </div>
      </div>
    </div>
  )
}
