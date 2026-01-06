import React, { useEffect, useState } from 'react'


export default function Dashboard() {
  const [data,setData]=useState([])
  const API_BASE_URL = process.env.API_BASE_URL

  useEffect(()=>{
    fetch(`${API_BASE_URL}/Dashboard`,{
      method:"GET",
      credentials:"include",
    }).then((res)=>res.json()).then((data)=>{
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
        ))}
      </ul>
    </>
  );
}
