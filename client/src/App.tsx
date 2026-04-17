import { useEffect } from "react"
import { isAxiosError } from "axios"
import { Route, Routes } from "react-router-dom"
import { useDispatch } from "react-redux"
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
import AuthVerifyPage from "./pages/authVerify/AuthVerifyPage"
import RegisterPage from "./pages/register/RegisterPage"
import { authApi } from "./api/api"
import { setUser } from "./redux/features/users/userSlice"
import type { AppDispatch } from "./redux/store"

function App() {
  const dispatch = useDispatch<AppDispatch>()

  useEffect(() => {
    const restoreSession = async (): Promise<void> => {
      try {
        const response = await authApi.getCurrentUser()
        dispatch(setUser({
          id: response.user.id,
          fullname: response.user.fullname,
          email: response.user.email,
          role: response.user.role,
          telegram: response.user.telegram ?? "",
          phone: response.user.phone ?? "",
        }))
      } catch (error) {
        // 401 на /auth/me при отсутствии сессии — ожидаемое поведение
        if (isAxiosError(error) && error.response?.status === 401) {
          dispatch(setUser(null))
          return
        }

        dispatch(setUser(null))
      }
    }

    void restoreSession()
  }, [dispatch])

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
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/auth/verify" element={<AuthVerifyPage />} />
        <Route path="/profile" element={<ProfilePage />} />
        <Route path="*" element={<NotFound />} />
      </Routes>

      <Footer />
    </>
  )
}

export default App
