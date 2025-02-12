import { useNavigate } from "react-router";
import { useAuth } from "../context/AuthContext";
import { useEffect } from "react";


const ProtectedRoute = ({ children }) => {
    const { isAuthenticated, logout } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        const checkAuth = async () => {
            const response = await fetch('http://localhost/api/check', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem("token")}`
                },
            })

            if (response.status === 401) {
                logout()
                navigate("/login")
            }
        }

        checkAuth()

        if (!isAuthenticated) {
            navigate("/login")
        }
    }, [isAuthenticated, navigate, logout])

    return children
};

export default ProtectedRoute;