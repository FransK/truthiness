import { useState } from "react";
import { ExperimentViewer } from "./features/ExperimentViewer";

export default function App() {
  const [username, setUsername] = useState<string>("");

  function onLogin(user: string) {
    setUsername(user);
  }

  return (
    <>
      <ExperimentViewer onLogin={onLogin} isLoggedIn={username !== ""} />
    </>
  );
}
