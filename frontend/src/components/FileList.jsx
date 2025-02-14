// eslint-disable-next-line react/prop-types
const FileList = ({ files }) => {
  return (
    <div>
      <h3 className="text-2xl font-bold">Files list</h3>
      <ul className="flex flex-col gap-2 py-4">
        {files?.map((file, index) => (
          <li key={index}>- {file.filename}</li>
        ))}
      </ul>
    </div>
  )
}

export default FileList