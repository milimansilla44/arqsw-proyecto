import React from "react";

export const OrdenItem =(
    {   order_id,
        date,
        total_price,
        user_id,
        details
})=>{
    return(
      <div className="ordenes">
        <div className="orden">
          <table>
         <thead>
           <tr>
             <th>Monto Final</th>
             <th>Fecha</th>
           </tr>
         </thead>
         <tbody>
            <tr>
               <td>{total_price}</td>
               <td>{date}</td>
             </tr>
         </tbody>
       </table>
       <h2>Detalle de la compra</h2>
       <table>
        <thead>
           <tr>
             <th>Precio Unitario</th>
             <th>Cantidad</th>
             <th>Total</th>
           </tr>
         </thead>
         <tbody>
         {
           details.map((detail) =>(      
                <tr>
                <td>{detail.price}</td>
                <td>{detail.quantity}</td>
                <td>{detail.total_price}</td>
              </tr>
                ))
           }
         </tbody>
       </table>
       <br></br>
        </div>
        <div>
        </div>
        </div>
    )
}