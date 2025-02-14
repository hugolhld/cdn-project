import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Link, useNavigate } from "react-router"
import React from "react"
import { useSnackbar } from "notistack"
import { useAuth } from "../context/AuthContext"

export function LoginForm({
  // eslint-disable-next-line react/prop-types
  className,
  ...props
}) {

  const navigate = useNavigate()
  const { enqueueSnackbar } = useSnackbar()
  const [data, setData] = React.useState({
    email: "",
    password: "",
  })

  const { login } = useAuth()

  const handleChange = (e) => {
    setData({ ...data, [e.target.id]: e.target.value })
  }

  const onSubmit = async (e) => {
    e.preventDefault()

    await fetch(`${import.meta.env.VITE_API_URL}/api/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data),
    })
      .then(response => {
        if (response.ok) {
          console.log("success")
          enqueueSnackbar("Login successful", { variant: "success" })
          return response.json()
        } else {
          console.log("error")
          enqueueSnackbar("Login failed", { variant: "error" })
        }
      })
      .then(data => {
        console.log(data.data.token)
        login(data.data.token)
        navigate("/")
      })
  }

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
          <CardDescription>
            Enter your email below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={(e) => onSubmit(e)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                  onChange={handleChange}
                />
              </div>
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="#"
                    className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <Input id="password" type="password" required onChange={handleChange} />
              </div>
              <Button type="submit" className="w-full">
                Login
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Don&apos;t have an account?{" "}
              <Link to={'/register'} className="underline underline-offset-4">
                Sign up
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
