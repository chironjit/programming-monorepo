package main

import (
	"fmt"
)

// CSS constant containing all dark mode styles
const CSS = `
<style>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}
body {
	background-color: #1a1a1a;
	color: #e0e0e0;
	font-family: Arial, sans-serif;
	line-height: 1.6;
	min-height: 100vh;
	display: flex;
	align-items: center;
	justify-content: center;
}
.container {
	max-width: 400px;
	padding: 2rem;
	background-color: #2d2d2d;
	border-radius: 8px;
	box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
	text-align: center;
}
h1 { margin-bottom: 1rem; color: #ffffff; }
h3 { color: #ffffff; margin-bottom: 0.5rem; }
p { margin-bottom: 1rem; }
.btn {
	display: inline-block;
	padding: 0.75rem 1.5rem;
	background-color: #4a9eff;
	color: white;
	text-decoration: none;
	border-radius: 4px;
	margin: 0.25rem;
	border: none;
	cursor: pointer;
	transition: background-color 0.3s;
}
.btn:hover { background-color: #3a8eef; }
.btn.error { background-color: #ff4757; }
.btn.error:hover { background-color: #ff3747; }
.btn.success { background-color: #2ed573; }
.btn.success:hover { background-color: #1eb863; }
.form { margin-bottom: 1.5rem; }
input {
	width: 100%;
	padding: 0.75rem;
	margin-bottom: 1rem;
	background-color: #3d3d3d;
	border: 1px solid #555;
	border-radius: 4px;
	color: #e0e0e0;
}
input::placeholder { color: #999; }
a { color: #4a9eff; text-decoration: none; }
a:hover { text-decoration: underline; }
.info-box {
	background-color: #3d3d3d;
	padding: 1rem;
	border-radius: 4px;
	margin: 1rem 0;
	border-left: 4px solid #4a9eff;
}
.error-box {
	background-color: #4d2d2d;
	padding: 1.5rem;
	border-radius: 4px;
	margin-bottom: 1.5rem;
	border-left: 4px solid #ff4757;
}
</style>
`

// Homepage function
func homepage() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Homepage</title>
	%s
</head>
<body>
	<div class="container">
		<h1>Welcome</h1>
		<p>This is the homepage of our application.</p>
		<a href="/login-page" class="btn">Login</a>
		<a href="/register-page" class="btn success">Register</a>
		<a href="/protected-page" class="btn">Protected Page</a>
	</div>
</body>
</html>`, CSS)
}

// Login function
func loginPage() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Login</title>
	%s
</head>
<body>
	<div class="container">
		<h1>Login</h1>
		<form class="form" method="POST" action="/login">
			<input type="text" name="username" placeholder="Username" required>
			<input type="password" name="password" placeholder="Password" required>
			<button type="submit" class="btn">Login</button>
		</form>
		<a href="/register-page">Don't have an account? Register</a><br><br>
		<a href="/">Back to Homepage</a>
	</div>
</body>
</html>`, CSS)
}

// Register function
func registerPage() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Register</title>
	%s
</head>
<body>
	<div class="container">
		<h1>Register</h1>
		<form class="form" method="POST" action="/register">
			<input type="text" name="username" placeholder="Username" required>
			<input type="password" name="password" placeholder="Password" required>
			<button type="submit" class="btn success">Register</button>
		</form>
		<a href="/login-page">Already have an account? Login</a><br><br>
		<a href="/">Back to Homepage</a>
	</div>
</body>
</html>`, CSS)
}

// Protected function
func protectedPage(csrfToken string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Protected Content</title>
	%s
</head>
<body>
	<div class="container">
		<h1>Protected Content</h1>
		<p>This page contains sensitive information that requires authentication.</p>
		<div class="info-box">
			<h3>Confidential Data</h3>
			<p>User dashboard with personal information and settings.</p>
		</div>
		<form method="POST" action="/logout" style="display: inline;">
			<input type="hidden" name="csrf_token" value="%s">
			<button type="submit" class="btn error">Logout</button>
		</form>
		<a href="/" class="btn">Back to Homepage</a>
	</div>
</body>
</html>`, CSS, csrfToken)
}

// Denied function
func deniedPage() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Access Denied</title>
	%s
</head>
<body>
	<div class="container">
		<div class="error-box">
			<h1>⚠️ Access Denied</h1>
			<p>You are not authenticated. Please log in to access this page.</p>
		</div>
		<a href="/login" class="btn">Go to Login</a>
		<a href="/" class="btn">Back to Homepage</a>
	</div>
</body>
</html>`, CSS)
}
