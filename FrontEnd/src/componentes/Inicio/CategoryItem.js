import React from "react";

export const CategoryItem =(
    {Id,
    Name,
})=>{
    return(
        <div className="categoria">
        <div className="categoria_footer">
            <h1>{Name}</h1>
        </div>
        </div>
    )
}