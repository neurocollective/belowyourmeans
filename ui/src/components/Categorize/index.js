import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS, CATEGORIZE } from '../../constants';

const Categorize = ({ stateManager }) => {

    const {
      ops: {
        [CATEGORIZE]: {
          getUncategorizedExpenditures,
          // setExpenditurePage,
        }
      },
      state: {
        [CATEGORIZE]: {
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

export default Categorize;
