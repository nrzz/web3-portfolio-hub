import React, { useState, useEffect } from 'react'

interface PerformanceData {
  period: string
  data: Array<{
    date: string
    value: string
    change: string
  }>
  total_return: string
  best_day: string
  worst_day: string
}

interface AllocationData {
  by_network: Record<string, {
    value: string
    percentage: string
    asset_count: number
  }>
  by_asset: Record<string, {
    value: string
    percentage: string
    amount: string
    network: string
  }>
}

export const Analytics: React.FC = () => {
  const [performance, setPerformance] = useState<PerformanceData | null>(null)
  const [allocation, setAllocation] = useState<AllocationData | null>(null)
  const [selectedPeriod, setSelectedPeriod] = useState('30d')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchAnalyticsData()
  }, [selectedPeriod])

  const fetchAnalyticsData = async () => {
    try {
      setLoading(true)
      
      // Mock performance data
      const mockPerformance: PerformanceData = {
        period: selectedPeriod,
        total_return: "+15.67%",
        best_day: "+8.2%",
        worst_day: "-3.1%",
        data: Array.from({ length: 30 }, (_, i) => ({
          date: new Date(Date.now() - (29 - i) * 24 * 60 * 60 * 1000).toISOString().split('T')[0],
          value: (100000 + Math.random() * 50000).toFixed(2),
          change: (Math.random() * 10 - 5).toFixed(2) + '%'
        }))
      }

      // Mock allocation data
      const mockAllocation: AllocationData = {
        by_network: {
          ethereum: { value: "$56,444.63", percentage: "45%", asset_count: 5 },
          polygon: { value: "$31,358.13", percentage: "25%", asset_count: 3 },
          bsc: { value: "$25,086.50", percentage: "20%", asset_count: 2 },
          arbitrum: { value: "$12,543.25", percentage: "10%", asset_count: 2 }
        },
        by_asset: {
          ETH: { value: "$48,750.00", percentage: "39%", amount: "2.5", network: "ethereum" },
          USDC: { value: "$10,000.00", percentage: "8%", amount: "10,000", network: "ethereum" },
          MATIC: { value: "$32,500.00", percentage: "26%", amount: "50,000", network: "polygon" },
          BNB: { value: "$25,000.00", percentage: "20%", amount: "10", network: "bsc" },
          ARB: { value: "$12,500.00", percentage: "7%", amount: "5,000", network: "arbitrum" }
        }
      }

      setPerformance(mockPerformance)
      setAllocation(mockAllocation)
    } catch (error) {
      console.error('Failed to fetch analytics data:', error)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Analytics</h1>
          <p className="text-gray-600">Portfolio performance and allocation insights</p>
        </div>
        <div className="flex space-x-2">
          {['7d', '30d', '90d', '1y'].map((period) => (
            <button
              key={period}
              onClick={() => setSelectedPeriod(period)}
              className={`px-3 py-2 rounded-lg font-medium ${
                selectedPeriod === period
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              {period}
            </button>
          ))}
        </div>
      </div>

      {/* Performance Overview */}
      {performance && (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-medium text-gray-900 mb-4">Performance Summary</h3>
            <div className="space-y-4">
              <div>
                <p className="text-sm text-gray-500">Total Return ({selectedPeriod})</p>
                <p className="text-2xl font-bold text-green-600">{performance.total_return}</p>
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-500">Best Day</p>
                  <p className="text-lg font-medium text-green-600">{performance.best_day}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">Worst Day</p>
                  <p className="text-lg font-medium text-red-600">{performance.worst_day}</p>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow p-6 lg:col-span-2">
            <h3 className="text-lg font-medium text-gray-900 mb-4">Performance Chart</h3>
            <div className="h-64 flex items-end justify-between space-x-1">
              {performance.data.slice(-14).map((point, index) => {
                const height = (parseFloat(point.value) / 150000) * 100
                const isPositive = point.change.startsWith('+')
                return (
                  <div key={index} className="flex-1 flex flex-col items-center">
                    <div
                      className={`w-full rounded-t ${
                        isPositive ? 'bg-green-500' : 'bg-red-500'
                      }`}
                      style={{ height: `${Math.max(height, 5)}%` }}
                    ></div>
                    <div className="text-xs text-gray-500 mt-2 transform rotate-45 origin-left">
                      {point.date.split('-').slice(1).join('/')}
                    </div>
                  </div>
                )
              })}
            </div>
          </div>
        </div>
      )}

      {/* Allocation Breakdown */}
      {allocation && (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Network Allocation */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-medium text-gray-900 mb-4">Network Allocation</h3>
            <div className="space-y-4">
              {Object.entries(allocation.by_network).map(([network, data]) => (
                <div key={network} className="flex items-center justify-between">
                  <div className="flex items-center">
                    <div className="w-4 h-4 rounded-full bg-blue-500 mr-3"></div>
                    <div>
                      <p className="text-sm font-medium text-gray-900 capitalize">{network}</p>
                      <p className="text-xs text-gray-500">{data.asset_count} assets</p>
                    </div>
                  </div>
                  <div className="text-right">
                    <p className="text-sm font-medium text-gray-900">{data.value}</p>
                    <p className="text-xs text-gray-500">{data.percentage}</p>
                  </div>
                </div>
              ))}
            </div>
            
            {/* Pie Chart Visualization */}
            <div className="mt-6">
              <div className="flex justify-center">
                <div className="relative w-32 h-32">
                  <svg className="w-32 h-32 transform -rotate-90" viewBox="0 0 32 32">
                    <circle
                      cx="16"
                      cy="16"
                      r="16"
                      fill="none"
                      stroke="#e5e7eb"
                      strokeWidth="8"
                    />
                    <circle
                      cx="16"
                      cy="16"
                      r="16"
                      fill="none"
                      stroke="#3b82f6"
                      strokeWidth="8"
                      strokeDasharray={`${2 * Math.PI * 16}`}
                      strokeDashoffset={`${2 * Math.PI * 16 * 0.55}`}
                    />
                  </svg>
                  <div className="absolute inset-0 flex items-center justify-center">
                    <span className="text-sm font-medium text-gray-900">45%</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Asset Allocation */}
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-medium text-gray-900 mb-4">Asset Allocation</h3>
            <div className="space-y-4">
              {Object.entries(allocation.by_asset).map(([asset, data]) => (
                <div key={asset} className="flex items-center justify-between">
                  <div className="flex items-center">
                    <div className="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center mr-3">
                      <span className="text-sm font-medium text-gray-900">{asset}</span>
                    </div>
                    <div>
                      <p className="text-sm font-medium text-gray-900">{asset}</p>
                      <p className="text-xs text-gray-500">{data.amount} on {data.network}</p>
                    </div>
                  </div>
                  <div className="text-right">
                    <p className="text-sm font-medium text-gray-900">{data.value}</p>
                    <p className="text-xs text-gray-500">{data.percentage}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* Risk Metrics */}
      <div className="bg-white rounded-lg shadow p-6">
        <h3 className="text-lg font-medium text-gray-900 mb-4">Risk Metrics</h3>
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div className="text-center">
            <p className="text-sm text-gray-500">Volatility</p>
            <p className="text-2xl font-bold text-gray-900">12.5%</p>
            <p className="text-xs text-gray-500">Annualized</p>
          </div>
          <div className="text-center">
            <p className="text-sm text-gray-500">Sharpe Ratio</p>
            <p className="text-2xl font-bold text-gray-900">1.85</p>
            <p className="text-xs text-gray-500">Risk-adjusted return</p>
          </div>
          <div className="text-center">
            <p className="text-sm text-gray-500">Max Drawdown</p>
            <p className="text-2xl font-bold text-red-600">-8.2%</p>
            <p className="text-xs text-gray-500">Peak to trough</p>
          </div>
          <div className="text-center">
            <p className="text-sm text-gray-500">Beta</p>
            <p className="text-2xl font-bold text-gray-900">0.92</p>
            <p className="text-xs text-gray-500">vs BTC</p>
          </div>
        </div>
      </div>

      {/* Recent Performance */}
      {performance && (
        <div className="bg-white rounded-lg shadow">
          <div className="px-6 py-4 border-b border-gray-200">
            <h3 className="text-lg font-medium text-gray-900">Recent Performance</h3>
          </div>
          <div className="overflow-x-auto">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Date</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Value</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Change</th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {performance.data.slice(-10).reverse().map((point, index) => (
                  <tr key={index}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{point.date}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                      ${parseFloat(point.value).toLocaleString()}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <span className={`text-sm font-medium ${
                        point.change.startsWith('+') ? 'text-green-600' : 'text-red-600'
                      }`}>
                        {point.change}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      )}
    </div>
  )
} 