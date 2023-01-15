import styled from "styled-components";
import { GitHub, LinkedIn } from "@mui/icons-material";
import { Link } from "@mui/material";

const FooterContainer = styled.div`
    text-align: center;
    position: absolute;
    bottom: 0;
    width: 100% !important;
    margin-bottom: 10px;
`;

const Footer = () => {
  return (
      <FooterContainer>
        <Link href="https://github.com/JoeLanglands">
            <GitHub color="secondary"/>
        </Link>
        <Link href="https://www.linkedin.com/in/joe-langlands-3686891b0/">
            <LinkedIn color="secondary" />
        </Link>
      </FooterContainer>
  );
};

export default Footer;
