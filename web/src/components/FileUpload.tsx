import React, { useState } from "react";
import { Upload } from "lucide-react";

export function FileUpload() {
  const [isUploading, setIsUploading] = useState(false);
  const [file, setFile] = useState<File | null>(null);
  const [error, setError] = useState("");

  const handleSelectFile = (event: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = event.target.files?.[0];
    setFile(selectedFile || null);
  };

  const handleUploadFile = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");

    if (!file) {
      setError("No file selected");
      return;
    }

    const token = localStorage.getItem("token");
    const formData = new FormData(e.currentTarget);
    setIsUploading(true);

    try {
      const response = await fetch(
        `${import.meta.env.VITE_REST_ADDR}/v1/upload`,
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${token}`,
          },
          body: formData,
        }
      );
      if (!response.ok) {
        if (response.status == 401) {
          setError("Authorization failed. Please check your credentials.");
        } else if (response.status == 403) {
          setError("Could not upload. Insufficient permissions.");
        } else {
          setError(`Error: ${response.status} - ${response.statusText}`);
          throw new Error(error);
        }
        return;
      }

      const result = await response.json();
      console.log("Upload successful:", result);
    } catch (err) {
      console.error("Error uploading file:", err);
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="w-full max-w-md">
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleUploadFile}>
        {/* File Upload Input */}
        <label
          className={`flex flex-col items-center justify-center w-full h-32 border-2 border-dashed rounded-lg cursor-pointer ${
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

        {/* Additional Form Inputs */}
        <div className="my-4">
          <label className="block mb-1" htmlFor="experiment">
            Experiment name:
          </label>
          <input
            type="text"
            id="experiment"
            name="experiment"
            className="w-full p-2 border rounded"
            required
          />
        </div>
        <div className="my-4">
          <label className="block mb-1" htmlFor="date">
            Experiment date:
          </label>
          <input
            type="text"
            id="date"
            name="date"
            className="w-full p-2 border rounded"
            required
          />
        </div>
        <div className="my-4">
          <label className="block mb-1" htmlFor="location">
            Experiment location:
          </label>
          <input
            type="text"
            id="location"
            name="location"
            className="w-full p-2 border rounded"
            required
          />
        </div>

        {/* Submit Button */}
        <button
          type="submit"
          className="w-full px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600"
          disabled={isUploading}
        >
          {isUploading ? "Uploading..." : "Upload"}
        </button>
      </form>
    </div>
  );
}
