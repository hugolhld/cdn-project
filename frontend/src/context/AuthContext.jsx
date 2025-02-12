import { createContext, useContext, useReducer, useEffect } from "react";
import { useNavigate } from "react-router"; // ✅ Utiliser useNavigate

const AuthContext = createContext();

const authReducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("token", action.payload); // ✅ Stocker le token
      return { ...state, isAuthenticated: true, token: action.payload };
    case "LOGOUT":
      localStorage.removeItem("token");
      return { ...state, isAuthenticated: false, token: null };
    default:
      return state;
  }
};

// eslint-disable-next-line react/prop-types
export const AuthProvider = ({ children }) => {
  const navigate = useNavigate(); // ✅ Permet la redirection
  const [state, dispatch] = useReducer(authReducer, {
    isAuthenticated: !!localStorage.getItem("token"), // ✅ Vérifie si un token existe
    token: localStorage.getItem("token") || null,
  });

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      dispatch({ type: "LOGIN", payload: token }); // ✅ Charge le token au démarrage
    }
  }, []); // ✅ Exécute une seule fois au chargement

  const login = (token) => {
    dispatch({ type: "LOGIN", payload: token });
    navigate("/"); // ✅ Redirige après login
  };

  const logout = () => {
    dispatch({ type: "LOGOUT" });
    navigate("/login"); // ✅ Redirige après logout
  };

  return (
    <AuthContext.Provider value={{ ...state, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuth = () => useContext(AuthContext);
