import React, { useState, useEffect } from 'react'
import { motion } from 'framer-motion'
import { 
  Wallet, 
  TrendingUp, 
  TrendingDown, 
  DollarSign, 
  BarChart3, 
  Bell, 
  RefreshCw, 
  Plus,
  Eye,
  EyeOff,
  AlertCircle,
  CheckCircle,
  Clock,
  Sparkles,
  Crown,
  Coins,
  Activity,
  Target,
  Zap,
  ArrowRight,
  ArrowUpRight,
  ArrowDownRight
} from 'lucide-react'
import { 
  LineChart, 
  Line, 
  AreaChart, 
  Area, 
  PieChart, 
  Pie, 
  Cell, 
  ResponsiveContainer, 
  XAxis, 
  YAxis, 
  CartesianGrid, 
  Tooltip, 
  Legend,
  BarChart,
  Bar,
  ComposedChart
} from 'recharts'
import { useAuth } from '../contexts/AuthContext'
import { Link } from 'react-router-dom'

interface PortfolioSummary {
  total_value: string
  total_change_24h: string
  total_change_7d: string
  total_change_30d: string
  asset_count: number
  network_count: number
  top_assets: Array<{
    symbol: string
    name: string
    amount: string
    value: string
    change_24h: string
    network: string
  }>
  network_allocation: Record<string, string>
}

// Advanced chart data
const portfolioHistory = [
  { date: 'Jan', value: 100000, volume: 50000, change: 5.2 },
  { date: 'Feb', value: 95000, volume: 45000, change: -2.1 },
  { date: 'Mar', value: 110000, volume: 60000, change: 8.7 },
  { date: 'Apr', value: 105000, volume: 55000, change: -1.2 },
  { date: 'May', value: 120000, volume: 70000, change: 12.3 },
  { date: 'Jun', value: 125432, volume: 75000, change: 4.8 },
]

const networkData = [
  { name: 'Ethereum', value: 45, color: '#3B82F6', change: 2.1 },
  { name: 'Polygon', value: 25, color: '#8B5CF6', change: 1.8 },
  { name: 'BSC', value: 20, color: '#F59E0B', change: -0.5 },
  { name: 'Arbitrum', value: 10, color: '#10B981', change: 3.2 },
]

const assetPerformance = [
  { symbol: 'ETH', performance: 12.5, volume: 45000, price: 3200 },
  { symbol: 'USDC', performance: 0.1, volume: 25000, price: 1.00 },
  { symbol: 'MATIC', performance: 8.7, volume: 15000, price: 0.85 },
  { symbol: 'BNB', performance: 5.2, volume: 20000, price: 580 },
  { symbol: 'LINK', performance: 15.3, volume: 12000, price: 18.50 },
]

const recentTransactions = [
  { id: 1, type: 'buy', asset: 'ETH', amount: '0.5', value: '$1,250', time: '2 min ago', status: 'completed' },
  { id: 2, type: 'sell', asset: 'USDC', amount: '500', value: '$500', time: '15 min ago', status: 'completed' },
  { id: 3, type: 'buy', asset: 'MATIC', amount: '1000', value: '$800', time: '1 hour ago', status: 'pending' },
  { id: 4, type: 'transfer', asset: 'ETH', amount: '0.1', value: '$250', time: '2 hours ago', status: 'completed' },
]

const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-gray-900 text-white p-3 rounded-lg border border-gray-700">
        <p className="font-medium">{label}</p>
        {payload.map((entry: any, index: number) => (
          <p key={index} style={{ color: entry.color }}>
            {entry.name}: ${entry.value?.toLocaleString()}
          </p>
        ))}
      </div>
    )
  }
  return null
}

export const Dashboard: React.FC = () => {
  const { user } = useAuth()
  const [summary, setSummary] = useState<PortfolioSummary | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [selectedTimeframe, setSelectedTimeframe] = useState('24h')
  const [showValues, setShowValues] = useState(true)

  const isFreeUser = user?.plan === 'free'

  useEffect(() => {
    fetchDashboardData()
  }, [])

  const fetchDashboardData = async () => {
    try {
      setLoading(true)
      
      const mockSummary: PortfolioSummary = {
        total_value: "$125,432.50",
        total_change_24h: "+2.45%",
        total_change_7d: "+8.32%",
        total_change_30d: "+15.67%",
        asset_count: 12,
        network_count: 4,
        top_assets: [
          {
            symbol: "ETH",
            name: "Ethereum",
            amount: "2.5",
            value: "$4,875.00",
            change_24h: "+1.2%",
            network: "ethereum"
          },
          {
            symbol: "USDC",
            name: "USD Coin",
            amount: "10,000",
            value: "$10,000.00",
            change_24h: "0.0%",
            network: "ethereum"
          },
          {
            symbol: "MATIC",
            name: "Polygon",
            amount: "5,000",
            value: "$3,250.00",
            change_24h: "+3.1%",
            network: "polygon"
          },
          {
            symbol: "BNB",
            name: "Binance Coin",
            amount: "15",
            value: "$2,850.00",
            change_24h: "+0.8%",
            network: "bsc"
          }
        ],
        network_allocation: {
          "ethereum": "45%",
          "polygon": "25%",
          "bsc": "20%",
          "arbitrum": "10%"
        }
      }
      
      setSummary(mockSummary)
    } catch (err) {
      setError('Failed to load dashboard data')
      console.error('Dashboard error:', err)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 flex items-center justify-center">
        <motion.div
          animate={{ rotate: 360 }}
          transition={{ duration: 1, repeat: Infinity, ease: "linear" }}
          className="w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full"
        />
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 flex items-center justify-center">
        <motion.div
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          className="bg-red-50 border border-red-200 rounded-xl p-6 max-w-md"
        >
          <div className="flex items-center">
            <AlertCircle className="w-6 h-6 text-red-600 mr-3" />
            <div>
              <h3 className="text-lg font-medium text-red-800">Error</h3>
              <div className="mt-1 text-sm text-red-700">{error}</div>
            </div>
          </div>
        </motion.div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900">
      {/* Header */}
      <div className="relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-purple-600/10"></div>
        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-8 pb-12">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            className="flex flex-col lg:flex-row lg:justify-between lg:items-center gap-4"
          >
            <div>
              <motion.h1 
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                className="text-4xl font-bold bg-gradient-to-r from-gray-900 to-gray-600 dark:from-white dark:to-gray-300 bg-clip-text text-transparent"
              >
                Advanced Portfolio Dashboard
              </motion.h1>
              <motion.p
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.1 }}
                className="text-gray-600 dark:text-gray-300 mt-2"
              >
                Real-time analytics and performance tracking
              </motion.p>
            </div>
            
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.2 }}
              className="flex items-center space-x-4"
            >
              <button
                onClick={() => setShowValues(!showValues)}
                className="p-2 bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-200"
              >
                {showValues ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
              <button
                onClick={fetchDashboardData}
                className="p-2 bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-200"
              >
                <RefreshCw className="w-5 h-5" />
              </button>
            </motion.div>
          </motion.div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-12">
        {/* Portfolio Overview Cards */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8"
        >
          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl flex items-center justify-center">
                <Wallet className="w-6 h-6 text-white" />
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">Total Value</span>
            </div>
            <div className="space-y-2">
              <div className="text-2xl font-bold text-gray-900 dark:text-white">
                {showValues ? summary?.total_value : '••••••'}
              </div>
              <div className="flex items-center text-sm">
                <TrendingUp className="w-4 h-4 text-green-500 mr-1" />
                <span className="text-green-600 dark:text-green-400">
                  {summary?.total_change_24h}
                </span>
                <span className="text-gray-500 dark:text-gray-400 ml-1">24h</span>
              </div>
            </div>
          </div>

          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-gradient-to-r from-green-500 to-green-600 rounded-xl flex items-center justify-center">
                <TrendingUp className="w-6 h-6 text-white" />
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">7D Change</span>
            </div>
            <div className="space-y-2">
              <div className="text-2xl font-bold text-gray-900 dark:text-white">
                {summary?.total_change_7d}
              </div>
              <div className="text-sm text-gray-500 dark:text-gray-400">
                Weekly performance
              </div>
            </div>
          </div>

          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl flex items-center justify-center">
                <Coins className="w-6 h-6 text-white" />
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">Assets</span>
            </div>
            <div className="space-y-2">
              <div className="text-2xl font-bold text-gray-900 dark:text-white">
                {summary?.asset_count}
              </div>
              <div className="text-sm text-gray-500 dark:text-gray-400">
                Tracked assets
              </div>
            </div>
          </div>

          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300">
            <div className="flex items-center justify-between mb-4">
              <div className="w-12 h-12 bg-gradient-to-r from-orange-500 to-orange-600 rounded-xl flex items-center justify-center">
                <Activity className="w-6 h-6 text-white" />
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">Networks</span>
            </div>
            <div className="space-y-2">
              <div className="text-2xl font-bold text-gray-900 dark:text-white">
                {summary?.network_count}
              </div>
              <div className="text-sm text-gray-500 dark:text-gray-400">
                Connected chains
              </div>
            </div>
          </div>
        </motion.div>

        {/* Advanced Charts Section */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.3 }}
          className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8"
        >
          {/* Portfolio Performance Chart */}
          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg">
            <h3 className="text-xl font-bold text-gray-900 dark:text-white mb-6">Portfolio Performance</h3>
            <ResponsiveContainer width="100%" height={300}>
              <ComposedChart data={portfolioHistory}>
                <defs>
                  <linearGradient id="colorValue" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#3B82F6" stopOpacity={0.3}/>
                    <stop offset="95%" stopColor="#3B82F6" stopOpacity={0}/>
                  </linearGradient>
                </defs>
                <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
                <XAxis dataKey="date" stroke="#6B7280" />
                <YAxis stroke="#6B7280" />
                <Tooltip content={<CustomTooltip />} />
                <Legend />
                <Area 
                  type="monotone" 
                  dataKey="value" 
                  stroke="#3B82F6" 
                  strokeWidth={3}
                  fill="url(#colorValue)"
                  name="Portfolio Value"
                />
                <Bar dataKey="volume" fill="#10B981" opacity={0.6} name="Volume" />
              </ComposedChart>
            </ResponsiveContainer>
          </div>

          {/* Network Allocation Chart */}
          <div className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg">
            <h3 className="text-xl font-bold text-gray-900 dark:text-white mb-6">Network Allocation</h3>
            <ResponsiveContainer width="100%" height={300}>
              <PieChart>
                <Pie
                  data={networkData}
                  cx="50%"
                  cy="50%"
                  innerRadius={60}
                  outerRadius={100}
                  paddingAngle={5}
                  dataKey="value"
                >
                  {networkData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip 
                  contentStyle={{ 
                    backgroundColor: '#1F2937', 
                    border: 'none', 
                    borderRadius: '8px',
                    color: '#F9FAFB'
                  }}
                />
              </PieChart>
            </ResponsiveContainer>
            <div className="mt-4 space-y-2">
              {networkData.map((network, index) => (
                <div key={network.name} className="flex items-center justify-between">
                  <div className="flex items-center">
                    <div 
                      className="w-3 h-3 rounded-full mr-2"
                      style={{ backgroundColor: network.color }}
                    />
                    <span className="text-sm text-gray-600 dark:text-gray-300">{network.name}</span>
                  </div>
                  <div className="flex items-center space-x-2">
                    <span className="text-sm font-medium text-gray-900 dark:text-white">{network.value}%</span>
                    <span className={`text-xs ${network.change >= 0 ? 'text-green-600' : 'text-red-600'}`}>
                      {network.change >= 0 ? '+' : ''}{network.change}%
                    </span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </motion.div>

        {/* Asset Performance Chart */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.4 }}
          className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-lg mb-8"
        >
          <h3 className="text-xl font-bold text-gray-900 dark:text-white mb-6">Asset Performance</h3>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={assetPerformance}>
              <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
              <XAxis dataKey="symbol" stroke="#6B7280" />
              <YAxis stroke="#6B7280" />
              <Tooltip content={<CustomTooltip />} />
              <Legend />
              <Bar dataKey="performance" fill="#3B82F6" name="Performance %" />
              <Bar dataKey="volume" fill="#10B981" name="Volume" />
            </BarChart>
          </ResponsiveContainer>
        </motion.div>

        {/* Top Assets and Recent Transactions */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.5 }}
          className="grid grid-cols-1 lg:grid-cols-2 gap-8"
        >
          {/* Top Assets */}
          <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
            <div className="px-6 py-4 border-b border-gray-100 dark:border-gray-700">
              <h3 className="text-xl font-bold text-gray-900 dark:text-white">Top Assets</h3>
            </div>
            <div className="p-6">
              <div className="space-y-4">
                {summary?.top_assets.map((asset, index) => (
                  <motion.div
                    key={index}
                    initial={{ opacity: 0, x: -20 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: index * 0.1 }}
                    className="flex items-center justify-between p-4 bg-gray-50 dark:bg-gray-700 rounded-xl hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
                  >
                    <div className="flex items-center">
                      <div className="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-500 rounded-xl flex items-center justify-center text-white font-bold">
                        {asset.symbol}
                      </div>
                      <div className="ml-4">
                        <div className="font-medium text-gray-900 dark:text-white">{asset.name}</div>
                        <div className="text-sm text-gray-500 dark:text-gray-400">{asset.symbol}</div>
                      </div>
                    </div>
                    <div className="text-right">
                      <div className="font-medium text-gray-900 dark:text-white">{asset.value}</div>
                      <div className={`text-sm font-medium ${
                        asset.change_24h.startsWith('+') ? 'text-green-600' : 'text-red-600'
                      }`}>
                        {asset.change_24h}
                      </div>
                    </div>
                  </motion.div>
                ))}
              </div>
            </div>
          </div>

          {/* Recent Transactions */}
          <div className="bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
            <div className="px-6 py-4 border-b border-gray-100 dark:border-gray-700">
              <h3 className="text-xl font-bold text-gray-900 dark:text-white">Recent Transactions</h3>
            </div>
            <div className="p-6">
              <div className="space-y-4">
                {recentTransactions.map((tx, index) => (
                  <motion.div
                    key={tx.id}
                    initial={{ opacity: 0, x: 20 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: index * 0.1 }}
                    className="flex items-center justify-between p-4 bg-gray-50 dark:bg-gray-700 rounded-xl hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
                  >
                    <div className="flex items-center">
                      <div className={`w-10 h-10 rounded-xl flex items-center justify-center ${
                        tx.type === 'buy' ? 'bg-green-100 text-green-600' :
                        tx.type === 'sell' ? 'bg-red-100 text-red-600' :
                        'bg-blue-100 text-blue-600'
                      }`}>
                        {tx.type === 'buy' ? <ArrowUpRight className="w-5 h-5" /> :
                         tx.type === 'sell' ? <ArrowDownRight className="w-5 h-5" /> :
                         <Activity className="w-5 h-5" />}
                      </div>
                      <div className="ml-4">
                        <div className="font-medium text-gray-900 dark:text-white capitalize">{tx.type} {tx.asset}</div>
                        <div className="text-sm text-gray-500 dark:text-gray-400">{tx.time}</div>
                      </div>
                    </div>
                    <div className="text-right">
                      <div className="font-medium text-gray-900 dark:text-white">{tx.value}</div>
                      <div className={`text-sm font-medium ${
                        tx.status === 'completed' ? 'text-green-600' : 'text-yellow-600'
                      }`}>
                        {tx.status}
                      </div>
                    </div>
                  </motion.div>
                ))}
              </div>
            </div>
          </div>
        </motion.div>

        {/* Upgrade Banner for Free Users */}
        {isFreeUser && (
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.6 }}
            className="mt-8 bg-gradient-to-r from-purple-600 to-blue-600 rounded-2xl p-8 text-white"
          >
            <div className="flex items-center justify-between">
              <div>
                <h3 className="text-2xl font-bold mb-2">Unlock Premium Analytics</h3>
                <p className="text-purple-100 mb-4">
                  Get advanced charts, real-time alerts, and unlimited portfolio tracking with our premium plan.
                </p>
                <Link
                  to="/subscription"
                  className="inline-flex items-center px-6 py-3 bg-white text-purple-600 rounded-lg font-semibold hover:bg-gray-100 transition-colors"
                >
                  <Crown className="mr-2 w-5 h-5" />
                  Upgrade Now
                </Link>
              </div>
              <div className="hidden md:block">
                <BarChart3 className="w-16 h-16 text-purple-200" />
              </div>
            </div>
          </motion.div>
        )}
      </div>
    </div>
  )
}