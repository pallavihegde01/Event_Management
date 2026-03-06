import React from 'react'
import HomePage from './pages/HomePage'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import EventPage from './pages/EventPage'

const App = () => {
  return (
     <BrowserRouter>

      <Routes>

        <Route path="/" element={<HomePage />} />

        <Route path="/event" element={<EventPage />} />

      </Routes>

    </BrowserRouter>
  )
}

export default App