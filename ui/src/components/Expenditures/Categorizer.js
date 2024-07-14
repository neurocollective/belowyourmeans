import React, { useEffect } from 'react';
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
import SetCategory from '../SetCategory';

const Categorizer = ({ stateManager, categoryName, expenditureId, expenditureDescription }) => {

  const {
    ops: {
      [CATEGORIES]: {
        setSelectedCategory,
        updateCategoryForExpenditure,
      },
    },
    state: {
      [EXPENDITURES]: {
        bruh,
      },
      [CATEGORIES]: {
        categories,
        selectionsMap,
      }
    }
  } = stateManager;

  const notCategorized = categoryName === UNCATEGORIZED;
  // const displayCategorizationOption = expenditureIdToCategorize === expenditureId;

  const categoryClass = notCategorized ? 'display-category-uncategorized' : 'display-category-default';

  return (
    <div>
      <div className='display-category'>
        <span className={categoryClass}>{categoryName}</span>
      </div>
      <div>
        <SetCategory
          notCategorized={notCategorized}
          categoryName={categoryName}
          select={setSelectedCategory}
          update={updateCategoryForExpenditure}
          categories={categories}
          expenditureId={expenditureId}
          expenditureDescription={expenditureDescription}
          selectionsMap={selectionsMap}
        />
      </div>
    </div>
  );
}

export default Categorizer;
