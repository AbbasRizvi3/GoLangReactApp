import React, { useState } from 'react'
import './Auth.css'
import { Link } from 'react-router-dom'

export default function Signup() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirm, setConfirm] = useState('')
  const [error, setError] = useState('')
  const API_BASE_URL = process.env.REACT_APP_API_BASE_URL

  function handleSignup({ name, email, password }) {
    console.log("API_BASE_URL =", API_BASE_URL) 
    fetch(`${API_BASE_URL}/Signup`,{
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, email, password }),
    })
    .then((r) => {
      if (r.ok) {
        r.json().then((user) => console.log(user));
      } else {
        r.json().then((err) => setError(err.error));
      } 
    })
  }

  function handleSubmit(e) {
    e.preventDefault()
    setError('')
    if (!name) return setError('Please enter your name')
    if (!email) return setError('Please enter your email')
    if (!password || password.length < 6) return setError('Password must be at least 6 characters')
    if (password !== confirm) return setError('Passwords do not match')
    handleSignup({ name, email, password })
    // Frontend-only: pretend to submit
    console.log('Signing up', { name, email })
  }

  return (
    <div className="auth-page">
      <div className="auth-card">
        <h2 className="auth-title">Create account</h2>
        <form onSubmit={handleSubmit} className="auth-form">
          <label className="auth-label">Full name</label>
          <input className="auth-input" value={name} onChange={e => setName(e.target.value)} placeholder="Jane Doe" />

          <label className="auth-label">Email</label>
          <input className="auth-input" type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="you@example.com" />

          <label className="auth-label">Password</label>
          <input className="auth-input" type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Minimum 6 characters" />

          <label className="auth-label">Confirm password</label>
          <input className="auth-input" type="password" value={confirm} onChange={e => setConfirm(e.target.value)} placeholder="Repeat password" />

          {error && <div className="auth-error">{error}</div>}

          <button className="auth-button" type="submit">Sign up</button>
        </form>

        <div className="auth-footer">
          <span>Already have an account? <Link to="/login">Login</Link></span>
        </div>
      </div>
    </div>
  )
}
