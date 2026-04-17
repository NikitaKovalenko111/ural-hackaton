import { useEffect, useState } from "react"
import type { JSX } from "react"
import type React from "react"
import { isAxiosError } from "axios"
import { Link, Navigate } from "react-router-dom"
import { useParams } from "react-router-dom"
import RegistrationModal from "../../components/registrationModal/RegistrationModal"
import { eventsApi, hubsApi, mentorsApi } from "../../api/api"
import type { IEvent, IHub, IMentor } from "../../types"

const IndexPage: React.FC = (): JSX.Element => {
    const { hubId } = useParams<{ hubId: string }>()
    const parsedHubId = Number(hubId)
    const [isRegistrationOpen, setIsRegistrationOpen] = useState<boolean>(false)
    const [selectedEventTitle, setSelectedEventTitle] = useState<string>("Ближайшее событие хаба")
    const [hub, setHub] = useState<IHub | null>(null)
    const [hubEvents, setHubEvents] = useState<IEvent[]>([])
    const [hubMentors, setHubMentors] = useState<IMentor[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [isNotFound, setIsNotFound] = useState<boolean>(false)
    const [errorMessage, setErrorMessage] = useState<string>("")

    useEffect(() => {
        if (!Number.isFinite(parsedHubId) || parsedHubId <= 0) {
            setIsLoading(false)
            setIsNotFound(true)
            return
        }

        const loadHub = async (): Promise<void> => {
            setIsLoading(true)
            setIsNotFound(false)
            setErrorMessage("")

            try {
                const [hubData, eventsData, mentorsData] = await Promise.all([
                    hubsApi.getHubById(parsedHubId),
                    eventsApi.getEvents(),
                    mentorsApi.getMentors(),
                ])

                const hubEventsData = eventsData
                    .filter((event) => event.hubId === parsedHubId)
                    .sort((a, b) => new Date(a.start).getTime() - new Date(b.start).getTime())

                const hubMentorsData = mentorsData.filter((mentor) => mentor.hubId === parsedHubId)

                setHub(hubData)
                setHubEvents(hubEventsData)
                setHubMentors(hubMentorsData)
            } catch (error) {
                setHub(null)
                setHubEvents([])
                setHubMentors([])

                if (isAxiosError(error) && error.response?.status === 404) {
                    setIsNotFound(true)
                } else {
                    setErrorMessage("Не удалось загрузить страницу хаба")
                }
            } finally {
                setIsLoading(false)
            }
        }

        void loadHub()
    }, [parsedHubId])

    if (isLoading) {
        return (
            <main>
                <section className="section section--hubs-overview">
                    <div className="container">
                        <p className="hubs-search__empty">Загрузка страницы хаба...</p>
                    </div>
                </section>
            </main>
        )
    }

    if (isNotFound) {
        return <Navigate to="/hubs" replace />
    }

    if (errorMessage) {
        return (
            <main>
                <section className="section section--hubs-overview">
                    <div className="container">
                        <p className="hubs-search__empty">{errorMessage}</p>
                        <Link to="/hubs" className="btn btn--outline btn--sm">Вернуться к списку хабов</Link>
                    </div>
                </section>
            </main>
        )
    }

    if (!hub) {
        return <Navigate to="/hubs" replace />
    }

    const openRegistrationModal = (eventTitle: string): void => {
        setSelectedEventTitle(eventTitle)
        setIsRegistrationOpen(true)
    }

    const formatEventDateTime = (value: string): string => {
        const date = new Date(value)
        if (Number.isNaN(date.getTime())) {
            return value
        }

        return new Intl.DateTimeFormat("ru-RU", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
            hour: "2-digit",
            minute: "2-digit",
        }).format(date)
    }

    const nearestEvent = hubEvents[0]
    const hubDescription = hub.description?.trim() || "Описание хаба скоро появится"

    const statusText = hub.status === "open" ? "● Открыт" : hub.status === "busy" ? "● Высокая загрузка" : "● Закрыт"
    const statusLabel = hub.status === "open" ? "Открыт" : hub.status === "busy" ? "Высокая загрузка" : "Закрыт"
    const mapsQuery = encodeURIComponent(`${hub.city} ${hub.address}`)

    return (
        <main>
            <section className="hero">
                <div className="container hero__card">
                    <div className="hero__info">
                        <div className="status-header">
                            <div className="status-header__icon">🕒</div>
                            <div>
                                <h1 className="status-header__title">{hub.name}</h1>
                                <p className="status-header__schedule">График сегодня: {hub.schedule}</p>
                            </div>
                            <span className="badge--open">{statusText}</span>
                        </div>
                        <div className="stats__container">
                            <div className="stats__box">
                                <div className="stats__label">Количество людей сейчас</div>
                                <div className="stats__value">{hub.occupancy}</div>
                                <div className="stats__unit">чел.</div>
                                <div className="stats__update-text">Обн. 5 мин</div>
                            </div>
                            <div className="event-box">
                                <div className="event-box__header">
                                    <span>Ближайшее событие</span>
                                    <Link to="/events" className="link--small">Все события</Link>
                                </div>
                                {nearestEvent ? (
                                    <div className="event-box__details">
                                        <h3>{nearestEvent.name}</h3>
                                        <p>Начало: {formatEventDateTime(nearestEvent.start)}</p>
                                    </div>
                                ) : (
                                    <div className="event-box__details">
                                        <h3>Событий пока нет</h3>
                                        <p>Следите за обновлениями расписания</p>
                                    </div>
                                )}
                                <button
                                    type="button"
                                    className="btn btn--primary"
                                    onClick={() => openRegistrationModal(nearestEvent?.name ?? "Ближайшее событие хаба")}
                                    disabled={!nearestEvent}
                                >
                                    Записаться
                                </button>
                                <div className="event-box__status">🕒 Сейчас {hub.schedule} → {hub.status === "closed" ? "закрыт" : "открыт"}</div>
                            </div>
                        </div>
                    </div>
                    <div className="hero__image-wrapper">
                        <div className="hero__image"></div>
                        <div className="image-overlay">{hubDescription}</div>
                        <div className="wifi-badge">📶 Wi-Fi: free</div>
                    </div>
                </div>
            </section>

            <section className="section section--hubs-overview">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">О хабе</h2>
                            <p className="section__subtitle">{hubDescription}</p>
                        </div>
                        <Link to="/hubs" className="btn btn--outline btn--sm">Все хабы</Link>
                    </div>
                    <div className="contacts__wrapper hub-overview">
                        <div className="contacts__info hub-overview__info">
                            <div className="hub-overview__grid">
                                <article className="hub-overview__card">
                                    <p className="hub-overview__label">Локация</p>
                                    <p className="hub-overview__value">📍 {hub.city}, {hub.address}</p>
                                </article>
                                <article className="hub-overview__card">
                                    <p className="hub-overview__label">Режим работы</p>
                                    <p className="hub-overview__value">🕒 {hub.schedule}</p>
                                </article>
                                <article className="hub-overview__card">
                                    <p className="hub-overview__label">Статус</p>
                                    <p className="hub-overview__value">{statusLabel}</p>
                                </article>
                                <article className="hub-overview__card">
                                    <p className="hub-overview__label">Сейчас в хабе</p>
                                    <p className="hub-overview__value">👥 {hub.occupancy} чел.</p>
                                </article>
                            </div>
                            <p className="hub-overview__note">Лучше приходить вне пиковых часов: в утренние и дневные интервалы загрузка ниже.</p>
                        </div>
                        <div className="contacts__map hub-overview__map">
                            <div className="map-placeholder hub-overview__map-placeholder">
                                <div className="hub-overview__map-pin">📍 {hub.city}</div>
                                <p className="hub-overview__map-address">{hub.address}</p>
                                <a className="btn btn--primary btn--sm" href={`https://yandex.ru/maps/?text=${mapsQuery}`} target="_blank" rel="noreferrer">Открыть на карте</a>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <section className="section section--events">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Ближайшие события</h2>
                            <p className="section__subtitle">Следите за расписанием и присоединяйтесь</p>
                        </div>
                        <Link to="/events" className="btn btn--primary btn--sm">Все события</Link>
                    </div>
                    <div className="events-grid">
                        {hubEvents.map((eventItem) => (
                            <article className="event-card" key={eventItem.id}>
                                <div className="event-card__header">
                                    <span className="event-card__time">🕒 {formatEventDateTime(eventItem.start)}</span>
                                </div>
                                <h3 className="event-card__title">{eventItem.name}</h3>
                                <p className="event-card__description">{eventItem.description}</p>
                                <p className="event-card__role">Окончание: {formatEventDateTime(eventItem.end)}</p>
                                <div className="mentor-card__actions">
                                    <button type="button" className="btn btn--primary btn--sm" onClick={() => openRegistrationModal(eventItem.name)}>
                                        Записаться
                                    </button>
                                </div>
                            </article>
                        ))}
                    </div>
                    {hubEvents.length === 0 ? <p className="hubs-search__empty">В этом хабе пока нет событий.</p> : null}
                </div>
            </section>

            <section className="section section--mentors">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Лучшие менторы</h2>
                            <p className="section__subtitle">Эксперты, которые помогут ускорить ваш рост</p>
                        </div>
                        <Link to="/mentors" className="btn btn--primary btn--sm">Стать ментором</Link>
                    </div>
                    <div className="mentors-grid">
                        {hubMentors.map((mentor) => (
                            <article className="mentor-card" key={mentor.mentorId ?? mentor.id}>
                                <div className="mentor-card__top">
                                    <span className="tag">👤 Ментор</span>
                                    <span className="rating">ID: {mentor.mentorId ?? "-"}</span>
                                </div>
                                <h3 className="mentor-card__title">{mentor.fullname}</h3>
                                <p className="mentor-card__role">{mentor.email}</p>
                                <div className="mentor-card__actions">
                                    <Link to="/mentors" className="btn btn--primary btn--sm">Подробнее</Link>
                                </div>
                            </article>
                        ))}
                    </div>
                    {hubMentors.length === 0 ? <p className="hubs-search__empty">В этом хабе пока нет менторов.</p> : null}
                </div>
            </section>

            <section className="section section--contacts">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Контакты</h2>
                            <p className="section__subtitle">Всегда на связи — приходите, звоните, пишите</p>
                        </div>
                        <Link to="/events" className="btn btn--primary btn--sm">Как добраться</Link>
                    </div>
                    <div className="contacts__wrapper">
                        <div className="contacts__info">
                            <div className="contacts__item">📍 Город, Ул. Примерная, 15</div>
                            <div className="contacts__item">📞 +7 (000) 000-00-00</div>
                            <div className="contacts__item">✉️ hello@yellowhub.example</div>
                            <div className="contacts__item">✈️ Telegram</div>
                        </div>
                        <div className="contacts__map">
                            <div className="map-placeholder"></div>
                        </div>
                    </div>
                </div>
            </section>

            <RegistrationModal
                isOpen={isRegistrationOpen}
                onClose={() => setIsRegistrationOpen(false)}
                eventTitle={selectedEventTitle}
            />
        </main>
    )
}

export default IndexPage