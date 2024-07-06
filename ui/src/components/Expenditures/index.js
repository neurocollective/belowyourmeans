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
} from '../../constants';

const CategoryContent = (props) => {

  const {
    displayCategorizationOption,
    notCategorized,
    setSelectedCategory,
    categories,
    category,
    updateCategory,
  } = props;

  if (notCategorized || displayCategorizationOption) {
    return (
      <React.Fragment>
        <select value={category} onChange={() => setSelectedCategory(category)}>
          {categories.map((c) => {
            return <option value={c["display_name"]}>{c["display_name"]}</option>
          })}
        </select>
        &nbsp;
        <button onChange={() => updateCategory()}>Set Category</button>
        <div>
          <input name="categorize-all" value={true} type="radio" onChange={() => {}} />
          <label>
            In the future, set same category for all expenditures with this exact description
          </label>
        </div>
        <div>
          <input name="categorize-all" value={false} type="radio" onChange={() => {}} />
          <label>
            Set category only for this single transaction
          </label>

        </div>
      </React.Fragment>
    );
  }

  return null;

  // for later

  // return (
  //   <React.Fragment>
  //     {category}
  //     &nbsp;
  //     <button onChange={() => {}}>Change Category</button>
  //   </React.Fragment>
  // );
}
const Categorizer = ({ stateManager, category, expenditureId }) => {

  const {
    ops: {
      [CATEGORIES]: {
        setSelectedCategory,
        updateCategory,
      },
    },
    state: {
      [EXPENDITURES]: {
        bruh,
      },
      [CATEGORIES]: {
        categories,
        expenditureIdToCategorize,
      }
    }
  } = stateManager;

  const notCategorized = category === 'uncategorized';
  const displayCategorizationOption = expenditureIdToCategorize === expenditureId;

  return (
    <div>
      <div>
        <span>{category}</span>
      </div>
      <div>
        <CategoryContent
          notCategorized={notCategorized}
          displayCategorizationOption={displayCategorizationOption}
          category={category}
          setSelectedCategory={setSelectedCategory}
          updateCategory={updateCategory}
          categories={categories}
        />
      </div>
    </div>
  );
}

const DisplayExpenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          setMonth,
        },
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

    if (loading) {
      return (
        <div>
          Loading...
        </div>
      );
    }

  return (
    <div>
      <div>
        <span>Month:</span>
        &nbsp;
        <select value={month} onChange={(e) => setMonth(e.target.value)}>
          {MONTHS.map((monthName, index) => <option value={index}>{monthName}</option>)}
        </select>
      </div>
      <div>
        {expenditures.length} results
      </div>
      {expenditures.length ? <div>Page {expenditurePage}</div> : null}
      <ul class="expenditure-list">
        {expenditures.map((ex, index) => {

          const { value, description, category = 'uncategorized' } = ex;

          return (
            <li class="expenditure-list-item" key={ex.id}>
              <div class="expenditure-list-item-details-container">
                <div>
                  <div class="expenditure-line expenditure-amount">
                    ${String(value).replace("-", "")}
                  </div>
                  <div class="expenditure-line">
                    {description}
                  </div>
                </div>
              </div>
              <div class="expenditure-line">
                <Categorizer
                  category={category}
                  expenditureId={ex.id}
                  stateManager={stateManager}
                />
              </div>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

const ExpendituresNav = (props) => {
  return (
    <nav> 
      {"nav goes hurr"}
    </nav>
  );
}

const Expenditures = ({ stateManager }) => {

    const {
      ops: {
        [EXPENDITURES]: {
          getExpenditures,
          // setExpenditurePage,
        },
        [CATEGORIES]: {
          getCategories,
        }
      },
      state: {
        [EXPENDITURES]: {
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
  }, [])

  if (loading) {
    return (
      <section className="flex centered header-nav">
        Loading...
      </section>
    );    
  }

  return (
    <section className="flex centered header-nav">
      <DisplayExpenditures stateManager={stateManager} />
    </section>
  );
}

export default Expenditures;
