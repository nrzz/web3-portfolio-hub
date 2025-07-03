import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react'
import { User, Subscription } from '../types'

interface AuthContextType {
  user: User | null
  login: (email: string, password: string) => Promise<void>
  register: (name: string, email: string, password: string) => Promise<void>
  logout: () => void
  isLoading: boolean
  getSubscription: () => Promise<Subscription | null>
  updateSubscription: (tier: string, status?: string) => Promise<void>
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}

interface AuthProviderProps {
  children: ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    // Check if user is logged in on app start
    const token = localStorage.getItem('token')
    if (token) {
      // Validate token and get user info
      // For now, we'll just set a mock user
      setUser({
        id: '1',
        email: 'user@example.com',
        is_active: true,
        plan: 'free',
        subscription_tier: 'basic',
        subscription_status: 'active',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      })
    }
    setIsLoading(false)
  }, [])

  const login = async (email: string, password: string) => {
    try {
      // In a real app, you'd make an API call here
      const response = await fetch('/api/v1/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!response.ok) {
        throw new Error('Login failed')
      }

      const data = await response.json()
      localStorage.setItem('token', data.token)
      setUser(data.user)
    } catch (error) {
      console.error('Login error:', error)
      throw error
    }
  }

  const register = async (name: string, email: string, password: string) => {
    try {
      const response = await fetch('/api/v1/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, email, password }),
      })

      if (!response.ok) {
        throw new Error('Registration failed')
      }

      const data = await response.json()
      localStorage.setItem('token', data.token)
      setUser(data.user)
    } catch (error) {
      console.error('Registration error:', error)
      throw error
    }
  }

  const logout = () => {
    localStorage.removeItem('token')
    setUser(null)
  }

  const getSubscription = async (): Promise<Subscription | null> => {
    try {
      const token = localStorage.getItem('token')
      if (!token) return null

      const response = await fetch('/api/v1/user/subscription', {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Failed to fetch subscription')
      }

      const data = await response.json()
      return data
    } catch (error) {
      console.error('Get subscription error:', error)
      return null
    }
  }

  const updateSubscription = async (tier: string, status?: string): Promise<void> => {
    try {
      const token = localStorage.getItem('token')
      if (!token) throw new Error('Not authenticated')

      const response = await fetch('/api/v1/user/subscription', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({ tier, status }),
      })

      if (!response.ok) {
        throw new Error('Failed to update subscription')
      }

      const data = await response.json()
      
      // Update user state with new subscription info
      if (user) {
        setUser({
          ...user,
          subscription_tier: data.subscription_tier,
          subscription_status: data.subscription_status,
        })
      }
    } catch (error) {
      console.error('Update subscription error:', error)
      throw error
    }
  }

  const value = {
    user,
    login,
    register,
    logout,
    isLoading,
    getSubscription,
    updateSubscription,
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
} 