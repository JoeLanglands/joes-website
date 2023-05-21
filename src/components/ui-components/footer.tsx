import styled from "styled-components";
import { GitHub, LinkedIn } from "@mui/icons-material";
import { Link } from "@mui/material";

const FooterContainer = styled.div`
  display: flex;
  justify-content: center;
  gap: 1em;
  flex-direction: row;
  position: absolute;
  bottom: 0;
  width: 100% !important;
  background-color: ${({ theme }) => theme.palette.primary.main};
  padding: 10px;
  right: 0;
`;

const Footer = () => {
  return (
    <FooterContainer>
      <Link href="https://github.com/JoeLanglands">
        <GitHub color="secondary" />
      </Link>
      <Link href="https://www.linkedin.com/in/joe-langlands-3686891b0/">
        <LinkedIn color="secondary" />
      </Link>
    </FooterContainer>
  );
};

export default Footer;
