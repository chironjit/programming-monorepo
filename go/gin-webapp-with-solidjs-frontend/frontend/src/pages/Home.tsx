import { Component } from "solid-js";
import { A, useNavigate } from "@solidjs/router";
import { useAuth } from "../contexts/AuthContext";

const Home: Component = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  const handleLoginClick = () => {
    if (isAuthenticated()) {
      navigate("/dashboard");
    } else {
      navigate("/login");
    }
  };

  const handleGetStartedClick = () => {
    if (isAuthenticated()) {
      navigate("/dashboard");
    } else {
      navigate("/register");
    }
  };

  return (
    <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex flex-col items-center justify-center min-h-screen text-center">
          <h1 class="text-5xl font-bold text-gray-900 mb-6">
            Demo web app with <span class="text-blue-600">SolidJS</span> and{" "}
            <span class="text-indigo-600">Gin</span>
          </h1>

          <p class="text-xl text-gray-600 mb-12 max-w-2xl">
            A simple demo web application built with SolidJS and the Gin Go web
            framework
          </p>

          <div class="flex space-x-4">
            <button
              onClick={handleGetStartedClick}
              class="bg-blue-600 text-white px-8 py-3 rounded-lg font-semibold hover:bg-blue-700 transition-colors"
            >
              Get Started
            </button>

            <button
              onClick={handleLoginClick}
              class="bg-white text-blue-600 px-8 py-3 rounded-lg font-semibold border border-blue-600 hover:bg-blue-50 transition-colors"
            >
              Login
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
