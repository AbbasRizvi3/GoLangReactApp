import React, { useEffect } from 'react'

export default function Dashboard() {

  useEffect(()=>{
    fetch("http://localhost:8000/Dashboard").then((res)=>res.json()).then((data)=>{
      console.log(data)
    })
  })
  return (
    <div>Dashboard</div>
  )
}
