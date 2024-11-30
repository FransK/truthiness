import React from "react";
import { Upload } from "lucide-react";
import { parseCSV } from "../utils/csv";
import { db } from "../db/db";

export function FileUpload() {
  const [isUploading, setIsUploading] = React.useState(false);

  const handleFileUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (!file) return;

    try {
      setIsUploading(true);
      const experimentData = await parseCSV(file);
      await db.experiments.add(experimentData);
      event.target.value = "";
    } catch (error) {
      console.error("Error uploading file:", error);
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="w-full max-w-md">
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
            {isUploading ? "Uploading..." : "Click to upload CSV file"}
          </p>
        </div>
        <input
          type="file"
          className="hidden"
          accept=".csv"
          onChange={handleFileUpload}
          disabled={isUploading}
        />
      </label>
    </div>
  );
}
