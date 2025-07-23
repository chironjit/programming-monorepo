import { Component } from "solid-js";
import { useNavigate } from "@solidjs/router";
import { useAuth } from "../contexts/AuthContext";

const Dashboard: Component = () => {
  const navigate = useNavigate();
  const { user, logout } = useAuth();

  const handleLogout = () => {
    logout();
    navigate("/");
  };

  return (
    <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Header */}
      <div class="bg-white shadow">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="flex justify-between items-center py-6">
            <div>
              <h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>
              <p class="text-gray-600">Welcome back, {user()?.username || 'User'}!</p>
            </div>
            <button
              onClick={handleLogout}
              class="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors"
            >
              Logout
            </button>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        {/* Welcome Card */}
        <div class="bg-white overflow-hidden shadow rounded-lg">
          <div class="p-5">
            <div class="text-center">
              <h2 class="text-2xl font-bold text-gray-900 mb-4">🎉 Login Successful!</h2>
              <p class="text-gray-600 mb-6">
                You have successfully logged in and your JWT token is valid.
              </p>
              <div class="bg-green-50 border border-green-200 rounded-lg p-4">
                <p class="text-green-800 font-semibold">
                  Authentication Status: ✅ Valid
                </p>
                <p class="text-green-600 text-sm">
                  JWT token stored and authenticated successfully
                </p>
              </div>
            </div>
          </div>
        </div>

        {/* Profile Card */}
        <div class="bg-white overflow-hidden shadow rounded-lg mt-6">
          <div class="p-5">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-10 h-10 bg-blue-600 rounded-full flex items-center justify-center">
                  <span class="text-white font-medium text-lg">
                    {user()?.username?.charAt(0).toUpperCase() || 'U'}
                  </span>
                </div>
              </div>
              <div class="ml-5">
                <h3 class="text-lg leading-6 font-medium text-gray-900">
                  Profile
                </h3>
                <p class="text-sm text-gray-500">Your account information</p>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-5 py-3">
            <div class="space-y-1">
              <p class="text-sm text-gray-600">
                <span class="font-medium">Username:</span> {user()?.username || 'Unknown'}
              </p>
              <p class="text-sm text-gray-600">
                <span class="font-medium">Demo Account:</span> Yes
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
