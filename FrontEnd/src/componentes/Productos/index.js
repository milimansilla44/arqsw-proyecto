import React, { useContext, useEffect, useState} from "react";
import { ProductoItem } from "./ProductoItem";
import { ProductosBuscador } from "./buscador";
import { Link } from "react-router-dom";


export const ProductosLista = () => {

    const [productos,setProductos] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/product')
    .then((response) => response.json());
    setProductos(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])
    return(
        <>

        <div className="todos">
        <ul>
        <li>
          <Link to="/cart" ><h1>MI CARRITO</h1></Link>
        </li>
        </ul>
        </div>
        <div className="todos">
        <ul>
        <li>
          <Link to="/order" ><h1>MIS COMPRAS</h1></Link>
        </li>
        </ul>
        </div>
        <ProductosBuscador/>
        
        
        <h2>NUESTROS PRODUCTOS</h2>
        <div className="productos">
            {
                productos.map(producto =>(
                  <ProductoItem 
                  key={producto.id}
                  product_id={producto.product_id}
                  name={producto.name}
                  product_unit_price={producto.product_unit_price}
                  Category_id={producto.Category_id}
                  stock={producto.stock}
                  picture_url={producto.picture_url}
                  description={producto.description}
                  /> 
                ))
            }
            
        </div>
        </>
    )
}

