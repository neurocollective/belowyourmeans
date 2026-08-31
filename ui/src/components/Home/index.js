import React, { useEffect } from 'react';
import { CATEGORIES, EXPENDITURES } from '../../constants'

const Home = ({ stateManager }) => {

  // const { stateManager: { state } } = props;

  const {
    ops: {
      [CATEGORIES]: {
        getCategories,
      },
      [EXPENDITURES]: {
        getExpenditures,       
      }
    },
    state: {
      [EXPENDITURES]: {
        month,
      },
      [CATEGORIES]: {
        categories,
      }
    }
  } = stateManager;

  // useEffect(() => {
  //   getExpenditures();
  // }, [month]);

  // useEffect(() => {
  //   getCategories();
  // }, [categories]);

  // useEffect(() => {
  //   getUncategorizedExpenditures();
  // }, []);

  return (
    <section>
      <br />
      <br />
      <br />
      i am the main stuff
      <br />
      <br />
      <br />
    </section>
  );
};

export default Home;
