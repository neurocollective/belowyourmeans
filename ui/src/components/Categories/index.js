import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS, CATEGORIES } from '../../constants';
import SetCategory from '../SetCategory';

const Categories = ({ stateManager }) => {

    const {
      ops: {
        [CATEGORIES]: {
          getUncategorizedExpenditures,
          getCategories,
          // setExpenditurePage,
        }
      },
      state: {
        [CATEGORIES]: {
          test,
          categories,
          expenditureNames,
          expenditureNamePage,
        }
      }
    } = stateManager;

  useEffect(() => {
    getUncategorizedExpenditures();
  }, []);

  useEffect(() => {
    getCategories();
  }, []);

  const pageTimesTen = 10 * expenditureNamePage;

  const expenditurePage = expenditureNames.slice(pageTimesTen, pageTimesTen + 10);

  return (
    <div className="flex centered header-nav">
      <div>
        <div>
          <ul style={{ "listStyle": "none" }}>
            {expenditurePage.map((expenditure) => {
              return (
                <li key={expenditure.id}>
                  {expenditure.description}
                  {/*<SetCategory
                    notCategorized={notCategorized}
                    categoryName={categoryName}
                    select={setSelectedCategory}
                    update={updateCategoryForExpenditure}
                    categories={categories}
                    expenditureId={expenditureId}
                    expenditureDescription={expenditureDescription}
                    selectionsMap={selectionsMap}
                  />*/}
                </li>
              );
            })}
          </ul>       
        </div>
        <div>
          <button>{"<-"}</button>
          &nbsp;
          Page: {String(expenditureNamePage + 1)}
          &nbsp;
          <button>{"->"}</button>
        </div>
      </div>
    </div>
  );
}

export default Categories;
