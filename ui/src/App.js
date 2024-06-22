import './App.css';
import { useState, useEffect } from 'react';
import StateStore from './state';
import Header from './components/Header';
import Login from './components/Login';
import Home from './components/Home';
import Router from './components/Router';
import Loading from './components/Loading';
import Footer from './components/Footer';
import { LOGIN, LOADING, HOME } from './constants';

const { INITIAL_STATE, buildStateManager } = StateStore;

function App() {

  const [state, setState] = useState(INITIAL_STATE);
  const stateManager = buildStateManager(state, setState);

  useEffect(() => {
    stateManager.ops[LOGIN].checkIfLoggedIn();
  }, []);

  console.log("state in App.js:", state);

  return (
    <>
      <Header stateManager={stateManager} />
      <main className="App">
        <Router stateManager={stateManager}>
          <Login stateManager={stateManager} navigation={LOGIN} />
          <Home stateManager={stateManager} navigation={HOME} />
          <Loading stateManager={stateManager} navigation={LOADING} />
        </Router>
      </main>
      <Footer stateManager={stateManager} />
    </>
  );
}

export default App;
