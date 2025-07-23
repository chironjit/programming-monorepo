const getApiBaseUrl = (): string => {
  // In production (Fly.io), use the production URL
  if (window.location.hostname === 'backend-purple-dream-570.fly.dev') {
    return 'https://backend-purple-dream-570.fly.dev';
  }
  // For local development
  return 'http://localhost:8080';
};

const API_BASE_URL = getApiBaseUrl();

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  password: string;
}

export interface User {
  id: string;
  username: string;
  created_at: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  user: User;
}

export interface TokenPair {
  access_token: string;
  refresh_token: string;
}

export interface RefreshRequest {
  refresh_token: string;
}

export interface SystemMetrics {
  go_version: string;
  num_goroutines: number;
  num_cpu: number;
  memory_alloc_mb: number;
  memory_total_mb: number;
  memory_sys_mb: number;
  gc_runs: number;
  uptime: string;
  goos: string;
  goarch: string;
}

export interface LiveDataResponse {
  server_time: string;
  counter: number;
  last_updated: string;
  system_metrics: SystemMetrics;
}

class ApiService {
  private getAuthHeaders(): HeadersInit {
    const accessToken = localStorage.getItem('access_token');
    return {
      'Content-Type': 'application/json',
      ...(accessToken && { Authorization: `Bearer ${accessToken}` }),
    };
  }

  private async refreshTokenIfNeeded(): Promise<boolean> {
    const refreshToken = localStorage.getItem('refresh_token');
    if (!refreshToken) return false;

    try {
      const response = await fetch(`${API_BASE_URL}/auth/refresh`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ refresh_token: refreshToken }),
      });

      if (response.ok) {
        const tokenPair: TokenPair = await response.json();
        localStorage.setItem('access_token', tokenPair.access_token);
        localStorage.setItem('refresh_token', tokenPair.refresh_token);
        return true;
      } else {
        // Refresh token is invalid, clear storage
        localStorage.removeItem('access_token');
        localStorage.removeItem('refresh_token');
        return false;
      }
    } catch (error) {
      console.error('Token refresh failed:', error);
      return false;
    }
  }

  private async makeAuthenticatedRequest(url: string, options: RequestInit = {}): Promise<Response> {
    // First attempt with current token
    let response = await fetch(url, {
      ...options,
      headers: {
        ...this.getAuthHeaders(),
        ...options.headers,
      },
    });

    // If unauthorized, try to refresh token
    if (response.status === 401) {
      const refreshed = await this.refreshTokenIfNeeded();
      if (refreshed) {
        // Retry with new token
        response = await fetch(url, {
          ...options,
          headers: {
            ...this.getAuthHeaders(),
            ...options.headers,
          },
        });
      }
    }

    return response;
  }

  async login(credentials: LoginRequest): Promise<AuthResponse> {
    // Create basic auth header to match backend expectations
    const basicAuth = btoa(`${credentials.username}:${credentials.password}`);
    
    const response = await fetch(`${API_BASE_URL}/login`, {
      method: 'POST',
      headers: {
        'Authorization': `Basic ${basicAuth}`,
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Invalid credentials');
    }

    const data = await response.json();
    
    // Transform backend response to match expected AuthResponse format
    return {
      access_token: data.token,
      refresh_token: data.token, // Backend only provides one token
      user: {
        id: 1,
        username: credentials.username,
        email: `${credentials.username}@demo.com`,
      }
    };
  }

  async register(userData: RegisterRequest): Promise<AuthResponse> {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: this.getAuthHeaders(),
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Registration failed');
    }

    return response.json();
  }

  async logout(): Promise<void> {
    const refreshToken = localStorage.getItem('refresh_token');
    if (refreshToken) {
      try {
        await fetch(`${API_BASE_URL}/auth/logout`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ refresh_token: refreshToken }),
        });
      } catch (error) {
        // Even if logout API fails, we still clear local storage
        console.error('Logout API error:', error);
      }
    }
    
    // Clear tokens regardless of API response
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
  }

  async getCurrentUser(): Promise<User> {
    const response = await this.makeAuthenticatedRequest(`${API_BASE_URL}/users/me`);

    if (!response.ok) {
      throw new Error('Failed to fetch user data');
    }

    return response.json();
  }

  async getLiveData(): Promise<LiveDataResponse> {
    const response = await this.makeAuthenticatedRequest(`${API_BASE_URL}/live-data`);

    if (!response.ok) {
      throw new Error('Failed to fetch live data');
    }

    return response.json();
  }
}

export const apiService = new ApiService();