import Navbar from "./components/ui-components/navbar";
import { ThemeProvider } from "@mui/system";
import { Routes, Route} from "react-router-dom";
import theme from "./theme";

import Home from "./components/Home";
import About from "./components/About";
import Projects from "./components/Projects";
import Footer from "./components/ui-components/footer";

function App() {
  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Navbar />
        <Routes>
          <Route index element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/projects" element={<Projects/>} />
        </Routes>
        <Footer/>
      </div>
    </ThemeProvider>
  );
}

export default App;
