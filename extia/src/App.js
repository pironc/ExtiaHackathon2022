import logo from './logo.svg';
import React from 'react';
import './App.css';

// import json from './input.json';

function httpGet(url) {
  let xmlHttpReq = new XMLHttpRequest();
  xmlHttpReq.open("GET", url, false);
  xmlHttpReq.send(null);
  return xmlHttpReq.responseText;
}
const json = JSON.parse(httpGet('https://localhost:8000/front?city1=Paris&city2=Barcelona'));

function App() {
  return (
    <div>
      <header className="navbar">
        <select className="dropdown">
          <option value="angers">Angers</option>
          <option selected value="barcelona">Barcelona</option>
          <option value="paris">Paris</option>
          <option value="lyon">Lyon</option>
        </select>
        Vs.
        <select className="dropdown">
          <option value="angers">Angers</option>
          <option value="barcelona">Barcelona</option>
          <option selected value="paris">Paris</option>
          <option value="lyon">Lyon</option>
        </select>
      </header>
      <header className="content">
        {
          json.map(city => {
            return(
              <div className="widget">
                <a className="title">{ city.city }</a>
                <div className="content">Rent (1 bed city center) : { city.rent }e</div>
                <div className="content">Beer (in a bar) : { city.rent }e</div>
              </div>
            )
          })
        }
      </header>
    </div>
  );
}

export default App;
