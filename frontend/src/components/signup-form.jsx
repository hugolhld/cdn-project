import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import React from "react"
import { Link, useNavigate } from "react-router"
import { useSnackbar } from "notistack"

// eslint-disable-next-line react/prop-types
export function SignupForm({ className, ...props }) {

    const { enqueueSnackbar } = useSnackbar()
    const navigate = useNavigate()

    const [data, setData] = React.useState({
        name: "",
        email: "",
        password: "",
        role: "Paris",
    })

    const handleChange = (e) => {
        setData({ ...data, [e.target.id]: e.target.value })
    }

    const onSubmit = async (e) => {
        e.preventDefault()

        await fetch(`${import.meta.env.VITE_API_URL}/api/member`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
            .then(response => {
                if (response.ok) {
                    console.log("success")
                    enqueueSnackbar("Account created successfully", { variant: "success" })
                    navigate("/login")
                } else {
                    console.log("error")
                    enqueueSnackbar("Account creation failed", { variant: "error" })
                }
            })
    }



    return (
        <div className={cn("flex flex-col gap-6", className)} {...props}>
            <Card>
                <CardHeader>
                    <CardTitle className="text-2xl">Sign Up</CardTitle>
                    <CardDescription>Create a new account</CardDescription>
                </CardHeader>
                <CardContent>
                    <form onSubmit={(e) => onSubmit(e)}>
                        <div className="flex flex-col gap-6">
                            <div className="grid gap-2">
                                <Label htmlFor="name">Name</Label>
                                <Input id="name" type="text" placeholder="John Doe" required onChange={handleChange} />
                            </div>
                            <div className="grid gap-2">
                                <Label htmlFor="email">Email</Label>
                                <Input id="email" type="email" placeholder="m@example.com" required onChange={handleChange} />
                            </div>
                            <div className="grid gap-2">
                                <Label htmlFor="password">Password</Label>
                                <Input id="password" type="password" required onChange={handleChange} />
                            </div>
                            <Button type="submit" className="w-full">
                                Sign Up
                            </Button>
                        </div>
                        <div className="mt-4 text-center text-sm">
                            Already have an account?{" "}
                            <Link to="/login" className="underline underline-offset-4">
                                Log in
                            </Link>
                        </div>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}

