import Cookies from "universal-cookie";

const Cookie= new Cookies()

export function CookieUser(id_user,token){
  
    let cookie = Cookie.get("username");
   
    if(cookie == undefined){
      Cookie.set("username", id_user + "," + token + ";", {path: "/"});
      return
    }
    return
  }