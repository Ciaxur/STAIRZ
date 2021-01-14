import React from 'react';
import {
  makeStyles,
  Grid, Typography,
} from '@material-ui/core';

// Component Imports
import Activity from './Activity';


const useStyles = makeStyles(theme => ({
  root: {
    marginTop: 20,
    flexGrow: 1,
    overflow: 'hidden',
  },
  detail: {
    textAlign: 'left',
    color: theme.palette.text.secondary,
  },
}));


export default function(): JSX.Element {
  // Hooks
  const styles = useStyles();

  return (
    <div className={styles.root}>
      <Grid
        container
        spacing={2}
        direction='row'
        justify='center'
        alignItems='flex-start'
      >

        <Activity title='Title' renderBody={() => (
          <Typography className={styles.detail} >
            Details
          </Typography>
        )} />
        
      </Grid>

    </div>
  );
}