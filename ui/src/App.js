import './App.css';
import { useState, useEffect } from 'react';
import StateStore from './state';
import Header from './components/Header';
import Login from './components/Login';
import { LOGIN } from './constants';

const { INITIAL_STATE, buildStateManager } = StateStore;

function App() {

  const [state, setState] = useState(INITIAL_STATE);
  // stateManager = { state, ops, changes }
  const stateManager = buildStateManager(state, setState);

  useEffect(() => {
    stateManager.ops[LOGIN].checkIfLoggedIn();
  }, []);

  console.log("state in App.js:", state);

  return (
    <>
      <Header stateManager={stateManager} />
      <main className="App">
        <Login stateManager={stateManager} />
        <section>
          <br />
          <br />
          <br />
          i am the main stuff
          <br />
          <br />
          <br />
        </section>
      </main>
      <footer className="App-footer">
        <a class="bym-link" href="mailto:david@neurocollective.io">Contact Us</a>
      </footer>
    </>
  );
}

export default App;
