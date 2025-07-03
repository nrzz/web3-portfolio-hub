import React from 'react'
import { Routes, Route } from 'react-router-dom'
import { Layout } from './components/Layout'
import { Home } from './pages/Home'
import { Login } from './pages/Login'
import { Dashboard } from './pages/Dashboard'
import { Portfolio } from './pages/Portfolio'
import { Analytics } from './pages/Analytics'
import { Alerts } from './pages/Alerts'
import { Settings } from './pages/Settings'
import { Subscription } from './pages/Subscription'
import { ForumHome } from './pages/ForumHome'
import { AskQuestion } from './pages/AskQuestion'
import { AuthProvider } from './contexts/AuthContext'
import { QuestionDetail } from './pages/QuestionDetail'

function App() {
  return (
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/dashboard" element={<Layout><Dashboard /></Layout>} />
        <Route path="/portfolio" element={<Layout><Portfolio /></Layout>} />
        <Route path="/analytics" element={<Layout><Analytics /></Layout>} />
        <Route path="/alerts" element={<Layout><Alerts /></Layout>} />
        <Route path="/settings" element={<Layout><Settings /></Layout>} />
        <Route path="/subscription" element={<Layout><Subscription /></Layout>} />
        <Route path="/forum" element={<Layout><ForumHome /></Layout>} />
        <Route path="/forum/ask" element={<Layout><AskQuestion /></Layout>} />
        <Route path="/forum/:id" element={<Layout><QuestionDetail /></Layout>} />
      </Routes>
    </AuthProvider>
  )
}

export default App 