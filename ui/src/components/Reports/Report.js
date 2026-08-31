import React, { useEffect } from 'react';
import { CATEGORIES, EXPENDITURES, HEADER, REPORTS } from '../../constants';

const Report = ({ report, priorReport, getLine }) => {

  if (!report) {
    return null;
  }

  return (
    <div>
      {report.map((line) => {

        const [name, average] = getLine(line);

        return (
          <div key={name}>
            <span>{name}</span>
            <span>:</span>
            &nbsp;
            <span>{average}</span>
            {priorReport && priorReport[name] && (
              <span>prior year: {average}</span>
            )}
          </div>
        );
      })}
    </div>
  );
};

export default Report;

// http://localhost:8080/api/report/annualized?year=2025&months=6