import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS, CATEGORIES } from '../../constants';

const Categories = ({ stateManager }) => {

    const {
      ops: {
        [CATEGORIES]: {
          getUncategorizedExpenditures,
          // setExpenditurePage,
        }
      },
      state: {
        [CATEGORIES]: {
          test,
        }
      }
    } = stateManager;

  useEffect(() => {
    getUncategorizedExpenditures();
  }, []);

  return (
    <div className="flex centered header-nav">
      Catgeorize goez hurr
    </div>
  );
}

export default Categories;
