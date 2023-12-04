import React, {useState} from "react";
import { ProductoItem } from "./ProductoItem";
import swal from "sweetalert2";
import "./buscador.css";

export const ProductosBuscador = ()=>{

    const [productos,setProductos] = useState([]);
    const [busqueda, setBusqueda]= useState("");
    
    const fetchApi = async()=>{
      
        const response = await fetch('http://localhost:8090/productXpalabraClave/'+busqueda)
        .then((response) => response.json());
        if (response.status == 400) {
          swal.fire({
            icon: 'error',
            text: "No se encontro el producto",
          }).then((result) => {
            if (result.isConfirmed) {
                window.location.reload();
            }})
       }else{
        
        setProductos(response)
        console.log(response);
       }
        };

    const handleChange=e=>{
     setBusqueda(e.target.value);
     
   
      };

      const handleSubmit= (event)=>{
        event.preventDefault();
        
      
        fetchApi();

    };
    
    return(
        <>
        <h1 className="title"></h1>
        <div className="containerInput" >
        <input
           
          className="form-controlinputBuscar"
          value={busqueda}
          placeholder="Ingresar Producto"
          onChange={handleChange}
         
        />
        <input
        className="botonBuscar"
        value = "Buscar"
         type = "button"
        onClick = {handleSubmit}
        />
      </div>
        <div className="vacio">
        <div className="productos">
            {
                productos.map(producto =>(
                  <ProductoItem  key={producto.id}
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
        </div>
        </>
    )
}