import { useEffect, useState } from 'react'

function App() {
  const [message, setMessage] = useState("");
  useEffect(()=>{
    const response = fetch("http://localhost:3000/health");
    response.then(res=>res.json()).then(data=>{
      setMessage(data.message);
    })
  },[])
  return (
    <div className="App">
      Hello from Atmos.

      Server Status : {message}
    </div>
  )
}

export default App
