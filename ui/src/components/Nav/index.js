import React from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS } from '../../constants';

const Nav = ({ stateManager }) => {

  const {
    ops: {
      [NAVIGATION]: {
        navigate,
      },
      [LOGIN]: {
        logout,
      }
    },
  } = stateManager;

  return (
    <nav className="flex centered header-nav">
      <div className="nav-link-container">
        <a className="nav-link" href="#" onClick={(e) => {e.preventDefault(); navigate(EXPENDITURES);}}>
          Expenditures
        </a>
      </div>
      <div className="nav-link-container">
        <a  className="nav-link" href="#" onClick={(e) => {e.preventDefault(); navigate(REPORTS);}}>
          Reports
        </a>
      </div>
      <div className="nav-link-container">
        <a  className="nav-link" href="#" onClick={(e) => {e.preventDefault(); logout;}}>
          Sign Out
        </a>
      </div>
    </nav>
  );
}

export default Nav;
