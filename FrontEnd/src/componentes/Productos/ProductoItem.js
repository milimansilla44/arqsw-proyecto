import React from "react";
import {addToCart} from "../cookies/add2car"
import Cookies from "universal-cookie";
import swal from "sweetalert";

function notLogin(){
    swal("Debe loguearse")
    window.location.replace("/")
}
const Cookie = new Cookies();
let cookie = Cookie.get("username")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
export const ProductoItem =(
    {
    product_id,
    name,
    picture_url,
    product_unit_price,
    stock,
    description,
    Category_id
})=>{


    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
            <img className="image" src={picture_url} alt=""/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{name}</h1>
            {stock>0?<p>Stock: {stock}</p>:
            <p>Sin Stock</p>}
            
            <p className="price">$ {product_unit_price}</p>
        </div>
        {
        <div className="buttom">
        {stock>0?
            id_user!="undefined"?
            <button className="btn" onClick={()=> addToCart(product_id)}>
                Agregar
            </button>:
            <button className="btn" onClick={()=>notLogin()}>
            Agregar
        </button>:
            <p>Este Producto no esta Disponible</p>}
        </div>
        
        }
        </div>
    )
}