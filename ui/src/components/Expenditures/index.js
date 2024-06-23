import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS } from '../../constants';


const DisplayExpenditures = ({ expenditures, expenditurePage }) => {
  return (
    <ul>
      {expenditures.map((ex) => {

        const { value, description, category = 'uncategorized' } = ex;

        return (
          <li>
            <div>
              {String(value).replace("-", "")}
            </div>
            <div>
              {description}
            </div>
            <div>
              {category}
            </div>
          </li>
        );
      })}
    </ul>
  );
};

const ExpendituresNav = (props) => {
  return (
    <nav> 
      {}
    </nav>
  );
}

const Expenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          getExpenditures,
          // setExpenditurePage,
        }
      },
      state: {
        [EXPENDITURES]: {
          expenditures,
          expenditurePage = 1,
        }
      }
    } = stateManager;

  useEffect(() => {
    getExpenditures();
  }, []);

  return (
    <div className="flex centered header-nav">
      <DisplayExpenditures expenditures={expenditures} page={expenditurePage} />
    </div>
  );
}

export default Expenditures;
