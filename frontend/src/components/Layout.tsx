import React from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { motion, AnimatePresence } from 'framer-motion'
import { Home, PieChart, Wallet, Bell, Settings, Sun, Moon, Plus, Menu, X, Crown, Sparkles, LogOut, User, ChevronDown } from 'lucide-react'
import { Switch } from '@headlessui/react'
import { useAuth } from '../contexts/AuthContext'

const navLinks = [
  { name: 'Dashboard', path: '/dashboard', icon: Home },
  { name: 'Portfolio', path: '/portfolio', icon: Wallet },
  { name: 'Analytics', path: '/analytics', icon: PieChart, requiresSubscription: true },
  { name: 'Alerts', path: '/alerts', icon: Bell, requiresSubscription: true },
  { name: 'Settings', path: '/settings', icon: Settings },
]

export const Layout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const location = useLocation()
  const navigate = useNavigate()
  const { user, logout } = useAuth()
  const [darkMode, setDarkMode] = React.useState(() => {
    if (typeof window !== 'undefined') {
      return localStorage.getItem('darkMode') === 'true' || 
             window.matchMedia('(prefers-color-scheme: dark)').matches
    }
    return false
  })
  const [mobileMenuOpen, setMobileMenuOpen] = React.useState(false)
  const [userMenuOpen, setUserMenuOpen] = React.useState(false)

  const isFreeUser = user?.plan === 'free'

  React.useEffect(() => {
    if (darkMode) {
      document.documentElement.classList.add('dark')
      localStorage.setItem('darkMode', 'true')
    } else {
      document.documentElement.classList.remove('dark')
      localStorage.setItem('darkMode', 'false')
    }
  }, [darkMode])

  const handleLogout = () => {
    logout()
    navigate('/login')
    setUserMenuOpen(false)
  }

  const sidebarVariants = {
    hidden: { x: -300, opacity: 0 },
    visible: { 
      x: 0, 
      opacity: 1,
      transition: { 
        type: 'spring', 
        stiffness: 100, 
        damping: 20,
        staggerChildren: 0.1
      }
    }
  }

  const navItemVariants = {
    hidden: { x: -20, opacity: 0 },
    visible: { 
      x: 0, 
      opacity: 1,
      transition: { type: 'spring', stiffness: 300, damping: 30 }
    }
  }

  return (
    <div className="min-h-screen flex bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 transition-all duration-500">
      {/* Mobile Menu Overlay */}
      <AnimatePresence>
        {mobileMenuOpen && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 bg-black bg-opacity-50 z-40 md:hidden"
            onClick={() => setMobileMenuOpen(false)}
          />
        )}
      </AnimatePresence>

      {/* Sidebar */}
      <motion.aside
        variants={sidebarVariants}
        initial="hidden"
        animate="visible"
        className={`fixed md:relative z-50 flex flex-col w-80 md:w-72 bg-white/90 dark:bg-gray-950/90 backdrop-blur-xl border-r border-gray-200/50 dark:border-gray-800/50 shadow-2xl ${
          mobileMenuOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'
        } transition-transform duration-300 ease-in-out`}
      >
        {/* Logo Section */}
        <motion.div 
          variants={navItemVariants}
          className="flex items-center justify-center h-20 border-b border-gray-200/50 dark:border-gray-800/50 bg-gradient-to-r from-blue-600 via-purple-600 to-indigo-600"
        >
          <div className="flex items-center space-x-3">
            <div className="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center backdrop-blur-sm">
              <Wallet className="w-6 h-6 text-white" />
            </div>
            <span className="text-2xl font-bold text-white tracking-tight">Web3Dash</span>
          </div>
        </motion.div>

        {/* Plan Status */}
        <motion.div 
          variants={navItemVariants}
          className="px-6 py-4 border-b border-gray-200/50 dark:border-gray-800/50"
        >
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              {isFreeUser ? (
                <div className="w-6 h-6 bg-gradient-to-r from-gray-400 to-gray-500 rounded-full flex items-center justify-center shadow-sm">
                  <span className="text-xs text-white font-bold">F</span>
                </div>
              ) : (
                <div className="w-6 h-6 bg-gradient-to-r from-yellow-400 via-orange-500 to-red-500 rounded-full flex items-center justify-center shadow-sm">
                  <Crown className="w-3 h-3 text-white" />
                </div>
              )}
              <div>
                <div className="text-sm font-medium text-gray-900 dark:text-white">
                  {isFreeUser ? 'Free Plan' : 'Premium Plan'}
                </div>
                <div className="text-xs text-gray-500 dark:text-gray-400">
                  {isFreeUser ? 'Basic features only' : 'Full access'}
                </div>
              </div>
            </div>
          </div>
        </motion.div>

        {/* Navigation */}
        <nav className="flex-1 py-8 px-6 space-y-3">
          {navLinks.map(({ name, path, icon: Icon, requiresSubscription }, index) => {
            const isLocked = requiresSubscription && isFreeUser
            const isActive = location.pathname === path
            
            return (
              <motion.div key={path} variants={navItemVariants}>
                <Link
                  to={path}
                  onClick={() => setMobileMenuOpen(false)}
                  className={`group relative flex items-center px-4 py-3 rounded-xl font-medium transition-all duration-300 transform hover:scale-105 ${
                    isActive
                      ? 'bg-gradient-to-r from-blue-500 to-purple-500 text-white shadow-lg shadow-blue-500/25'
                      : isLocked
                      ? 'text-gray-400 dark:text-gray-500 cursor-not-allowed opacity-50'
                      : 'text-gray-700 dark:text-gray-300 hover:bg-gray-100/80 dark:hover:bg-gray-800/80'
                  }`}
                >
                  <Icon className={`w-5 h-5 mr-3 transition-all duration-300 ${
                    isActive 
                      ? 'text-white' 
                      : isLocked
                      ? 'text-gray-400'
                      : 'text-blue-500 group-hover:scale-110'
                  }`} />
                  {name}
                  {isLocked && (
                    <div className="ml-auto">
                      <Sparkles className="w-4 h-4 text-gray-400" />
                    </div>
                  )}
                  {isActive && (
                    <motion.div
                      layoutId="activeTab"
                      className="absolute inset-0 bg-gradient-to-r from-blue-500 to-purple-500 rounded-xl -z-10"
                      initial={false}
                      transition={{ type: "spring", stiffness: 500, damping: 30 }}
                    />
                  )}
                </Link>
              </motion.div>
            )
          })}
        </nav>

        {/* Upgrade Banner for Free Users */}
        {isFreeUser && (
          <motion.div 
            variants={navItemVariants}
            className="mx-6 mb-4 p-4 bg-gradient-to-r from-purple-500 via-pink-500 to-red-500 rounded-xl text-white shadow-lg"
          >
            <div className="flex items-center space-x-2 mb-2">
              <Crown className="w-4 h-4" />
              <span className="text-sm font-medium">Upgrade to Premium</span>
            </div>
            <p className="text-xs text-purple-100 mb-3">
              Unlock analytics, alerts, and advanced features
            </p>
            <Link
              to="/subscription"
              onClick={() => setMobileMenuOpen(false)}
              className="block w-full bg-white/20 hover:bg-white/30 text-white text-xs font-medium py-2 px-3 rounded-lg text-center transition-colors backdrop-blur-sm"
            >
              Upgrade Now
            </Link>
          </motion.div>
        )}

        {/* Dark Mode Toggle */}
        <motion.div 
          variants={navItemVariants}
          className="p-6 border-t border-gray-200/50 dark:border-gray-800/50"
        >
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium text-gray-600 dark:text-gray-400">Dark Mode</span>
            <Switch
              checked={darkMode}
              onChange={setDarkMode}
              className={`${
                darkMode 
                  ? 'bg-gradient-to-r from-blue-500 to-purple-500' 
                  : 'bg-gray-300 dark:bg-gray-600'
              } relative inline-flex h-7 w-12 items-center rounded-full transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2`}
            >
              <span className="sr-only">Toggle dark mode</span>
              <motion.span
                layout
                className={`inline-block h-5 w-5 transform rounded-full bg-white shadow-lg transition-all duration-300 ${
                  darkMode ? 'translate-x-6' : 'translate-x-1'
                }`}
                whileTap={{ scale: 0.9 }}
              />
              {darkMode ? (
                <Moon className="absolute left-1.5 top-1 w-4 h-4 text-blue-100" />
              ) : (
                <Sun className="absolute right-1.5 top-1 w-4 h-4 text-yellow-400" />
              )}
            </Switch>
          </div>
        </motion.div>
      </motion.aside>

      {/* Main Content */}
      <div className="flex-1 flex flex-col min-h-screen">
        {/* Top Bar */}
        <motion.header 
          initial={{ y: -20, opacity: 0 }}
          animate={{ y: 0, opacity: 1 }}
          transition={{ delay: 0.2 }}
          className="sticky top-0 z-30 bg-white/80 dark:bg-gray-950/80 backdrop-blur-xl border-b border-gray-200/50 dark:border-gray-800/50 flex items-center h-16 px-4 md:px-8 shadow-sm"
        >
          <div className="flex-1 flex items-center">
            <button
              onClick={() => setMobileMenuOpen(true)}
              className="md:hidden p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            >
              <Menu className="w-6 h-6 text-gray-600 dark:text-gray-400" />
            </button>
            <motion.h1 
              key={location.pathname}
              initial={{ opacity: 0, x: 20 }}
              animate={{ opacity: 1, x: 0 }}
              className="text-xl font-bold text-gray-900 dark:text-gray-100 ml-4 md:ml-0"
            >
              {navLinks.find(l => l.path === location.pathname)?.name || 'Dashboard'}
            </motion.h1>
          </div>
          <div className="flex items-center space-x-4">
            {/* Plan Badge */}
            <motion.div
              whileHover={{ scale: 1.05 }}
              className={`px-3 py-1 rounded-full text-xs font-medium shadow-sm ${
                isFreeUser 
                  ? 'bg-gray-100 text-gray-600 dark:bg-gray-800 dark:text-gray-400'
                  : 'bg-gradient-to-r from-yellow-400 to-orange-500 text-white'
              }`}
            >
              {isFreeUser ? 'Free' : 'Premium'}
            </motion.div>
            
            {/* Notifications */}
            <motion.button
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              className="relative p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            >
              <Bell className="w-5 h-5 text-gray-600 dark:text-gray-400" />
              <span className="absolute -top-1 -right-1 w-3 h-3 bg-red-500 rounded-full animate-pulse"></span>
            </motion.button>
            
            {/* User Menu */}
            <div className="relative">
              <motion.button 
                whileHover={{ scale: 1.05 }}
                onClick={() => setUserMenuOpen(!userMenuOpen)}
                className="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
              >
                <div className="w-8 h-8 rounded-full bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center font-bold text-white text-sm shadow-lg">
                  {user?.email?.charAt(0).toUpperCase() || 'U'}
                </div>
                <ChevronDown className={`w-4 h-4 text-gray-600 dark:text-gray-400 transition-transform ${userMenuOpen ? 'rotate-180' : ''}`} />
              </motion.button>

              {/* User Dropdown Menu */}
              <AnimatePresence>
                {userMenuOpen && (
                  <motion.div
                    initial={{ opacity: 0, y: -10, scale: 0.95 }}
                    animate={{ opacity: 1, y: 0, scale: 1 }}
                    exit={{ opacity: 0, y: -10, scale: 0.95 }}
                    className="absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 rounded-lg shadow-lg border border-gray-200 dark:border-gray-700 py-2 z-50"
                  >
                    <div className="px-4 py-2 border-b border-gray-200 dark:border-gray-700">
                      <p className="text-sm font-medium text-gray-900 dark:text-white">{user?.email}</p>
                      <p className="text-xs text-gray-500 dark:text-gray-400">{isFreeUser ? 'Free Plan' : 'Premium Plan'}</p>
                    </div>
                    <button
                      onClick={() => navigate('/settings')}
                      className="w-full flex items-center px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
                    >
                      <User className="w-4 h-4 mr-2" />
                      Profile Settings
                    </button>
                    <button
                      onClick={handleLogout}
                      className="w-full flex items-center px-4 py-2 text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors"
                    >
                      <LogOut className="w-4 h-4 mr-2" />
                      Sign Out
                    </button>
                  </motion.div>
                )}
              </AnimatePresence>
            </div>
          </div>
        </motion.header>

        {/* Page Content */}
        <AnimatePresence mode="wait">
          <motion.main
            key={location.pathname}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -20 }}
            transition={{ duration: 0.4, ease: 'easeInOut' }}
            className="flex-1 p-4 md:p-8"
          >
            {children}
          </motion.main>
        </AnimatePresence>
      </div>

      {/* Floating Action Button */}
      <motion.button
        initial={{ scale: 0, opacity: 0 }}
        animate={{ scale: 1, opacity: 1 }}
        transition={{ delay: 0.5, type: 'spring', stiffness: 200 }}
        whileHover={{ scale: 1.1 }}
        whileTap={{ scale: 0.9 }}
        className="fixed bottom-6 right-6 w-14 h-14 bg-gradient-to-r from-blue-500 to-purple-500 text-white rounded-full shadow-lg hover:shadow-xl transition-all duration-300 z-40 flex items-center justify-center backdrop-blur-sm"
      >
        <Plus className="w-6 h-6" />
      </motion.button>
    </div>
  )
} 