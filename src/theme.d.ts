declare module "@mui/material/styles" {
  interface Theme {
    palette: {
      type: string;
      primary: {
        main: string;
        contrastText: string;
      };
      secondary: {
        main: string;
      };
      background: {
        default: string;
      };
      text: {
        primary: string;
        secondary: string;
      };
      divider: string;
    };
    typography: {
      h1: {
        fontFamily: string;
        fontSize: string;
      };
    };
  }

  // allow configuration using `createTheme`
  // interface ThemeOptions {
  //   status?: {
  //     danger?: string;
  //   };
  // }
}
