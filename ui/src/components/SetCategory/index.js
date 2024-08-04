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
    updatePattern,
    notCategorized,
    categories,
    categoryName,
    categoryPattern,
    expenditureId,
    expenditureDescription,
    selectionsMap,
    fullName = true,
    allowPatterns = true,
  } = props;

  if (notCategorized) {

    const selectedCategory = selectionsMap?.[expenditureId]?.name || "default";

    const submit = () => {
      update(expenditureDescription, categoryName, expenditureId);
    };
    const handleSelect = (e) => {
      console.log('select\'s expenditureId, e.target.value', expenditureId, e.target.value);
      select(expenditureId, e.target.value);
    }

    return (
      <React.Fragment>
        {fullName && (
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
        )}
        {allowPatterns && (
          <div>
            <input type="checkbox" value="fullName" checked={!fullName} onChange={() => {}}/>
            <label>Only use part of description text</label>
          </div>
        )}
        {!fullName && allowPatterns && (
          <input type="text" value={categoryPattern} />
        )}
      </React.Fragment>
    );
  }

  return null;
}

export default SetCategory;
