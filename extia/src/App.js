import logo from './logo.svg';
import './App.css';

import input from './input.json';
const obj = JSON.stringify(input);
const json = JSON.parse(obj);
var array = []

function App() {
  for (const key in json) {
    array.push(key)
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          {array}
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
