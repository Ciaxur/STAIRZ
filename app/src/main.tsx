import React from 'react';
import ReactDOM from 'react-dom';
import { makeStyles } from '@material-ui/core';

// Component Imports
import ActivityBar from './Components/ActivityBar';
import ProcessNodes from './Components/ProcessNode';

const useStyles = makeStyles({
  root: {},
});

const App = () => {
  const styles = useStyles();

  
  return (
    <div className={styles.root}>
      <ActivityBar />
      <ProcessNodes />
    </div>
  );
};

ReactDOM.render(
  <App />,
  document.getElementById('root'),
);