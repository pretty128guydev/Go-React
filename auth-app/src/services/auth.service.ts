import axios from "axios";
import { jwtDecode } from "jwt-decode";

const API_URL = "http://localhost:8080/api/auth/";

interface DecodedToken {
    username: string;
    email: string;
    password: string;
    id: string;
  }

export const register = (username: string, email: string, password: string) => {
  return axios.post(API_URL + "signup", {
    username,
    email,
    password,
  });
};

export const login = (username: string, password: string) => {
  return axios
    .post(API_URL + "signin", {
      username,
      password,
    })
    .then((response) => {
      if (response.data.token) {
        localStorage.setItem("user", JSON.stringify(response.data));
      }
      console.log(response.data)
      return response.data;
    });
};

export const logout = () => {
  localStorage.removeItem("user");
};

export const getCurrentUser = () => {
    const userStr = localStorage.getItem("user");
    if (userStr) {
      const user = JSON.parse(userStr);
      const token = user.token;
  
      if (token) {
        try {
          const decodedToken = jwtDecode<DecodedToken>(token);
          console.log(decodedToken)
          return decodedToken;
        } catch (e) {
          console.error("Failed to decode token:", e);
          return null;
        }
      }
    }
    return null;
  };
