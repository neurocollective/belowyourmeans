import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS, CATEGORIES } from '../../constants';
import SetCategory from '../SetCategory';

const Categories = ({ stateManager }) => {

    const {
      ops: {
        [CATEGORIES]: {
          getUncategorizedExpenditures,
          getCategories,
          setSelectedBroadCategory,
          applyCategoryItem
          // setExpenditurePage,
        }
      },
      state: {
        [CATEGORIES]: {
          loading,
          test,
          categories,
          expenditureNames,
          expenditureNamePage,
          broadSelectionsMap,
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

  if (loading) {
    return (
      <div className="flex centered header-nav">
        LOADING...
      </div>
    );
  }

  return (
    <div className="flex centered header-nav">
      <div>
        <div>
          <ul style={{ "listStyle": "none" }}>
            {expenditurePage.map((expenditure) => {

              const { description } = expenditure;

              const select = (id, value) => {
                setSelectedBroadCategory(id, value);
              };

              const applyCategories = (expenditureDescription, categoryName, expenditureId) => {
                console.log(expenditureDescription, categoryName, expenditureId);
                applyCategoryItem(expenditureDescription, categoryName, expenditureId);
              };

              return (
                <li key={expenditure.id}>
                  <div>
                    {expenditure.description}
                  </div>
                  <div>
                    <SetCategory
                      notCategorized={true}
                      categoryName={broadSelectionsMap[expenditure.id] || "default"}
                      select={select}
                      update={applyCategories}
                      categories={categories}
                      expenditureId={expenditure.id}
                      expenditureDescription={description}
                      selectionsMap={broadSelectionsMap}
                    />
                  </div>
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
