import React, { useState } from "react";
import { ChevronRight, ChevronDown, Folder, File, PlusCircle, Upload, Trash } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

const FileTreeManager = () => {
  const [structure, setStructure] = useState([]);
  const [newFolderName, setNewFolderName] = useState("");
  const [selectedPath, setSelectedPath] = useState(null);
  const [fileUpload, setFileUpload] = useState(null);

  const addFolder = () => {
    if (!newFolderName) return;
    const newFolder = {
      id: `${selectedPath ? selectedPath.id + "/" : ""}${newFolderName}`,
      name: newFolderName,
      type: "folder",
      children: []
    };
    if (selectedPath) {
      selectedPath.children = [...(selectedPath.children || []), newFolder];
      setStructure([...structure]);
    } else {
      setStructure([...structure, newFolder]);
    }
    setNewFolderName("");
  };

  const uploadFile = () => {
    if (!fileUpload) return;
    const newFile = {
      id: `${selectedPath ? selectedPath.id + "/" : ""}${fileUpload.name}`,
      name: fileUpload.name,
      type: "file"
    };
    if (selectedPath) {
      selectedPath.children = [...(selectedPath.children || []), newFile];
      setStructure([...structure]);
    } else {
      setStructure([...structure, newFile]);
    }
    setFileUpload(null);
  };

  const toggleFolder = (item) => {
    item.open = !item.open;
    setStructure([...structure]);
  };

  const deleteItem = (item, parent = null) => {
    if (parent) {
      parent.children = parent.children.filter(i => i !== item);
      setStructure([...structure]);
    } else {
      setStructure(structure.filter(i => i !== item));
    }
  };

  const renderTree = (items, parent = null) => {
    return (
      <ul>
        {items.map((item) => (
          <li key={item.id} className="ml-4 flex items-center gap-2 cursor-pointer">
            {item.type === "folder" && (
              <span onClick={() => toggleFolder(item)}>
                {item.open ? <ChevronDown /> : <ChevronRight />}
              </span>
            )}
            {item.type === "folder" ? <Folder className="w-5 h-5" /> : <File className="w-5 h-5" />}
            <span onClick={() => setSelectedPath(item)}>{item.name}</span>
            <Button onClick={() => deleteItem(item, parent)} variant="ghost" className="ml-auto">
              <Trash className="w-4 h-4 text-red-500" />
            </Button>
            {item.open && item.children && renderTree(item.children, item)}
          </li>
        ))}
      </ul>
    );
  };

  return (
    <Card className="w-96 p-4">
      <CardContent>
        <div className="mb-4 flex gap-2">
          <Input placeholder="Nom du dossier" value={newFolderName} onChange={(e) => setNewFolderName(e.target.value)} />
          <Button onClick={addFolder}><PlusCircle className="w-5 h-5" /></Button>
        </div>
        <div className="mb-4 flex gap-2">
          <input type="file" onChange={(e) => setFileUpload(e.target.files[0])} />
          <Button onClick={uploadFile}><Upload className="w-5 h-5" /></Button>
        </div>
        <div className="space-y-2">
          {renderTree(structure)}
        </div>
      </CardContent>
    </Card>
  );
};

export default FileTreeManager;
