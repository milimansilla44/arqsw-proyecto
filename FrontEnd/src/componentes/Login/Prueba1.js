
import React,{ useState} from "react";
import "./Prueba1.css";
import Cookies from 'universal-cookie';
import {CookieUser} from "../cookies/cookiesUser"
import swal from "sweetalert2";
const Cookie = new Cookies();

export default function GetUserByLogin(){
  
    const[username,setUserData]= useState("");
    const[password,setPassword] = useState("");
  
    const onChangeUser =  (username)=>{
        setUserData(username.target.value);
        
    }
    
    const onChangePas = (password)=>{
    setPassword(password.target.value)};

    
    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({'username': username, 'password': password})
      };

    const login = async()=>{
        fetch('http://127.0.0.1:8090/login',requestOptions)
        .then(response => {if (response.status == 400) {
           swal.fire({
            text: "Datos incorrectos",
            icon: 'error',
           }).then((result) => {
            if (result.isConfirmed) {
                window.location.reload();
                return response.json()
            }})
        }
        if(response.status==201){
          swal.fire({
            icon: 'success',
            text: "¡BIENVENIDO!"
          }
          ).then((result) => {
            if (result.isConfirmed) {
              window.location.replace("/inicio")
              return response.json()
            }})
        }
        return response.json()})
        .then(response => {
           Cookie.set("username", response.id_user + "," + response.token, {path: "/"})
    })
   
    };
   
    const handleSubmit= (event)=>{
        event.preventDefault();
        login();

    };


  
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Username </label>
          <input type="text" name="uid" required />
          <label> Password </label>
          <input type="password" name="upass" required />
        </div>
        <div className="button-container">
          <input type="submit" value="Ingresar" />
        </div>
      </form>
    </div>
  );

return(
  <div className="app">
    <div className="login-form">
    <form onSubmit={handleSubmit} >
    <ul className="ul">
    
    <h1 className="login">INGRESO</h1>
    
      <li>
    <input id="username" type={"text"} placeholder="Nombre de usuario" onChange={onChangeUser} value ={username} required></input>
      </li>
    <li>
    <input id="password" type={"password"} placeholder="Contraseña" onChange={onChangePas} value={password} required></input>
    </li>
    <input type="submit" value="Ingresar"></input>
    </ul>
    
    </form>
    </div>
    </div>

);
}
