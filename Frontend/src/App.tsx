import { useEffect, useState } from 'react'
import AqiPanel from './components/AQIPanel';
import Header from './components/Header';
import LocationStrip from './components/LocationStrip';
import type {WeatherData, AirQualityData} from './types';
import WeatherHero from './components/weatherCard';
function App() {
  const [weatherData, setWeatherData] = useState<WeatherData| null>(null);
  const [airQualityData, setAirQualityData] = useState<AirQualityData|null>(null);
  const [coordinates,setCoordinates] = useState({
    "Latitude":"22.390",
    "Longitude":"88.146",
  })
  const [isGPSAllowed,setIsGPSAllowed] = useState(false);
  useEffect(()=>{
    if(isGPSAllowed){
      const weatherResponse = fetch(`http://localhost:3000/weatherData?latitude=${coordinates.Latitude}&longitude=${coordinates.Longitude}`);
      weatherResponse.then(res=>res.json()).then(data=>{
        setWeatherData(data); 
        
      })   

      const aqiResponse = fetch(`http://localhost:3000/airQualityData?latitude=${coordinates.Latitude}&longitude=${coordinates.Longitude}`);
      aqiResponse.then((res)=>res.json()).then(data=>{
        setAirQualityData(data) 
      })
    }
    else if(!isGPSAllowed){
      const weatherResponse = fetch("http://localhost:3000/weatherData");
      weatherResponse.then(res=>res.json()).then(data=>{
        setWeatherData(data);
      })   

      const aqiResponse = fetch("http://localhost:3000/airQualityData");
      aqiResponse.then((res)=>res.json()).then(data=>{
        setAirQualityData(data)
      })
    }
  },[coordinates])
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
    <div className='bg-[#333537] text-[#bebeb4] flex flex-col justify-center items-center'>
      <div className="App h-fit max-w-300 ">
      <Header/>
      {weatherData && <LocationStrip  data={weatherData}></LocationStrip>}
      <div>
        <button onClick={getLatLong}>Locate Me</button>
      </div>
      {weatherData && <WeatherHero {...weatherData}/>}
      {airQualityData && <AqiPanel {...airQualityData} />}
    </div>
    </div>
  )
}

export default App
