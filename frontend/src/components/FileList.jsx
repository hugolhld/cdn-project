// eslint-disable-next-line react/prop-types
const FileList = ({ files }) => {
  return (
    <div>
      <h3>Files list</h3>
      <ul>
        {files?.map((file, index) => (
          <li key={index}>{file.filename}</li>
        ))}
      </ul>
    </div>
  )
}

export default FileList