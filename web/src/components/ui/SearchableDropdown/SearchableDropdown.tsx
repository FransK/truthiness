import "./SearchableDropdown.css";

import { useEffect, useRef, useState } from "react";

interface MyDropdownProps {
  options: string[];
  id: string;
  selectedVal: string;
  handleChange: (val: string) => void;
}

export default function MyDropdown({
  options,
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

  const selectOption = (option: string) => {
    setQuery(() => "");
    setIsOpen(false);
    handleChange(option);
  };

  function toggleHandler() {
    if (document.activeElement !== document.getElementById("searchBox")) {
      setIsOpen(false);
    } else {
      setIsOpen(true);
    }
  }

  function keyDownHandler() {
    setIsOpen(true);
  }

  const getDisplayValue = () => {
    if (query) return query;
    if (selectedVal) return selectedVal;

    return "";
  };

  let filter = (options: string[]) => {
    return options.filter(
      (option) => option.toLowerCase().indexOf(query.toLowerCase()) > -1
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
            placeholder="Select experiment..."
            id="searchBox"
            name="searchTerm"
            onChange={(e) => {
              setQuery(e.target.value);
              handleChange("");
            }}
            onKeyDown={keyDownHandler}
          />
        </div>
      </div>

      <div className={`options ${isOpen ? "open" : ""}`}>
        {filter(options).map((option, index) => {
          return (
            <div
              onClick={() => selectOption(option)}
              className={`option ${option === selectedVal ? "selected" : ""}`}
              key={`${id}-${index}`}
            >
              {option}
            </div>
          );
        })}
      </div>
    </div>
  );
}
