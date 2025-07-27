import './App.css';
import { useState, useEffect } from 'react';
import StateStore from './state';
import Header from './components/Header';
import Login from './components/Login';
import Home from './components/Home';
import Router from './components/Router';
import Loading from './components/Loading';
import Footer from './components/Footer';
import Expenditures from './components/Expenditures';
import Categories from './components/Categories';
import { LOGIN, LOADING, HOME, EXPENDITURES, CATEGORIES, HEADER } from './constants';

const { INITIAL_STATE, buildStateManager } = StateStore;

function App() {

  const [state, setState] = useState(INITIAL_STATE);
  const stateManager = buildStateManager(state, setState);

  const {
    ops: {
      [LOGIN]: {
        checkIfLoggedIn,
      },
      [CATEGORIES]: {
        getCategories,
        getUncategorizedExpenditures,
      },
      [EXPENDITURES]: {
        getExpenditures,       
      }
    },
    state: {
      [LOGIN]: {
        user: userId,
      },
      [EXPENDITURES]: {
        month,
      },
      [CATEGORIES]: {
        categories,
      },
      [HEADER]: {
        selectedYear,
      }
    }
  } = stateManager;

  useEffect(() => {
    console.log('useEffect 1 in App.js');
    checkIfLoggedIn();
  }, []);

  useEffect(() => {
    console.log('useEffect 2 in App.js');
    if (userId) {
      // getUncategorizedExpenditures();
      getCategories();
    }
  }, [userId, selectedYear]);

  console.log("state in App.js:", state);

  return (
    <>
      <Header stateManager={stateManager} />
      <main className="App">
        <Router stateManager={stateManager}>
          <Login stateManager={stateManager} navigation={LOGIN} />
          <Home stateManager={stateManager} navigation={HOME} />
          <Loading stateManager={stateManager} navigation={LOADING} />
          <Expenditures stateManager={stateManager} navigation={EXPENDITURES} />
          <Categories stateManager={stateManager} navigation={CATEGORIES} />
        </Router>
      </main>
      <Footer stateManager={stateManager} />
    </>
  );
}

export default App;
