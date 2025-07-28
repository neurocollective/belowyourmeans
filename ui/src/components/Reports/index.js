import React, { useEffect } from 'react';
import { CATEGORIES, EXPENDITURES, HEADER, REPORTS } from '../../constants';
import Report from './Report';

const DISPLAY_NAME = 'display_name';
const MONTHLY_AVERAGE = 'monthly_average';

const getLine = (line) => {
  const { [DISPLAY_NAME]: name, [MONTHLY_AVERAGE]: average } = line;
  return [name, average];
};

const buildPriorMap = (priorReport) => {
  if (!priorReport || !priorReport.length) {
    return null;
  }
  return priorReport.reduce((accumulator, line) => {
    const [name, average] = getLine(line);
    accumulator[name] = average;
    return accumulator;
  }, {});
};

const Reports = ({ stateManager }) => {

  const {
    ops: {
      [REPORTS]: {
        getReports,
      },
    },
    state: {
      [REPORTS]: {
        report,
        priorReport,
      }
    }
  } = stateManager;

  console.log('state[REPORTS]', stateManager.state[REPORTS]);

  useEffect(() => {
    getReports();
  }, []);

  const priorMap = buildPriorMap(priorReport)

  return (
    <section class="flex centered">
      <Report report={report} priorMap={priorMap} getLine={getLine} />
    </section>
  );
};

export default Reports;

// http://localhost:8080/api/report/annualized?year=2025&months=6