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

const SetCategory = (props) => {

  const {
    select,
    update,
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
      update(expenditureDescription, categoryName, expenditureId);
    };
    const handleSelect = (e) => select(expenditureId, e.target.value);

    return (
      <React.Fragment>
        <select
          className="set-category-dropdown"
          value={selectedCategory}
          onChange={handleSelect}
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

export default SetCategory;
