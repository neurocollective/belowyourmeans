import React, { useEffect } from 'react';
import { NAVIGATION, HOME, EXPENDITURES, LOGIN, REPORTS, CATEGORIES } from '../../constants';
import SetCategory from '../SetCategory';

const Categories = ({ stateManager }) => {

    const {
      ops: {
        [CATEGORIES]: {
          getUncategorizedExpenditures,
          getCategories,
          // updatePreAssignmentFullNameSelection,
          updatePreAssignmentPattern,
          // setSelectedBroadCategory,
          applyCategoryItem
          // setExpenditurePage,
        }
      },
      changes: {
        [CATEGORIES]: {
          updatePreAssignmentFullNameSelection,
        },
      },
      state: {
        [CATEGORIES]: {
          loading,
          test,
          categories,
          expenditureNames,
          expenditureNamePage,
          preAssignmentMap,
        }
      }
    } = stateManager;

  useEffect(() => {
    console.log('useEffect 1 in Categories.js');
    getUncategorizedExpenditures();
  }, [JSON.stringify(expenditureNames)]);

  // console.log(`size of expenditureNames: ${expenditureNames.length}`);

  const pageTimesTen = 10 * expenditureNamePage;

  const expenditurePageList = expenditureNames.slice(expenditureNamePage, expenditureNamePage + pageTimesTen);

  // console.log(`expenditureNamePage: ${expenditureNamePage} pageTimesTen: ${pageTimesTen} expenditureNamePage + pageTimesTen: ${expenditureNamePage + pageTimesTen}`);

  // console.log(`expenditurePageList size: ${expenditurePageList.length}`);

  if (loading) {
    return (
      <div className="flex centered header-nav">
        LOADING...
      </div>
    );
  }

  const applyCategories = (expenditureDescription, categoryName, expenditureId) => {
    console.log(expenditureDescription, categoryName, expenditureId);
    applyCategoryItem(expenditureDescription, categoryName, expenditureId);
  };

  const select = (id, value) => {
    updatePreAssignmentFullNameSelection(id, value);
  };

  const backPage = () => {

  };

  const forwardPage = () => {

  };

  return (
    <div className="flex centered header-nav">
      <div>
        <div>
          <ul style={{ "listStyle": "none" }}>
            {expenditurePageList.map((expenditure) => {

              const { description } = expenditure;

              const assignmentMapEntry = preAssignmentMap?.[expenditure.id] ?? {};

              return (
                <li key={expenditure.id}>
                  <div>
                    {expenditure.description}
                  </div>
                  <div>
                    <SetCategory
                      notCategorized={true}
                      categoryName={assignmentMapEntry.name || "default"}
                      categoryPattern={assignmentMapEntry.pattern ?? description}
                      fullName={assignmentMapEntry.fullName}
                      select={select}
                      update={applyCategories}
                      updatePattern={updatePreAssignmentPattern}
                      categories={categories}
                      expenditureId={expenditure.id}
                      expenditureDescription={description}
                      selectionsMap={preAssignmentMap}
                    />
                  </div>
                </li>
              );
            })}
          </ul>       
        </div>
        <div>
          {expenditureNamePage > 1 && <button onClick={backPage}>{"<-"}</button>}
          &nbsp;
          Page: {String(expenditureNamePage)}
          &nbsp;
          <button onClick={forwardPage}>{"->"}</button>
        </div>
      </div>
    </div>
  );
}

export default Categories;
