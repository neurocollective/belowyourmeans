import React, { useEffect } from 'react';
import Categorizer from './Categorizer';
import {
  MONTHS,
  NAVIGATION,
  HOME,
  EXPENDITURES,
  LOGIN,
  REPORTS,
  CATEGORIES,
  IGNORED,
  UNCATEGORIZED,
} from '../../constants';

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
      <ul class="expenditure-list">
        {expenditures.map((ex, index) => {

          const { value, description, ['category_name']: categoryName } = ex;

          return (
            <li class="expenditure-list-item" key={ex.id}>
              <div class="expenditure-list-item-details-container">
                <div>
                  <div class="expenditure-line expenditure-amount">
                    ${String(value).replace("-", "")}
                  </div>
                  <div class="expenditure-line">
                    {description}
                  </div>
                </div>
              </div>
              <div class="expenditure-line">
                <Categorizer
                  categoryName={categoryName}
                  expenditureDescription={description}
                  expenditureId={ex.id}
                  stateManager={stateManager}
                />
              </div>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

const Expenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          expenditures,
          getExpenditures,
        },
        [CATEGORIES]: {
          getCategories,
        }
      },
      state: {
        [EXPENDITURES]: {
          month,
          loading,
        }
      }
    } = stateManager;

  useEffect(() => {
    getExpenditures();
  }, [month]);

  useEffect(() => {
    getCategories();
  }, [])

  if (loading) {
    return (
      <section className="flex centered header-nav">
        Loading...
      </section>
    );    
  }

  return (
    <section className="flex centered header-nav">
      <DisplayExpenditures stateManager={stateManager} />
    </section>
  );
}

export default Expenditures;
