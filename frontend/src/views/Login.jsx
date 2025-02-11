import { LoginForm } from "@/components/login-form"
import { useAuth } from "@/context/AuthContext"
import { useNavigate } from "react-router"
import { useEffect } from "react"

const Login = () => {

    const { isAuthenticated } = useAuth()
    const navigate = useNavigate()

    useEffect(() => {
        if (isAuthenticated) {
            navigate('/')
        }
    }, [isAuthenticated, navigate])

    return (
        <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
            <div className="w-full max-w-sm">
                <LoginForm />
            </div>
        </div>
    )
}

export default Login