import { useEffect, useState } from 'react'

function App() {
  const [message, setMessage] = useState("");
  const [coordinates,setCoordinates] = useState({
    "Latitude":"22.390",
    "Longitude":"88.146",
  })
  const [isGPSAllowed,setIsGPSAllowed] = useState(false);
  useEffect(()=>{
    const response = fetch("http://localhost:3000/health");
    response.then(res=>res.json()).then(data=>{
      setMessage(data.message);
    })
  },[])
  function getLatLong(){
    if(navigator.geolocation){
      navigator.geolocation.getCurrentPosition(success,error)
    }
    else{
      setIsGPSAllowed(false)
    }
  }
  function success(pos:any){
    setIsGPSAllowed(true)
    setCoordinates({
      "Latitude": pos.coords.latitude.toString(),
      "Longitude": pos.coords.longitude.toString()
    })
  }
  function error(error:any){
    setIsGPSAllowed(false)
    if(error.code === error.PERMISSION_DENIED){
      alert("Please allow GPS access to get accurate weather data.")
    }
    else if(error.code === error.POSITION_UNAVAILABLE){
      alert("Location information is unavailable.")
    }
  }
  return (
    <div className="App">
      Hello from Atmos.

      <div>
        Server Status : {message}
      </div>
      <div>
        <button onClick={getLatLong}>Get Current Location</button>
        {isGPSAllowed && <div>
          <p>Latitude: {coordinates.Latitude}</p>
          <p>Longitude: {coordinates.Longitude}</p>
        </div>}
      </div>
    </div>
  )
}

export default App
