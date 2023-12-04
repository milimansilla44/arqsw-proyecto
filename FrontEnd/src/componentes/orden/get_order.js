import React, {useState, useEffect} from "react";
import Cookies from "universal-cookie";
import { OrdenItem } from "./order_item";
import swal from "sweetalert2";


const Cookie = new Cookies();

async function GetOrdersWithDetailsByUserId(id) {
    return fetch('http://localhost:8090/ordersWithDetails/' +id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
   }
export const GetOrders = ()=>{
    
    let cookie = Cookie.get("username")
    let id_user;
    let token;
    if(cookie!=undefined){
    let array = cookie.split(",")
    token = array[0]
     id_user=array[1]
    }
    else{
         id_user = "undefined"
    }
    const [ordenes,setOrdenes]=useState([]);
        async function Handle (id) {
            const response = await GetOrdersWithDetailsByUserId(id)
            if (response.status == 400) {
                swal.fire({
                    text: "No ha realizado ninguna orden",
                    icon: 'warning'
                   
                }).then((result)=>{
                    if(result.isConfirmed){
                        window.location.replace("/")
                    }
                })
                
             }else{
                setOrdenes(response)
             }
            
            };
    useEffect(()=>{
        if(id_user!="undefined"){
        Handle(token);
        }
        else{
            swal.fire({
                text: "Debe Loguearse",
                icon: 'warning'
            }).then((result) => {
                if (result.isConfirmed) {
                    window.location.replace("/");
                }})
            
        }
        },[])
        return(
            <>
            <div className="help">
            <h1 className="title1">MIS COMPRAS</h1>
            </div>

            <div>

                <div className="ordenes">
            
            {
                ordenes.map(orden =>(
                    <div>
                        <OrdenItem key={orden.order_id}
                    order_id = {orden.order_id}
                    total_price = {orden.total_price}
                    date = {orden.date}
                    user_id = {orden.user_id}
                    details = {orden.details}
                    /> 
                    </div>
                )
                )
            }
            </div>
            </div>
            
            </>      
        )
    }
   