import React from "react";
import { getCurrentUser } from "../services/auth.service";

const Profile: React.FC = () => {
  const currentUser = getCurrentUser();
  console.log(currentUser)

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
      <strong>Authorities:</strong>
      {/* <ul>
        {currentUser?.roles &&
          currentUser?.roles.map((role: string, index: number) => <li key={index}>{role}</li>)}
      </ul> */}
    </div>
  );
};

export default Profile;
