import React from 'react';
import {
  HEADER,
  NAVIGATION,
  HOME,
  EXPENDITURES,
  LOGIN,
  REPORTS,
  CATEGORIES,
} from '../../constants';

const Nav = ({ stateManager }) => {

  const {
    ops: {
      [NAVIGATION]: {
        navigate,
      },
      [LOGIN]: {
        logout,
      },
    },
    changes: {
      [HEADER]: {
        handleYearSelect,
      } = {}
    },
    state: {
      [HEADER]: {
        years,
        selectedYear = '2024',
      } = {}
    },
  } = stateManager;

  const goto = category => e => {
    e.preventDefault();
    navigate(category);
  };

  return (
    <div>
      <nav className="flex centered header-nav">
        <div className="nav-link-container">
          <a className="nav-link" href="#" onClick={goto(CATEGORIES)}>
            Categories
          </a>
        </div>
        <div className="nav-link-container">
          <a className="nav-link" href="#" onClick={goto(EXPENDITURES)}>
            Expenditures
          </a>
        </div>
        <div className="nav-link-container">
          <a  className="nav-link" href="#" onClick={goto(REPORTS)}>
            Reports
          </a>
        </div>
        <div className="nav-link-container">
          <a  className="nav-link" href="#" onClick={(e) => {e.preventDefault(); logout();}}>
            Sign Out
          </a>
        </div>
      </nav>
      <nav className="flex centered upcentered header-nav">
        <label>Year:</label>
        <select onChange={handleYearSelect} value={selectedYear}>
          {years.map(year => (<option value={year}>{year}</option>))}
        </select>
      </nav>
    </div>
  );
}

export default Nav;
