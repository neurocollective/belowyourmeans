import './App.css';
import { useState } from 'react';
import state from './state';
import Header from './components/Header';
import Login from './components/Login';

const { INITIAL_STATE, buildStateManager } = state;

function App() {

  const [state, setState] = useState(INITIAL_STATE);

  console.log("state in App.js:", state);

  const stateManager = buildStateManager(state, setState);

  return (
    <>
      <header className="App-header">
        <Header stateManager={stateManager} />
      </header>
      <main className="App">
        <Login stateManager={stateManager} />
      </main>
      <footer className="App-footer">
        <a class="bym-link" href="mailto:david@neurocollective.io">Contact Us</a>
      </footer>
    </>
  );
}

export default App;
