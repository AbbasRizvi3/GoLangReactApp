import React, { useEffect, useState } from 'react'


export default function Dashboard() {
  const [data,setData]=useState([])

  useEffect(()=>{
    fetch("http://localhost:8000/Dashboard").then((res)=>res.json()).then((data)=>{
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
