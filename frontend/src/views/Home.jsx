import { Button } from "@/components/ui/button"
import { useAuth } from "@/context/AuthContext"

const Home = () => {
    const { logout } = useAuth()

    return (
        <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
            <div className="w-full max-w-md flex flex-col justify-center gap-8">
                <h1 className="text-3xl font-semibold text-center">
                    Welcome to CDN-PROJECT
                </h1>
                <div className="flex flex-col gap-2 max-w-xs mx-auto">
                    <Button>
                        Upload your file
                    </Button>
                    <Button className="bg-red-500 hover:bg-red-600" onClick={logout}>
                        Logout
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default Home