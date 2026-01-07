import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'


export default function Dashboard() {
  const [data,setData]=useState({})
  const API_BASE_URL = process.env.REACT_APP_API_BASE_URL
  const navigate=useNavigate()

  function logoutHandler(){
    fetch(`${API_BASE_URL}/Logout`,{
      method:"POST",
      credentials:"include",
    }).then((res)=>{
      if(res.ok){
        console.log("Logged out successfully")
        navigate("/Login")
      }
    })
  }

  useEffect(()=>{
    fetch(`${API_BASE_URL}/Dashboard`,{
      method:"GET",
      credentials:"include",
    }).then((res)=>{
      if(!res.ok){
        navigate("/Login")
      }
      return res.json()
    }).then((data)=>{
      console.log(data)
      setData(data)
    })
  },[])

 return (
    <>
      <div>Dashboard</div>
      <ul>
        {data?.msg?.map((item, index) => (
          <li key={index}>{item}</li>
        ))}</ul>

        <button onClick={logoutHandler}>Logout</button>
    </>
  );
}
