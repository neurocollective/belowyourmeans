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

const CategoryContent = (props) => {

  const {
    setSelectedCategory,
    updateCategoryForExpenditure,
    notCategorized,
    categories,
    categoryName,
    expenditureId,
    expenditureDescription,
    selectionsMap,
  } = props;

  if (notCategorized) {

    const selectedCategory = selectionsMap[expenditureId] || "default"

    const submit = () => {
      console.log('CLICK, motherfucker');
      updateCategoryForExpenditure(expenditureDescription, categoryName, expenditureId);
    };
    const select = (e) => setSelectedCategory(expenditureId, e.target.value);

    return (
      <React.Fragment>
        <select
          className="set-category-dropdown"
          value={selectedCategory}
          onChange={select}
        >
          <option disabled value="default">Select A Category</option>
          {categories.map((c) => {
            const name = c["display_name"]
            return <option key={c.id} value={name}>{name}</option>
          })}
        </select>
        &nbsp;
        <button
          className="set-category-button"
          onClick={submit}
        >
          Set Category
        </button>
      </React.Fragment>
    );
  }

  return null;
}
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
        <CategoryContent
          notCategorized={notCategorized}
          categoryName={categoryName}
          setSelectedCategory={setSelectedCategory}
          updateCategoryForExpenditure={updateCategoryForExpenditure}
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
