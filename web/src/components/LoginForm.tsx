import { useState } from "react";

interface Props {
  onLogin: (username: string) => void;
}

export function LoginForm({ onLogin }: Props) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setError("");

    try {
      const response = await fetch(
        `${import.meta.env.VITE_REST_ADDR}/v1/authenticate`,
        {
          method: "POST",
          body: JSON.stringify({
            username: username,
            password: password,
          }),
        }
      );
      if (!response.ok) {
        setError("Error: " + response.statusText);
        onLogin("");
        return;
      }

      const result = await response.text();

      localStorage.setItem("token", result); // Save token
      onLogin(username);
    } catch (err) {
      setError("Error logging in:");
      onLogin("");
    }
  };

  return (
    <div className="w-full max-w-md">
      {error && <p style={{ color: "red" }}>{error}</p>}
      <form onSubmit={handleLogin}>
        <div className="my-4">
          <label className="block mb-1" htmlFor="username">
            Username:
          </label>
          <input
            type="text"
            id="username"
            name="username"
            placeholder="Username"
            className="w-full p-2 border rounded"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="my-4">
          <label className="block mb-1" htmlFor="password">
            Password:
          </label>
          <input
            type="password"
            id="password"
            name="password"
            placeholder="Password"
            className="w-full p-2 border rounded"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button
          type="submit"
          className="w-full px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600"
        >
          Login
        </button>
      </form>
    </div>
  );
}
