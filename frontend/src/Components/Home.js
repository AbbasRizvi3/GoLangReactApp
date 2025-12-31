import React from 'react'
import './Auth.css'
import { Link } from 'react-router-dom'

export default function Home() {
  return (
    <div className="home-hero">
      <div className="home-inner">
        <h1 className="hero-title">Welcome to GoLangReactApp</h1>
        <p className="hero-sub">A simple starter app with Go backend and React frontend.</p>
        <div className="hero-actions">
          <Link to="/login" className="btn primary">Login</Link>
          <Link to="/signup" className="btn outline">Sign Up</Link>
        </div>
      </div>
    </div>
  )
}
