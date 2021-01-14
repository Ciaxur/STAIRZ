import React from 'react';
import {
  makeStyles,
  AppBar,
  Toolbar,
  Typography,
} from '@material-ui/core';

const useStyles = makeStyles({
  quadrant_sm: {
    width: '20%',
  },
  quadrant_md: {
    width: '60%',
  },
  leftAlign: { 
    textAlign: 'left', 
  },
  center: { 
    textAlign: 'center', 
  },
  rightAlign: { 
    textAlign: 'right', 
  },
  title: {
    fontSize: 14,
  },
});


export default function ActivityBar(): JSX.Element {
  // Hooks
  const styles = useStyles();

  // States
  const [ time, setTime ] = React.useState<Date>(new Date());

  // Effects
  React.useEffect(() => { // On Mount/Unmount
    // Keep track of Time every second
    const timerID = setInterval(() => setTime(new Date()), 1000);
    
    // Clean up Method
    return () => {
      clearInterval(timerID);     // Clean up Timer
    };
  });
  
  return (
    <AppBar position='static'>
      <Toolbar variant='dense'>
        <div className={ `${styles.quadrant_sm} ${styles.leftAlign}` } >
          <Typography className={styles.title} variant='h6'>
            STAIRZ
          </Typography>
        </div>
        <div className={ `${styles.quadrant_md} ${styles.center}` } >
          <Typography className={styles.title} variant='h6'>
            {time.toLocaleDateString()} {time.toLocaleTimeString()}
          </Typography>
        </div>
        <div className={ `${styles.quadrant_sm} ${styles.rightAlign}` } >
          <Typography className={styles.title} variant='h6'>
            P3
          </Typography>
        </div>
      </Toolbar>
    </AppBar>
  );
}