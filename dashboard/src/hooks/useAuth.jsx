import { createContext, useContext, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { useLocalStorage } from "./useLocalStorage";
import axios from "axios";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useLocalStorage("user", null);
  const navigate = useNavigate();

  const login = async (data) => {
    axios
      .post(
        "http://localhost:8002/v1/user/signin",
        JSON.stringify({
          email: data.email,
          password: data.password,
        }),
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((response) => {
        console.log(response);
        setUser(data);
        navigate("/dashboard/profile", { replace: true });
      });
  };

  const signup = async (data) => {
    fetch("https://probe.muhesh.studio/server/dashboard/v1/user/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: data.email,
        name: data.name,
        password: data.password,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        // setUser(data);
        // navigate("/dashboard/profile", { replace: true });
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
