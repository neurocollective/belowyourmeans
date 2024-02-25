// import './App.css';
// import { useState, useEffect } from 'react';
// import StateStore from './state';
// import Header from './components/Header';
// import Login from './components/Login';
// import { LOGIN } from './constants';

// const { INITIAL_STATE, buildStateManager } = StateStore;

function Nav({ stateManager }) {

  return (
    <nav className="flex centered header-nav">
      <div className="nav-link-container">
        <a className="nav-link" href="#" onClick={(e) => {e.preventDefault(); console.log('clicky');}}>
          Expenditures
        </a>
      </div>
      <div className="nav-link-container">
        <a  className="nav-link" href="#" onClick={(e) => {e.preventDefault(); console.log('clicky');}}>
          Reports
        </a>
      </div>
      <div className="nav-link-container">
        <a  className="nav-link" href="#" onClick={(e) => {e.preventDefault(); console.log('clicky');}}>
          Sign Out
        </a>
      </div>
    </nav>
  );
}

export default Nav;
