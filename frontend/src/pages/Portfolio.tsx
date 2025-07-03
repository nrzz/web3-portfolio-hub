import React, { useState, useEffect } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { RefreshCw, Plus, ExternalLink, Copy, CheckCircle, AlertCircle } from 'lucide-react'
import { useAuth } from '../contexts/AuthContext'

interface Portfolio {
  id: string
  name: string
  addresses: Address[]
  created_at: string
  updated_at: string
}

interface Address {
  id: string
  address: string
  network: string
  label: string
  is_active: boolean
  balances: Balance[]
}

interface Balance {
  id: string
  token_address: string
  symbol: string
  name: string
  amount: string
  decimals: number
  price: string
  value: string
  updated_at: string
}

export const Portfolio: React.FC = () => {
  const { user } = useAuth()
  const [portfolios, setPortfolios] = useState<Portfolio[]>([])
  const [selectedPortfolio, setSelectedPortfolio] = useState<Portfolio | null>(null)
  const [loading, setLoading] = useState(true)
  const [refreshing, setRefreshing] = useState(false)
  const [showAddAddress, setShowAddAddress] = useState(false)
  const [copiedAddress, setCopiedAddress] = useState<string | null>(null)
  const [newAddress, setNewAddress] = useState({
    address: '',
    network: 'ethereum',
    label: ''
  })

  useEffect(() => {
    fetchPortfolios()
  }, [])

  const fetchPortfolios = async () => {
    try {
      setLoading(true)
      const token = localStorage.getItem('token')
      
      if (!token) {
        throw new Error('No authentication token')
      }

      const response = await fetch('/api/v1/portfolios', {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Failed to fetch portfolios')
      }

      const data = await response.json()
      setPortfolios(data.portfolios || [])
      
      if (data.portfolios && data.portfolios.length > 0) {
        setSelectedPortfolio(data.portfolios[0])
      }
    } catch (error) {
      console.error('Failed to fetch portfolios:', error)
      // Show error, do not use mock data
      setPortfolios([])
      setSelectedPortfolio(null)
      alert('Failed to fetch portfolios from backend. Please check your connection and try again.')
    } finally {
      setLoading(false)
    }
  }

  const handleAddAddress = async () => {
    if (!newAddress.address || !newAddress.network || !selectedPortfolio) return

    try {
      const token = localStorage.getItem('token')
      if (!token) throw new Error('No authentication token')

      const response = await fetch(`/api/v1/portfolios/${selectedPortfolio.id}/addresses`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          address: newAddress.address,
          network: newAddress.network,
          label: newAddress.label || 'New Address'
        }),
      })

      if (!response.ok) {
        throw new Error('Failed to add address')
      }

      // Refresh portfolios to get updated data
      await fetchPortfolios()
      setNewAddress({ address: '', network: 'ethereum', label: '' })
      setShowAddAddress(false)
    } catch (error) {
      console.error('Failed to add address:', error)
      alert('Failed to add address. Please try again.')
    }
  }

  const handleRefreshBalances = async (portfolioId: string) => {
    if (!selectedPortfolio) return
    
    try {
      setRefreshing(true)
      const token = localStorage.getItem('token')
      if (!token) throw new Error('No authentication token')

      const response = await fetch(`/api/v1/portfolios/${portfolioId}/balances/refresh`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Failed to refresh balances')
      }

      // Refresh portfolios to get updated balances
      await fetchPortfolios()
    } catch (error) {
      console.error('Failed to refresh balances:', error)
      alert('Failed to refresh balances. Please try again.')
    } finally {
      setRefreshing(false)
    }
  }

  const copyToClipboard = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text)
      setCopiedAddress(text)
      setTimeout(() => setCopiedAddress(null), 2000)
    } catch (error) {
      console.error('Failed to copy to clipboard:', error)
    }
  }

  const getNetworkIcon = (network: string) => {
    const icons: { [key: string]: string } = {
      ethereum: '🔵',
      polygon: '🟣',
      bsc: '🟡',
      arbitrum: '🔵'
    }
    return icons[network] || '🔗'
  }

  const getNetworkColor = (network: string) => {
    const colors: { [key: string]: string } = {
      ethereum: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
      polygon: 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200',
      bsc: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
      arbitrum: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200'
    }
    return colors[network] || 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200'
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <motion.div
          animate={{ rotate: 360 }}
          transition={{ duration: 1, repeat: Infinity, ease: "linear" }}
          className="w-12 h-12 border-4 border-blue-200 border-t-blue-600 rounded-full"
        />
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <motion.div 
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        className="flex justify-between items-center"
      >
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Portfolio</h1>
          <p className="text-gray-600 dark:text-gray-400">Manage your blockchain addresses and assets</p>
        </div>
        <div className="flex space-x-3">
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            onClick={() => selectedPortfolio && handleRefreshBalances(selectedPortfolio.id)}
            disabled={refreshing}
            className="flex items-center space-x-2 bg-gray-600 hover:bg-gray-700 disabled:bg-gray-400 text-white px-4 py-2 rounded-lg font-medium transition-colors"
          >
            <RefreshCw className={`w-4 h-4 ${refreshing ? 'animate-spin' : ''}`} />
            <span>{refreshing ? 'Refreshing...' : 'Refresh Balances'}</span>
          </motion.button>
          <motion.button
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            onClick={() => setShowAddAddress(true)}
            className="flex items-center space-x-2 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-300 shadow-lg hover:shadow-xl"
          >
            <Plus className="w-4 h-4" />
            <span>Add Address</span>
          </motion.button>
        </div>
      </motion.div>

      {/* Portfolio Selection */}
      {portfolios.length > 0 && (
        <motion.div 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="bg-white dark:bg-gray-800 rounded-xl shadow-lg border border-gray-200 dark:border-gray-700 p-6"
        >
          <h2 className="text-lg font-medium text-gray-900 dark:text-white mb-4">Select Portfolio</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            {portfolios.map((portfolio) => (
              <motion.div
                key={portfolio.id}
                whileHover={{ scale: 1.02 }}
                whileTap={{ scale: 0.98 }}
                onClick={() => setSelectedPortfolio(portfolio)}
                className={`p-4 border rounded-xl cursor-pointer transition-all duration-300 ${
                  selectedPortfolio?.id === portfolio.id
                    ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/20 shadow-lg'
                    : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600 bg-gray-50 dark:bg-gray-700/50'
                }`}
              >
                <h3 className="font-medium text-gray-900 dark:text-white">{portfolio.name}</h3>
                <p className="text-sm text-gray-500 dark:text-gray-400">{portfolio.addresses.length} addresses</p>
              </motion.div>
            ))}
          </div>
        </motion.div>
      )}

      {/* Add Address Modal */}
      <AnimatePresence>
        {showAddAddress && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 bg-black bg-opacity-50 overflow-y-auto h-full w-full z-50 flex items-center justify-center p-4"
          >
            <motion.div
              initial={{ scale: 0.9, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              exit={{ scale: 0.9, opacity: 0 }}
              className="relative bg-white dark:bg-gray-800 rounded-xl shadow-2xl w-full max-w-md p-6"
            >
              <h3 className="text-lg font-medium text-gray-900 dark:text-white mb-4">Add New Address</h3>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Address</label>
                  <input
                    type="text"
                    value={newAddress.address}
                    onChange={(e) => setNewAddress({ ...newAddress, address: e.target.value })}
                    className="mt-1 block w-full border border-gray-300 dark:border-gray-600 rounded-lg px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="0x..."
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Network</label>
                  <select
                    value={newAddress.network}
                    onChange={(e) => setNewAddress({ ...newAddress, network: e.target.value })}
                    className="mt-1 block w-full border border-gray-300 dark:border-gray-600 rounded-lg px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="ethereum">Ethereum</option>
                    <option value="polygon">Polygon</option>
                    <option value="bsc">BSC</option>
                    <option value="arbitrum">Arbitrum</option>
                  </select>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Label (Optional)</label>
                  <input
                    type="text"
                    value={newAddress.label}
                    onChange={(e) => setNewAddress({ ...newAddress, label: e.target.value })}
                    className="mt-1 block w-full border border-gray-300 dark:border-gray-600 rounded-lg px-3 py-2 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="My Wallet"
                  />
                </div>
              </div>
              <div className="flex justify-end space-x-3 mt-6">
                <button
                  onClick={() => setShowAddAddress(false)}
                  className="px-4 py-2 text-gray-600 dark:text-gray-400 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
                >
                  Cancel
                </button>
                <button
                  onClick={handleAddAddress}
                  className="px-4 py-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-lg hover:from-blue-700 hover:to-purple-700 transition-all duration-300"
                >
                  Add Address
                </button>
              </div>
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>

      {/* Portfolio Details */}
      {selectedPortfolio && (
        <div className="space-y-6">
          {/* Portfolio Summary */}
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2 }}
            className="bg-white dark:bg-gray-800 rounded-xl shadow-lg border border-gray-200 dark:border-gray-700 p-6"
          >
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold text-gray-900 dark:text-white">{selectedPortfolio.name}</h2>
              <span className="text-sm text-gray-500 dark:text-gray-400">
                {selectedPortfolio.addresses.length} addresses
              </span>
            </div>
            
            {/* Total Portfolio Value */}
            <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
              <div className="text-center p-4 bg-gradient-to-r from-blue-50 to-purple-50 dark:from-blue-900/20 dark:to-purple-900/20 rounded-xl">
                <p className="text-sm text-gray-500 dark:text-gray-400 mb-1">Total Value</p>
                <p className="text-3xl font-bold text-gray-900 dark:text-white">
                  ${selectedPortfolio.addresses.reduce((total, addr) => 
                    total + addr.balances.reduce((sum, bal) => sum + parseFloat(bal.value || '0'), 0), 0
                  ).toLocaleString()}
                </p>
              </div>
              <div className="text-center p-4 bg-gradient-to-r from-green-50 to-emerald-50 dark:from-green-900/20 dark:to-emerald-900/20 rounded-xl">
                <p className="text-sm text-gray-500 dark:text-gray-400 mb-1">Total Assets</p>
                <p className="text-3xl font-bold text-gray-900 dark:text-white">
                  {selectedPortfolio.addresses.reduce((total, addr) => total + addr.balances.length, 0)}
                </p>
              </div>
              <div className="text-center p-4 bg-gradient-to-r from-purple-50 to-pink-50 dark:from-purple-900/20 dark:to-pink-900/20 rounded-xl">
                <p className="text-sm text-gray-500 dark:text-gray-400 mb-1">Networks</p>
                <p className="text-3xl font-bold text-gray-900 dark:text-white">
                  {new Set(selectedPortfolio.addresses.map(addr => addr.network)).size}
                </p>
              </div>
              <div className="text-center p-4 bg-gradient-to-r from-orange-50 to-red-50 dark:from-orange-900/20 dark:to-red-900/20 rounded-xl">
                <p className="text-sm text-gray-500 dark:text-gray-400 mb-1">Last Updated</p>
                <p className="text-lg font-bold text-gray-900 dark:text-white">
                  {new Date(selectedPortfolio.updated_at).toLocaleDateString()}
                </p>
              </div>
            </div>
          </motion.div>

          {/* Addresses */}
          <div className="space-y-4">
            {selectedPortfolio.addresses.map((address, index) => (
              <motion.div
                key={address.id}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.3 + index * 0.1 }}
                className="bg-white dark:bg-gray-800 rounded-xl shadow-lg border border-gray-200 dark:border-gray-700 overflow-hidden"
              >
                <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700 bg-gradient-to-r from-gray-50 to-gray-100 dark:from-gray-700 dark:to-gray-800">
                  <div className="flex justify-between items-center">
                    <div className="flex-1">
                      <div className="flex items-center space-x-3">
                        <h3 className="text-lg font-medium text-gray-900 dark:text-white">
                          {address.label || 'Unnamed Address'}
                        </h3>
                        <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getNetworkColor(address.network)}`}>
                          {getNetworkIcon(address.network)} {address.network}
                        </span>
                      </div>
                      <div className="flex items-center space-x-2 mt-1">
                        <p className="text-sm text-gray-500 dark:text-gray-400 font-mono">{address.address}</p>
                        <button
                          onClick={() => copyToClipboard(address.address)}
                          className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded transition-colors"
                        >
                          {copiedAddress === address.address ? (
                            <CheckCircle className="w-4 h-4 text-green-500" />
                          ) : (
                            <Copy className="w-4 h-4 text-gray-400" />
                          )}
                        </button>
                        <a
                          href={`https://etherscan.io/address/${address.address}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="p-1 hover:bg-gray-200 dark:hover:bg-gray-600 rounded transition-colors"
                        >
                          <ExternalLink className="w-4 h-4 text-gray-400" />
                        </a>
                      </div>
                    </div>
                    <div className="text-right">
                      <p className="text-sm text-gray-500 dark:text-gray-400">Total Value</p>
                      <p className="text-xl font-bold text-gray-900 dark:text-white">
                        ${address.balances.reduce((sum, bal) => sum + parseFloat(bal.value || '0'), 0).toLocaleString()}
                      </p>
                    </div>
                  </div>
                </div>
                
                {/* Balances */}
                <div className="px-6 py-4">
                  <h4 className="text-sm font-medium text-gray-900 dark:text-white mb-4">Balances</h4>
                  <div className="space-y-3">
                    {address.balances.length > 0 ? (
                      address.balances.map((balance) => (
                        <motion.div
                          key={balance.id}
                          whileHover={{ scale: 1.02 }}
                          className="flex justify-between items-center p-4 bg-gray-50 dark:bg-gray-700 rounded-xl border border-gray-200 dark:border-gray-600"
                        >
                          <div className="flex items-center space-x-3">
                            <div className="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full flex items-center justify-center shadow-lg">
                              <span className="text-sm font-bold text-white">{balance.symbol}</span>
                            </div>
                            <div>
                              <p className="text-sm font-medium text-gray-900 dark:text-white">{balance.name}</p>
                              <p className="text-xs text-gray-500 dark:text-gray-400">
                                {parseFloat(balance.amount).toLocaleString()} {balance.symbol}
                              </p>
                            </div>
                          </div>
                          <div className="text-right">
                            <p className="text-sm font-bold text-gray-900 dark:text-white">${parseFloat(balance.value).toLocaleString()}</p>
                            <p className="text-xs text-gray-500 dark:text-gray-400">
                              ${parseFloat(balance.price).toLocaleString()} per {balance.symbol}
                            </p>
                          </div>
                        </motion.div>
                      ))
                    ) : (
                      <div className="text-center py-8">
                        <AlertCircle className="w-12 h-12 text-gray-400 mx-auto mb-2" />
                        <p className="text-gray-500 dark:text-gray-400">No balances found for this address</p>
                      </div>
                    )}
                  </div>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      )}

      {/* Empty State */}
      {portfolios.length === 0 && (
        <motion.div 
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className="text-center py-12"
        >
          <div className="w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </div>
          <h3 className="text-lg font-medium text-gray-900 dark:text-white mb-2">No portfolios yet</h3>
          <p className="text-gray-500 dark:text-gray-400 mb-6">Create your first portfolio to start tracking your assets</p>
          <motion.button 
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
            className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white px-6 py-3 rounded-lg font-medium transition-all duration-300 shadow-lg hover:shadow-xl"
          >
            Create Portfolio
          </motion.button>
        </motion.div>
      )}
    </div>
  )
} 