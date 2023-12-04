import React from "react";
import Mac from "../../images/labialmac.jpg";
import { Link } from "react-router-dom";

export const Header = () => {
  
  return (
    <header>
      <Link to="/inicio">
      <div className="logo">
        <img src={Mac} alt="" width="150" />
      </div>
      </Link>
      <ul>
        <li>
          <Link to="/inicio">INICIO</Link>
        </li>
        <li>
          <Link to="/productos">PRODUCTOS</Link>
        </li>
      </ul>
      <div className="cart">
        <box-icon name="cart"></box-icon>
      </div>
    </header>
  );
};