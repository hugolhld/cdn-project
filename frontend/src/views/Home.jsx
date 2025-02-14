import FileList from "@/components/FileList"
import { Button } from "@/components/ui/button"
import { useAuth } from "@/context/AuthContext"
import { useSnackbar } from "notistack"
import { useCallback, useEffect, useRef, useState } from "react"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

const Home = () => {
    const { logout } = useAuth()
    const { enqueueSnackbar } = useSnackbar()
    const [files, setFiles] = useState([])
    const ref = useRef()

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
                    setFiles([...files, {
                        filename: file.name,
                        size: file.size,
                        type: file.type,
                        date: new Date().toISOString()
                    }])
                    enqueueSnackbar("File uploaded successfully", { variant: "success" })
                    console.log("success")

                    ref.current.value = null
                } else {
                    console.log(response)
                    enqueueSnackbar("File upload failed", { variant: "error" })
                    console.log("error")
                }
            })
            .catch(error => {
                console.log(error)
                enqueueSnackbar("File upload failed", { variant: "error" })
            })
    }

    const getFiles = useCallback(async () => {
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
                    enqueueSnackbar("Failed to fetch files", { variant: "error" })
                }
            })
            .then(({ files }) => {
                setFiles(files)
            })
    }, [enqueueSnackbar])

    useEffect(() => {
        getFiles()
    }, [getFiles])

    return (
        <div className="flex flex-col w-full items-center justify-center px-4">
            <div className="w-full flex justify-between items-center py-6">
                <h1 className="text-3xl font-semibold text-center">
                    Welcome to CDN-PROJECT
                </h1>
                <Button className="bg-red-500 hover:bg-red-600" onClick={logout}>
                    Logout
                </Button>
            </div>
            <div className="w-full flex flex-col justify-start gap-8">
                <div className="flex gap-2  mx-auto">
                    <FileList files={files} />
                    <div className="flex flex-col justify-start w-full gap-1.5">
                        <Label htmlFor="file" className="text-xl">Upload a file</Label>
                        <Input ref={ref} id="file" type="file" onChange={onUpload} />
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Home