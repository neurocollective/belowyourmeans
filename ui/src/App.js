import './App.css';
import { useState } from 'react';
import state from './state';
import Header from './components/Header';
import Login from './components/Login';

const { INITIAL_STATE, buildStateManager } = state;

function App() {

  const [state, setState] = useState(INITIAL_STATE);

  const stateManager = buildStateManager(state, setState);

  return (
    <>
      <header className="App-header">
        <Header stateManager={stateManager} />
      </header>
      <main className="App">
        mane
        <Login stateManager={stateManager} />
      </main>
      <footer>
        footur
      </footer>
    </>
  );
}

export default App;
