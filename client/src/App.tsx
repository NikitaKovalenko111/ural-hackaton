import { Route, Routes } from "react-router-dom"
import Footer from "./components/footer/Footer"
import Header from "./components/header/Header"
import IndexPage from "./pages/index/IndexPage"
import "./styles/scss/style.scss"
import EventsPage from "./pages/events/EventsPage"
import MentorsPage from "./pages/mentors/MentorsPage"
import ProfilePage from "./pages/profile/ProfilePage"
import HubsPage from "./pages/hubs/HubsPage"
import NotFound from "./pages/notfound/NotFound"
import LoginPage from "./pages/login/LoginPage"

function App() {
  return (
    <>
      <Header />

      <Routes>
        <Route path="/" element={<HubsPage />} />
        <Route path="/hubs" element={<HubsPage />} />
        <Route path="/hubs/:hubId" element={<IndexPage />} />
        <Route path="/events" element={<EventsPage />} />
        <Route path="/mentors" element={<MentorsPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/profile" element={<ProfilePage />} />
        <Route path="*" element={<NotFound />} />
      </Routes>

      <Footer />
    </>
  )
}

export default App
