import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import axios from "axios";



export default function Home() {
  const [username, setUsername] = useState("");
  const router = useRouter();
  useEffect(() => {
    if(localStorage.getItem("username")){
      router.push("/chat");
  
    }
  });
  
  const handleLogin = async (ev: React.MouseEvent<HTMLButtonElement>) => {
    try {
      const userRes =  await axios.get("/api/get-users");
      const users = await userRes.data;
     
      if ((await userRes.data && users.includes(username.trim())) || username == ""){
        ev.preventDefault();
        ev.stopPropagation();

        throw Error("Username already exists")

      }
     else if (username.trim()) {
        localStorage.setItem("username", username);
        router.push("/chat");
      }
    } catch (error) {
      console.error("Failed to fetch users", error);
      return
    }

   
  };

  return (
    <div className="h-screen flex items-center justify-center">
      <div className="p-8 border rounded-md">
        <h1 className="text-2xl font-bold mb-4">Login</h1>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter your username"
          className="border p-2 w-full mb-4 rounded-md"
        />
        <button
          onClick={handleLogin}
          className="px-4 py-2 bg-blue-500 text-white rounded-md w-full"
        >
          Join Chat
        </button>
      </div>
    </div>
  );
}
