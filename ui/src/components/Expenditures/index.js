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

const Expenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          // expenditures,
          setMonth,
          getExpenditures,
        },
        [CATEGORIES]: {
          getCategories,
          categories,
        }
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

  useEffect(() => {
    getExpenditures();
  }, [month]);

  useEffect(() => {
    getCategories();
  }, [categories]);

  if (loading) {
    return (
      <section className="flex centered header-nav">
        Loading...
      </section>
    );    
  }

  return (
    <section className="flex centered header-nav">
      <div>
        <div>
          <span>Month:</span>
          &nbsp;
          <select value={month} onChange={(e) => setMonth(e.target.value)}>
            {MONTHS.map((monthName, index) => <option key={monthName} value={index}>{monthName}</option>)}
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
              <li className="expenditure-list-item" key={ex.id}>
                <div className="expenditure-list-item-details-container">
                  <div>
                    <div className="expenditure-line expenditure-amount">
                      ${String(value).replace("-", "")}
                    </div>
                    <div className="expenditure-line">
                      {description}
                    </div>
                  </div>
                </div>
                <div className="expenditure-line">
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
    </section>
  );
}

export default Expenditures;
