export interface User {
  id: string
  email: string
  discord_id?: string
  is_active: boolean
  plan?: 'free' | 'subscriber' | 'pro'
  subscription_tier: string
  subscription_status: string
  created_at: string
  updated_at: string
}

export interface Portfolio {
  id: string
  user_id: string
  name: string
  addresses: Address[]
  created_at: string
  updated_at: string
}

export interface Address {
  id: string
  portfolio_id: string
  address: string
  network: string
  label?: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface Balance {
  id: string
  address_id: string
  token_address?: string
  symbol: string
  name: string
  amount: string
  decimals: number
  price: string
  value: string
  updated_at: string
}

export interface Transaction {
  id: string
  portfolio_id: string
  tx_hash: string
  network: string
  token_address?: string
  amount: string
  block_number: number
  timestamp: string
  created_at: string
}

export interface Alert {
  id: string
  user_id: string
  type: string
  name: string
  conditions: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface PortfolioOverview {
  total_value: string
  daily_change: string
  weekly_change: string
  monthly_change: string
  asset_count: number
  addresses: Address[]
  balances: Balance[]
}

export interface NetworkStatus {
  ethereum: boolean
  polygon: boolean
  bsc: boolean
  arbitrum: boolean
}

export interface Subscription {
  subscription_tier: string
  subscription_status: string
}

export interface SubscriptionPlan {
  id: string
  name: string
  tier: string
  price: number
  currency: string
  interval: string
  features: string[]
  popular?: boolean
}

export interface SubscriptionUpdateRequest {
  tier: string
  status?: string
}

export interface ApiResponse<T> {
  data: T
  message?: string
  error?: string
} 