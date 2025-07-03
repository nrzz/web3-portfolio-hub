import React, { useState } from 'react'
import { motion } from 'framer-motion'
import { useAuth } from '../contexts/AuthContext'
import { Check, Crown, Sparkles, Zap, Shield, BarChart3, Bell, Download, Users, Globe } from 'lucide-react'

const plans = [
  {
    id: 'free',
    name: 'Free',
    price: 0,
    interval: 'forever',
    description: 'Perfect for getting started',
    features: [
      'Check wallet balances',
      'Basic portfolio overview',
      'Support for 4 networks',
      'Community support'
    ],
    limitations: [
      'No analytics',
      'No alerts',
      'No data export',
      'No priority support'
    ],
    popular: false,
    gradient: 'from-gray-400 to-gray-600'
  },
  {
    id: 'subscriber',
    name: 'Premium',
    price: 9.99,
    interval: 'month',
    description: 'For serious crypto investors',
    features: [
      'Everything in Free',
      'Advanced analytics & charts',
      'Price & balance alerts',
      'Data export (CSV, PDF)',
      'Priority support',
      'Historical performance',
      'Network allocation charts',
      'Transaction tracking'
    ],
    limitations: [],
    popular: true,
    gradient: 'from-blue-500 to-purple-600'
  },
  {
    id: 'pro',
    name: 'Pro',
    price: 29.99,
    interval: 'month',
    description: 'For professional traders',
    features: [
      'Everything in Premium',
      'Multi-wallet tracking',
      'Custom dashboards',
      'Real-time notifications',
      'API access',
      'Tax reporting tools',
      'DeFi protocol integration',
      'Advanced portfolio analytics'
    ],
    limitations: [],
    popular: false,
    gradient: 'from-yellow-400 to-orange-500'
  }
]

export const Subscription: React.FC = () => {
  const { user, updateSubscription } = useAuth()
  const [selectedPlan, setSelectedPlan] = useState(user?.plan === 'free' ? 'subscriber' : user?.plan || 'subscriber')
  const [isUpgrading, setIsUpgrading] = useState(false)

  const handleUpgrade = async (planId: string) => {
    if (planId === 'free') return
    
    setIsUpgrading(true)
    try {
      await updateSubscription(planId, 'active')
      // In a real app, you'd redirect to payment processor
      alert(`Upgraded to ${planId} plan! (This is a demo)`)
    } catch (error) {
      console.error('Upgrade failed:', error)
      alert('Upgrade failed. Please try again.')
    } finally {
      setIsUpgrading(false)
    }
  }

  const containerVariants = {
    hidden: { opacity: 0 },
    visible: {
      opacity: 1,
      transition: {
        staggerChildren: 0.1
      }
    }
  }

  const cardVariants = {
    hidden: { opacity: 0, y: 20 },
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        type: 'spring',
        stiffness: 100,
        damping: 15
      }
    }
  }

  return (
    <motion.div
      variants={containerVariants}
      initial="hidden"
      animate="visible"
      className="space-y-8"
    >
      {/* Header */}
      <motion.div variants={cardVariants} className="text-center">
        <motion.h1 
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="text-4xl font-bold bg-gradient-to-r from-gray-900 to-gray-600 dark:from-white dark:to-gray-300 bg-clip-text text-transparent mb-4"
        >
          Choose Your Plan
        </motion.h1>
        <motion.p 
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto"
        >
          Unlock powerful features to take your crypto portfolio to the next level
        </motion.p>
      </motion.div>

      {/* Feature Comparison */}
      <motion.div variants={cardVariants} className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-xl border border-gray-100 dark:border-gray-700">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-6">Feature Comparison</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {plans.map((plan) => (
            <div key={plan.id} className="text-center">
              <div className={`w-12 h-12 mx-auto mb-4 rounded-xl bg-gradient-to-r ${plan.gradient} flex items-center justify-center`}>
                {plan.id === 'free' ? (
                  <Sparkles className="w-6 h-6 text-white" />
                ) : plan.id === 'subscriber' ? (
                  <Crown className="w-6 h-6 text-white" />
                ) : (
                  <Zap className="w-6 h-6 text-white" />
                )}
              </div>
              <h3 className="text-lg font-semibold text-gray-900 dark:text-white mb-2">{plan.name}</h3>
              <div className="text-sm text-gray-600 dark:text-gray-400">
                {plan.features.length} features
              </div>
            </div>
          ))}
        </div>
      </motion.div>

      {/* Plans */}
      <motion.div variants={cardVariants} className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {plans.map((plan, index) => (
          <motion.div
            key={plan.id}
            variants={cardVariants}
            whileHover={{ y: -5, scale: 1.02 }}
            className={`relative bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-xl border-2 transition-all duration-300 ${
              selectedPlan === plan.id 
                ? 'border-blue-500 shadow-blue-500/25' 
                : 'border-gray-100 dark:border-gray-700'
            } ${plan.popular ? 'ring-2 ring-purple-500 ring-opacity-50' : ''}`}
          >
            {plan.popular && (
              <div className="absolute -top-3 left-1/2 transform -translate-x-1/2">
                <span className="bg-gradient-to-r from-purple-500 to-pink-500 text-white text-xs font-bold px-3 py-1 rounded-full">
                  MOST POPULAR
                </span>
              </div>
            )}

            <div className="text-center mb-6">
              <div className={`w-16 h-16 mx-auto mb-4 rounded-2xl bg-gradient-to-r ${plan.gradient} flex items-center justify-center`}>
                {plan.id === 'free' ? (
                  <Sparkles className="w-8 h-8 text-white" />
                ) : plan.id === 'subscriber' ? (
                  <Crown className="w-8 h-8 text-white" />
                ) : (
                  <Zap className="w-8 h-8 text-white" />
                )}
              </div>
              <h3 className="text-2xl font-bold text-gray-900 dark:text-white mb-2">{plan.name}</h3>
              <p className="text-gray-600 dark:text-gray-400 mb-4">{plan.description}</p>
              <div className="mb-6">
                <span className="text-4xl font-bold text-gray-900 dark:text-white">
                  ${plan.price}
                </span>
                <span className="text-gray-600 dark:text-gray-400">/{plan.interval}</span>
              </div>
            </div>

            <div className="space-y-4 mb-8">
              <h4 className="font-semibold text-gray-900 dark:text-white mb-3">Features:</h4>
              {plan.features.map((feature, featureIndex) => (
                <motion.div
                  key={featureIndex}
                  initial={{ opacity: 0, x: -20 }}
                  animate={{ opacity: 1, x: 0 }}
                  transition={{ delay: index * 0.1 + featureIndex * 0.05 }}
                  className="flex items-center space-x-3"
                >
                  <div className="w-5 h-5 bg-green-100 dark:bg-green-900 rounded-full flex items-center justify-center flex-shrink-0">
                    <Check className="w-3 h-3 text-green-600 dark:text-green-400" />
                  </div>
                  <span className="text-sm text-gray-700 dark:text-gray-300">{feature}</span>
                </motion.div>
              ))}
              
              {plan.limitations.length > 0 && (
                <>
                  <h4 className="font-semibold text-gray-900 dark:text-white mb-3 mt-6">Limitations:</h4>
                  {plan.limitations.map((limitation, limitationIndex) => (
                    <motion.div
                      key={limitationIndex}
                      initial={{ opacity: 0, x: -20 }}
                      animate={{ opacity: 1, x: 0 }}
                      transition={{ delay: index * 0.1 + limitationIndex * 0.05 }}
                      className="flex items-center space-x-3"
                    >
                      <div className="w-5 h-5 bg-red-100 dark:bg-red-900 rounded-full flex items-center justify-center flex-shrink-0">
                        <span className="w-3 h-3 text-red-600 dark:text-red-400 text-xs">×</span>
                      </div>
                      <span className="text-sm text-gray-500 dark:text-gray-400">{limitation}</span>
                    </motion.div>
                  ))}
                </>
              )}
            </div>

            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              onClick={() => handleUpgrade(plan.id)}
              disabled={isUpgrading || plan.id === 'free'}
              className={`w-full py-3 px-4 rounded-xl font-medium transition-all duration-300 ${
                plan.id === 'free'
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
                  : selectedPlan === plan.id
                  ? 'bg-gradient-to-r from-blue-500 to-purple-600 text-white shadow-lg'
                  : 'bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-600'
              }`}
            >
              {isUpgrading ? (
                <div className="flex items-center justify-center">
                  <div className="w-4 h-4 border-2 border-current border-t-transparent rounded-full animate-spin mr-2"></div>
                  Upgrading...
                </div>
              ) : plan.id === 'free' ? (
                'Current Plan'
              ) : (
                `Upgrade to ${plan.name}`
              )}
            </motion.button>
          </motion.div>
        ))}
      </motion.div>

      {/* Additional Features */}
      <motion.div variants={cardVariants} className="bg-gradient-to-r from-blue-50 to-purple-50 dark:from-blue-900/20 dark:to-purple-900/20 rounded-2xl p-8">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-6 text-center">
          Why Choose Premium?
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {[
            { icon: BarChart3, title: 'Advanced Analytics', desc: 'Deep insights into your portfolio performance' },
            { icon: Bell, title: 'Smart Alerts', desc: 'Never miss important price movements' },
            { icon: Download, title: 'Data Export', desc: 'Export your data for tax and analysis' },
            { icon: Shield, title: 'Priority Support', desc: 'Get help when you need it most' }
          ].map((feature, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: index * 0.1 }}
              className="text-center"
            >
              <div className="w-12 h-12 mx-auto mb-4 bg-gradient-to-r from-blue-500 to-purple-500 rounded-xl flex items-center justify-center">
                <feature.icon className="w-6 h-6 text-white" />
              </div>
              <h3 className="font-semibold text-gray-900 dark:text-white mb-2">{feature.title}</h3>
              <p className="text-sm text-gray-600 dark:text-gray-400">{feature.desc}</p>
            </motion.div>
          ))}
        </div>
      </motion.div>

      {/* FAQ */}
      <motion.div variants={cardVariants} className="bg-white dark:bg-gray-800 rounded-2xl p-6 shadow-xl border border-gray-100 dark:border-gray-700">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-6">Frequently Asked Questions</h2>
        <div className="space-y-4">
          {[
            {
              q: 'Can I cancel my subscription anytime?',
              a: 'Yes, you can cancel your subscription at any time. You\'ll continue to have access until the end of your billing period.'
            },
            {
              q: 'Do you offer refunds?',
              a: 'We offer a 30-day money-back guarantee for all paid plans.'
            },
            {
              q: 'Can I upgrade or downgrade my plan?',
              a: 'Yes, you can change your plan at any time. Changes take effect immediately.'
            },
            {
              q: 'Is my data secure?',
              a: 'Absolutely. We use bank-level encryption and never store your private keys.'
            }
          ].map((faq, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ delay: index * 0.1 }}
              className="border-b border-gray-200 dark:border-gray-700 pb-4 last:border-b-0"
            >
              <h3 className="font-semibold text-gray-900 dark:text-white mb-2">{faq.q}</h3>
              <p className="text-gray-600 dark:text-gray-400">{faq.a}</p>
            </motion.div>
          ))}
        </div>
      </motion.div>
    </motion.div>
  )
} 