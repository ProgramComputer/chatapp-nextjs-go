interface UserListProps {
  users: string[];
}

export default function UserList({ users }: UserListProps) {
  return (
    <div className="w-1/4 border-r p-4">
      <h2 className="font-bold mb-4">Users</h2>
      {users ? (users.map((user, index) => (
        <div key={index} className="mb-2">
          {user}
        </div>
      ))):null}
    </div>
  );
}
