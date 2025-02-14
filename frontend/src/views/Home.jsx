import FileList from "@/components/FileList"
import { Button } from "@/components/ui/button"
import { useAuth } from "@/context/AuthContext"
import { useSnackbar } from "notistack"
import { useEffect, useState } from "react"

const Home = () => {
    const { logout } = useAuth()
    const { enqueueSnackbar } = useSnackbar()
    const [files, setFiles] = useState([])

    const onUpload = async (e) => {
        const file = e.target.files[0]
        const formData = new FormData()
        formData.append('file', file)
        await fetch(`${import.meta.env.VITE_API_URL}/api/upload`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: formData
        })
        .then(response => {
            if (response.ok) {
                setFiles([...files, file])
                enqueueSnackbar("File uploaded successfully", { variant: "success" })
                console.log("success")
            } else {
                enqueueSnackbar("File upload failed", { variant: "error" })
                console.log("error")
            }
        })
    }

    const getFiles = async () => {
        await fetch(`${import.meta.env.VITE_API_URL}/api/files`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        })
        .then(response => {
            if (response.ok) {
                return response.json()
            } else {
                console.log("error")
            }
        })
        .then(({files}) => {
            console.log(files)
            setFiles(files)
        })
    }

    useEffect(() => {
        getFiles()
    }, [])

    return (
        <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
            <div className="w-full max-w-md flex flex-col justify-center gap-8">
                <h1 className="text-3xl font-semibold text-center">
                    Welcome to CDN-PROJECT
                </h1>
                <div className="flex flex-col gap-2 max-w-xs mx-auto">
                    <FileList files={files} />
                    <div className="py-4">
                        <input type="file" onChange={onUpload} />
                    </div>
                    <Button className="bg-red-500 hover:bg-red-600" onClick={logout}>
                        Logout
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default Home