import "./SearchableDropdown.css";

import { useEffect, useRef, useState } from "react";

interface MyDropdownProps {
  options: { [key: number]: string }[];
  label: string;
  id: string;
  selectedVal: string;
  handleChange: (val: string) => void;
}

export default function MyDropdown({
  options,
  label,
  id,
  selectedVal,
  handleChange,
}: MyDropdownProps) {
  const [query, setQuery] = useState("");
  const [isOpen, setIsOpen] = useState(false);

  const inputRef = useRef(null);

  useEffect(() => {
    document.addEventListener("click", toggleHandler);
    return () => document.removeEventListener("click", toggleHandler);
  }, []);

  const selectOption = (option: { [key: string]: string }) => {
    setQuery(() => "");
    handleChange(option[label]);
    setIsOpen((isOpen) => !isOpen);
  };

  function toggleHandler(e: MouseEvent) {
    setIsOpen(e && e.target === inputRef.current);
  }

  function toggle(e: React.MouseEvent<HTMLElement>) {
    setIsOpen(e && e.target === inputRef.current);
  }

  const getDisplayValue = () => {
    if (query) return query;
    if (selectedVal) return selectedVal;

    return "";
  };

  let filter = (options: { [key: string]: string }[]) => {
    return options.filter(
      (option) => option[label].toLowerCase().indexOf(query.toLowerCase()) > -1
    );
  };

  return (
    <div className={`dropdown ${isOpen ? "open" : ""}`}>
      <div className="control">
        <div className="selected-value">
          <input
            ref={inputRef}
            type="text"
            value={getDisplayValue()}
            name="searchTerm"
            onChange={(e) => {
              setQuery(e.target.value);
              handleChange("");
            }}
            onClick={toggle}
          />
        </div>
      </div>

      <div className={`options ${isOpen ? "open" : ""}`}>
        {filter(options).map((option, index) => {
          return (
            <div
              onClick={() => selectOption(option)}
              className={`option ${
                option[label] === selectedVal ? "selected" : ""
              }`}
              key={`${id}-${index}`}
            >
              {option[label]}
            </div>
          );
        })}
      </div>
    </div>
  );
}
