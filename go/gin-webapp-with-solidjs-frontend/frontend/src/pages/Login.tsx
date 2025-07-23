import { Component, createSignal } from "solid-js";
import { A, useNavigate } from "@solidjs/router";
import { useAuth } from "../contexts/AuthContext";

const Login: Component = () => {
  const [username, setUsername] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [error, setError] = createSignal("");
  const [loading, setLoading] = createSignal(false);
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleSubmit = async (e: SubmitEvent) => {
    e.preventDefault();
    e.stopPropagation();
    setError("");
    setLoading(true);

    try {
      await login({ username: username(), password: password() });
      navigate("/dashboard");
    } catch (err: any) {
      setError(err.message || "Login failed. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 py-12 px-4">
      <div class="max-w-md w-full">
        <div class="text-center mb-8">
          <h2 class="text-3xl font-bold text-gray-900">
            Sign in to your account
          </h2>
          <p class="mt-2 text-gray-600">
            Don't have an account?{" "}
            <A
              href="/register"
              class="text-blue-600 hover:text-blue-500 font-medium"
            >
              Sign up
            </A>
          </p>
        </div>

        <div class="bg-white py-8 px-6 shadow rounded-lg">
          <form onSubmit={handleSubmit} class="space-y-6">
            {error() && (
              <div class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded">
                {error()}
              </div>
            )}

            <div>
              <label
                for="username"
                class="block text-sm font-medium text-gray-700 mb-1"
              >
                Username
              </label>
              <input
                id="username"
                type="text"
                required
                value={username()}
                onInput={(e) => setUsername(e.currentTarget.value)}
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Enter your username"
              />
            </div>

            <div>
              <label
                for="password"
                class="block text-sm font-medium text-gray-700 mb-1"
              >
                Password
              </label>
              <input
                id="password"
                type="password"
                required
                value={password()}
                onInput={(e) => setPassword(e.currentTarget.value)}
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Enter your password"
              />
            </div>

            <button
              type="submit"
              disabled={loading()}
              class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {loading() ? "Signing In..." : "Sign In"}
            </button>
          </form>

          <div class="mt-6 text-center">
            <A href="/" class="text-sm text-gray-600 hover:text-gray-500">
              ← Back to home
            </A>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
