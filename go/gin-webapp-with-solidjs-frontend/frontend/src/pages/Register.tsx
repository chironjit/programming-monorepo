import { Component } from "solid-js";
import { A } from "@solidjs/router";

const Register: Component = () => {

  return (
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 py-12 px-4">
      <div class="max-w-md w-full">
        <div class="text-center mb-8">
          <h2 class="text-3xl font-bold text-gray-900">Demo App</h2>
          <p class="mt-2 text-gray-600">
            This is a demo application with a single user account.
          </p>
        </div>

        <div class="bg-white py-8 px-6 shadow rounded-lg text-center">
          <div class="mb-6">
            <h3 class="text-lg font-semibold text-gray-800 mb-4">Demo Credentials</h3>
            <div class="bg-gray-50 p-4 rounded-lg">
              <p class="text-sm text-gray-600 mb-2">Username: <strong class="text-gray-800">admin</strong></p>
              <p class="text-sm text-gray-600">Password: <strong class="text-gray-800">secret</strong></p>
            </div>
          </div>
          
          <div class="space-y-3">
            <A
              href="/login"
              class="block w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors"
            >
              Go to Login
            </A>
            
            <A
              href="/"
              class="block w-full text-gray-600 hover:text-gray-500 transition-colors"
            >
              ← Back to Home
            </A>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
