import { Route, Routes } from "react-router-dom"
import Footer from "./components/footer/Footer"
import Header from "./components/header/Header"
import IndexPage from "./pages/index/IndexPage"
import './styles/style.css'
import EventsPage from "./pages/events/EventsPage"
import MentorsPage from "./pages/mentors/MentorsPage"
import ProfilePage from "./pages/profile/ProfilePage"

function App() {
  return (
    <>
      <Header />

      <Routes>
        <Route path="/" element={<IndexPage />} />
        <Route path="/events" element={<EventsPage />} />
        <Route path="/mentors" element={<MentorsPage />} />
        <Route path="/profile" element={<ProfilePage />} />
      </Routes>

      <Footer />
    </>
  )
}

export default App
