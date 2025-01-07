interface Message {
  username: string;
  content: string;
}

interface MessageListProps {
  messages: Message[];
}

export default function MessageList({ messages }: MessageListProps) {
  return (
    <div className="flex-grow overflow-y-auto p-4">
      {messages ? (messages.map((msg, index) => (
        <div key={index} className="mb-2">
          <strong>{msg.username}</strong>: {msg.content}
        </div>
      ))):null}
    </div>
  );
}
