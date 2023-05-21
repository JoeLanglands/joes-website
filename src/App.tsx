import Navbar from "./components/ui-components/navbar";
import { ThemeProvider } from "@mui/system";
import { ThemeProvider as ScThemeProvider } from "styled-components";
import { Routes, Route} from "react-router-dom";
import theme from "./theme";

import Home from "./components/Home";
import About from "./components/About";
import Projects from "./components/Projects";
import Footer from "./components/ui-components/footer";

// Put in horrible double theme provider hack so that styled components can use the theme

function App() {
  return (
    <ThemeProvider theme={theme}>
      <ScThemeProvider theme={theme}>
        <div className="App">
          <Navbar />
          <Routes>
            <Route index element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/projects" element={<Projects />} />
          </Routes>
          <Footer />
        </div>
      </ScThemeProvider>
    </ThemeProvider>
  );
}

export default App;
