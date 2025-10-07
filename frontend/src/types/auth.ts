export interface User {
  id: string;
  created_at: string;
  updated_at: string;
  username: string;
  email: string;
  first_name: string;
  last_name: string;
  avatar: string;
  is_active: boolean;
  is_verified: boolean;
  last_login_at: string | null;
  roles: Role[];
}

export interface Role {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
  description: string;
  permissions: string[];
  is_active: boolean;
}

export interface Policy {
  id: string;
  created_at: string;
  updated_at: string;
  name: string;
  description: string;
  resource: string;
  action: string;
  effect: 'allow' | 'deny';
  conditions: Record<string, string>;
  is_active: boolean;
}

export interface Session {
  id: string;
  created_at: string;
  updated_at: string;
  user_id: string;
  token: string;
  expires_at: string;
  ip_address: string;
  user_agent: string;
  is_active: boolean;
  user: User;
}

export interface AuditLog {
  id: string;
  created_at: string;
  user_id: string | null;
  action: string;
  resource: string;
  resource_id: string | null;
  details: Record<string, any>;
  ip_address: string;
  user_agent: string;
  success: boolean;
  user?: User;
}

// Request/Response types
export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  user: User;
  access_token: string;
  refresh_token: string;
  expires_in: number;
}

export interface RefreshTokenRequest {
  refresh_token: string;
}

export interface CreateUserRequest {
  username: string;
  email: string;
  password: string;
  first_name?: string;
  last_name?: string;
  role_ids?: string[];
}

export interface UpdateUserRequest {
  username?: string;
  email?: string;
  first_name?: string;
  last_name?: string;
  avatar?: string;
  is_active?: boolean;
  role_ids?: string[];
}

export interface ChangePasswordRequest {
  current_password: string;
  new_password: string;
}

export interface CreateRoleRequest {
  name: string;
  description?: string;
  permissions: string[];
}

export interface UpdateRoleRequest {
  name?: string;
  description?: string;
  permissions?: string[];
  is_active?: boolean;
}

export interface CreatePolicyRequest {
  name: string;
  description?: string;
  resource: string;
  action: string;
  effect: 'allow' | 'deny';
  conditions?: Record<string, string>;
}

export interface UpdatePolicyRequest {
  name?: string;
  description?: string;
  resource?: string;
  action?: string;
  effect?: 'allow' | 'deny';
  conditions?: Record<string, string>;
  is_active?: boolean;
}

// API Response types
export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface PaginatedResponse<T> {
  items: T[];
  pagination: {
    current_page: number;
    total_pages: number;
    total_items: number;
    items_per_page: number;
  };
}

export interface ApiError {
  error: string;
  message?: string;
  details?: Record<string, any>;
}