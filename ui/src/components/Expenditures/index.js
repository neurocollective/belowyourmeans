import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS } from '../../constants';


const DisplayExpenditures = ({ expenditures }) => {
  return (
    <ul>
      {expenditures.map((e) => {
        return (
          <li>{e.value}</li>
        );
      })}
    </ul>
  );
};

const Expenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          getExpenditures,
        }
      },
      state: {
        [EXPENDITURES]: {
          expenditures,
        }
      }
    } = stateManager;

  useEffect(() => {
    getExpenditures();
  }, []);

  return (
    <div className="flex centered header-nav">
      <DisplayExpenditures expenditures={expenditures} />
    </div>
  );
}

export default Expenditures;
