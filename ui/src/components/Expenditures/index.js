import React, { useEffect } from 'react';
import { MONTHS, NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS } from '../../constants';


const DisplayExpenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          setMonth,
        },
      },
      state: {
        [EXPENDITURES]: {
          expenditures,
          expenditurePage = 1,
          month,
          loading,
        }
      }
    } = stateManager;

    if (loading) {
      return (
        <div>
          Loading...
        </div>
      );
    }

  return (
    <div>
      <div>
        <span>Month:</span>
        &nbsp;
        <select value={month} onChange={(e) => setMonth(e.target.value)}>
          {MONTHS.map((monthName, index) => <option value={index}>{monthName}</option>)}
        </select>
      </div>
      <div>
        {expenditures.length} results
      </div>
      {expenditures.length ? <div>Page {expenditurePage}</div> : null}
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
    </div>
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
          month,
        }
      }
    } = stateManager;

  useEffect(() => {
    getExpenditures();
  }, [month]);

  return (
    <section className="flex centered header-nav">
      <DisplayExpenditures stateManager={stateManager} />
    </section>
  );
}

export default Expenditures;
