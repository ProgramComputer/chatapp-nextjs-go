import { useEffect, useState } from "react";
import axios from "axios";
import ChatBox from "../components/ChatBox";
import MessageList from "../components/MessageList";
import UserList from "../components/UserList";

interface Message {
  username: string;
  content: string;
  timestamp: string;
}

export default function Chat() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [users, setUsers] = useState<string[]>([]);
  const username = typeof window !== "undefined" ? localStorage.getItem("username") : "";
  const insertUser = async () => {
    await axios.post("/api/insert-user", { username});

  };
  useEffect(() => {
    window.addEventListener('beforeunload',
      deleteUser)

    insertUser();
    const fetchMessages = async () => {
      try {
        const messageRes = await axios.get("/api/get-messages");
        setMessages(messageRes.data ? messageRes.data.reverse() : []);
      } catch (error) {
        console.error("Failed to fetch messages", error);
      }
    };

    const fetchUsers = async () => {
      try {
        const userRes = await axios.get("/api/get-users");
        setUsers(userRes.data);
      } catch (error) {
        console.error("Failed to fetch users", error);
      }
    };

    const interval = setInterval(() => {
      fetchMessages();
      fetchUsers();
    }, 1000); // Poll every 1 seconds to reduce database load

    return () => (clearInterval(interval));
  }, []);

  const handleSend = async (content: string) => {
    await axios.post("/api/send-message", { username, content });
  };
  const deleteUser = () =>  {
    localStorage.setItem("username","")

    return axios.post("/api/delete-user", { username}).then((response) => {
  });;
  };

  return (
    <div className="h-screen flex">
      <UserList users={users}/>
      <div className="flex flex-col flex-grow">
        <MessageList messages={messages} />
        <ChatBox onSend={handleSend} />
      </div>
    </div>
  );
}
