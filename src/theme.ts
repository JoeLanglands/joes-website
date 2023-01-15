import { createTheme } from "@mui/material/styles";


const theme = createTheme({
  palette: {
    // type: 'light',
    primary: {
      main: '#2b1d38',
      contrastText: '#54e484',
    },
    secondary: {
      main: '#b141f1',
    },
    background: {
      default: '#ffffff',
    },
    text: {
      primary: '#6071cc',
      secondary: '#58c7e0',
    },
    divider: '#f91aad',
  },
  typography: {
    h1: {
      fontFamily: 'Fira Code',
      fontSize: '5rem',
    },
  },
});

export default theme;