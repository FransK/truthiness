import React from "react";
import { Upload } from "lucide-react";

export function FileUpload() {
  const [isUploading, setIsUploading] = React.useState(false);
  const [file, setFile] = React.useState<File | null>(null);

  const handleSelectFile = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    setFile(file ? file : null);
  };

  async function handleUploadFile(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault(); //prevent redirects
    const formData = new FormData(e.currentTarget);

    setIsUploading(true);

    fetch(`${import.meta.env.VITE_REST_ADDR}/v1/upload`, {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        console.log(response.json());
      })
      .catch((error) => console.error("Error fetching experiments:", error))
      .finally(() => setIsUploading(false));
  }

  return (
    <div className="w-full max-w-md">
      <form onSubmit={handleUploadFile}>
        <label
          className={`flex flex-col items-center justify-center w-full h-32 border-2 border-dashed rounded-lg cursor-pointer
          ${
            isUploading
              ? "bg-gray-100 border-gray-300"
              : "border-blue-300 hover:bg-blue-50"
          }`}
        >
          <div className="flex flex-col items-center justify-center pt-5 pb-6">
            <Upload className="w-8 h-8 mb-2 text-blue-500" />
            <p className="mb-2 text-sm text-gray-500">
              {isUploading
                ? "Uploading..."
                : file
                ? file.name
                : "Click to upload CSV file"}
            </p>
          </div>
          <input
            type="file"
            name="file"
            id="file"
            className="hidden"
            accept=".csv"
            onChange={handleSelectFile}
          />
        </label>
        <div>
          <label className="flex w-full" htmlFor="experiment">
            Experiment name:
          </label>
          <input type="text" id="experiment" name="experiment" />
        </div>
        <div>
          <label className="flex w-full" htmlFor="date">
            Experiment date:
          </label>
          <input type="text" id="date" name="date" />
        </div>
        <div>
          <label className="flex w-full" htmlFor="experiment">
            Experiment location:
          </label>
          <input type="text" id="location" name="location" />
        </div>
        <div>
          <button type="submit">Upload</button>
        </div>
      </form>
    </div>
  );
}
