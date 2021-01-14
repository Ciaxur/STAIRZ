import React from 'react';
import {
  makeStyles,
  Container, Paper, Grid, Typography,
} from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  root: {
    marginTop: 20,
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(2),
  },
  title: {
    textAlign: 'center',
    color: theme.palette.text.primary,
  },
  
}));

interface Props {
  title: string,
  renderBody: () => JSX.Element,
}


export default function ({ title, renderBody }: Props): JSX.Element {
  // Hooks
  const styles = useStyles();

  return (
    <Grid item sm={3} xs={6}>
      <Paper className={styles.paper} variant='outlined'>
        <Typography className={styles.title} variant='h6' >
          {title}
        </Typography>

        {renderBody()}
      </Paper>
    </Grid>
  );
}