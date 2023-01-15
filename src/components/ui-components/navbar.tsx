import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { AppBar, Typography, Box, Tabs, Tab, Avatar } from "@mui/material";

import MyImage from "../../images/joe.jpg";

interface LinkTabProps {
  label?: string;
  href?: string;
  value?: string;
}

const LinkTab = (props: LinkTabProps) => {
  const navigate = useNavigate();
  return (
    <Tab
      component="a"
      onClick={(event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
        if (props.href) {
          navigate(props.href);
        }
        event.preventDefault();
      }}
      {...props}
    />
  );
};

const Navbar = () => {
  const [value, setValue] = useState("home");
  const handleChange = (event: React.SyntheticEvent, value: string) => {
    setValue(value);
  };

  return (
    <AppBar position="static" sx={{ paddingTop: "20px" }}>
      <Box
        maxWidth="xl"
        sx={{ display: "grid", gap: 1, gridTemplateColumns: "repeat(4, 1fr)", gridTemplateAreas: `". avatar header header"
        ". tabs tabs tabs"` }}
      >
        <Box gridArea={"avatar"}>
          <Avatar
            alt="Joe Langlands"
            src={MyImage}
            sx={{ width: 150, height: 150 }}
          />
        </Box>

        <Box gridArea={"header"}>
          <Typography variant="h1" component="div" sx={{paddingTop: '20px'}}>
            &lt;joe langlands/&gt;
          </Typography>
        </Box>

        <Box gridArea={"tabs"}>
          <Tabs
            value={value}
            onChange={handleChange}
            textColor="secondary"
            indicatorColor="secondary"
            aria-label="navigation bar"
          >
            <LinkTab value="home" label="Home" href="/" />
            <LinkTab value="about" label="About" href="/about" />
            <LinkTab value="projects" label="Projects" href="/projects" />
          </Tabs>
        </Box>
      </Box>
    </AppBar>
  );
};

export default Navbar;
