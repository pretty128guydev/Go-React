import React, { useState } from "react";
import { findUser, getCurrentUser } from "../services/auth.service";

const Profile: React.FC = () => {
  const currentUser = getCurrentUser();
  const [users, setUsers] = useState<any[]>([]);

  const onfindUser = () => {
    findUser()
      .then((response) => {
        setUsers(response);
      })
      .catch((error) => {
        console.error("Error finding users:", error);
      });
  };

  return (
    <div className="container">
      <header className="jumbotron">
        <h3>
          <strong>{currentUser?.username}</strong> Profile
        </h3>
      </header>
      <p>
        <strong>Email:</strong> {currentUser?.email}
      </p>
      <button className="btn-primary" onClick={onfindUser}>
        Find registerd users
      </button>
      <ul>
        {users.length > 0 ? (
          users.map((user, index) => (
            <>
              <li key={index}>Username: {user.username}</li>
              <li key={index}>Email: {user.email}</li>
            </>
          ))
        ) : (
          <p>No users found</p>
        )}
      </ul>
    </div>
  );
};

export default Profile;
