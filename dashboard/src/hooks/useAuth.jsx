import { createContext, useContext, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { useLocalStorage } from "./useLocalStorage";
import axios from "axios";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useLocalStorage("user", null);
  const navigate = useNavigate();

  const login = async (data) => {
    fetch(" https://probe.muhesh.studio/server/dashboard/v1/user/signin",{
    method: 'POST',
    body: JSON.stringify({
      email: data.email,
      password: data.password
    }), 
    headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json'
      }
    }
    ).then((response) => response.json())
    .then((data)=>{
      console.log(data);
      setUser(data.response.token);
      navigate("/apps", { replace: true });
    });
  };

  const signup = async (data) => {
    fetch(" https://probe.muhesh.studio/server/dashboard/v1/user/signup",{
        method: 'POST',
        body: JSON.stringify({
          email: data.email,
          password: data.password,
          name: data.name
        }), 
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json'
          }
        }
        ).then((response) => response.json())
        .then((data)=>{
      console.log(data);
      navigate("/login", { replace: true });
    });
  };

  const logout = () => {
    setUser(null);
    navigate("/", { replace: true });
  };

  const value = useMemo(
    () => ({
      user,
      login,
      logout,
      signup,
    }),
    [user]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
  return useContext(AuthContext);
};
